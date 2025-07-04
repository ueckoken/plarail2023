package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/ueckoken/plarail2023/backend/spec/state/v1/statev1connect"
	connectHandler "github.com/ueckoken/plarail2023/backend/state-manager/pkg/connect"
	"github.com/ueckoken/plarail2023/backend/state-manager/pkg/db"
	"github.com/ueckoken/plarail2023/backend/state-manager/pkg/mqtt_handler"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"
)

type AppEnv uint

const (
	Dev AppEnv = iota
	Test
	Prod
)

func (a AppEnv) String() string {
	switch a {
	case Dev:
		return "dev"
	case Test:
		return "test"
	case Prod:
		return "prod"
	default:
		return "unknown"
	}
}

var appEnv AppEnv = Dev
var (
	version = "develop"
	commit  = "deadbeef"
)

func init() {
	switch os.Getenv("APP_ENV") {
	case "prod":
		appEnv = Prod
	case "test":
		appEnv = Test
	default:
		appEnv = Dev
	}
	logger := func(appEnv AppEnv) *slog.Logger {
		switch appEnv {
		case Prod:
			return slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
				Level:     slog.LevelDebug,
				AddSource: true,
			}))
		default:
			return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
				Level:     slog.LevelDebug,
				AddSource: true,
			}))
		}
	}(appEnv)
	logger = logger.With("version", version).With("commit", commit).With("app_env", appEnv)
	slog.SetDefault(logger)
}

func NewTlsConfig() (*tls.Config, error) {
	certpool := x509.NewCertPool()
	ca, err := os.ReadFile("emqxsl-ca.pem")
	if err != nil {
		return nil, fmt.Errorf("failed to read CA certificate: %w", err)
	}
	ok := certpool.AppendCertsFromPEM(ca)
	if !ok {
		return nil, fmt.Errorf("failed to parse CA certificate")
	}
	return &tls.Config{
		RootCAs: certpool,
	}, nil
}

func main() {
	if os.Getenv("APP_ENV") != "prod" {
		err := godotenv.Load(".env")
		if err != nil {
			slog.Default().Error("Error loading .env file", slog.Any("error", err))
			os.Exit(1)
		}
	}
	baseCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx, stop := signal.NotifyContext(baseCtx, os.Interrupt)
	defer stop()
	go func() {
		<-ctx.Done()
		slog.Default().Info("signal received or canceled")
	}()

	eg, ctx := errgroup.WithContext(ctx)

	DBOpts := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	DBHandler, err := db.Open(ctx, DBOpts)
	if err != nil {
		slog.Default().Error("database connection failed", slog.Any("err", err))
		cancel()
		return
	}
	mqttClientOpts := mqtt.NewClientOptions()
	mqttClientOpts.AddBroker(os.Getenv("MQTT_BROKER_ADDR"))
	mqttClientOpts.Username = os.Getenv("MQTT_USERNAME")
	mqttClientOpts.Password = os.Getenv("MQTT_PASSWORD")
	mqttClientOpts.ClientID = os.Getenv("MQTT_CLIENT_ID")
	tlsconfig, err := NewTlsConfig()
	if err != nil {
		slog.Default().Error("failed to create TLS config", slog.Any("err", err))
		cancel()
		return
	}
	mqttClientOpts.SetTLSConfig(tlsconfig)

	mqttHandler, err := mqtt_handler.NewHandler(mqttClientOpts, DBHandler)
	if err != nil {
		slog.Default().Error("mqtt create client or handler failed,", slog.Any("err", err))
		cancel()
		return
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/debug/ping"))
	r.Use(httplog.RequestLogger(
		httplog.NewLogger(
			"http_server",
			httplog.Options{
				JSON:           appEnv == Prod,
				Concise:        false,
				RequestHeaders: true,
				Writer:         os.Stdout,
				Tags: map[string]string{
					"version": version,
					"commit":  commit,
				},
				LogLevel: slog.LevelWarn,
			},
		),
	),
	)
	r.Mount("/debug", middleware.Profiler())
	r.Mount(statev1connect.NewStateManagerServiceHandler(&connectHandler.StateManagerServer{DBHandler: DBHandler, MqttHandler: mqttHandler}))

	srv := &http.Server{
		Addr:              net.JoinHostPort("0.0.0.0", "8080"),
		Handler:           cors.AllowAll().Handler(h2c.NewHandler(r, &http2.Server{})),
		ReadHeaderTimeout: 60 * time.Second,
		BaseContext:       func(net.Listener) context.Context { return ctx },
	}
	eg.Go(func() error {
		errC := make(chan error)
		go func() {
			slog.Default().Info("start http server")
			if err := srv.ListenAndServe(); err != nil {
				slog.Default().Error("http server error", slog.Any("error", err))
				errC <- err
			}
		}()
		select {
		case err := <-errC:
			return fmt.Errorf("http server error: %w", err)
		case <-ctx.Done():
			slog.Default().Info("Interrupted at http server")
			return ctx.Err()
		}
	})
	// go operation.Handler()
	eg.Go(func() error {
		slog.Default().Info("start mqtt handler")
		err := mqttHandler.Start(ctx)
		if err != nil {
			return fmt.Errorf("mqtt handler error: %w", err)
		}
		return nil
	})

	// errGroup.Waitはeg.Goが全てerrorを返すまでwaitする
	if err := eg.Wait(); err != nil {
		slog.Default().Error("error in sub goroutine at main", slog.Any("error", err))
	}
	slog.Default().Info("shutting down server")
	newCtx, srvTimeOutCancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer srvTimeOutCancel()
	if err := srv.Shutdown(newCtx); err != nil {
		slog.Default().Error("failed to shutdown server gracefully", slog.Any("error", err))
	}
	
	// Close database connection
	DBHandler.Close(newCtx)
}
