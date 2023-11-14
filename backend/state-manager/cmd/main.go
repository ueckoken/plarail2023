package main

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/joho/godotenv"
	"github.com/ueckoken/plarail2023/backend/spec/state/v1/statev1connect"
	connectHandler "github.com/ueckoken/plarail2023/backend/state-manager/pkg/connect"
	"github.com/ueckoken/plarail2023/backend/state-manager/pkg/mqtt_handler"
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

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Default().Error("Error loading .env file")
	}
	baseCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx, stop := signal.NotifyContext(baseCtx, os.Interrupt)
	defer stop()

	go func() {
		<-ctx.Done()
		slog.Default().Info("signal received")
	}()

	eg, ctx := errgroup.WithContext(ctx)

	r := chi.NewRouter()
	// r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/debug/ping"))
	r.Mount("/debug", middleware.Profiler())
	r.Handle(statev1connect.NewStateManagerServiceHandler(&connectHandler.StateManagerServer{}))
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
				LogLevel: slog.LevelInfo,
			},
		),
	),
	)

	srv := &http.Server{
		Addr:              net.JoinHostPort("0.0.0.0", "8080"),
		Handler:           h2c.NewHandler(r, &http2.Server{}),
		ReadHeaderTimeout: 60 * time.Second,
		BaseContext:       func(net.Listener) context.Context { return ctx },
	}
	eg.Go(srv.ListenAndServe)
	//go operation.Handler()
	eg.Go(func() error {
		slog.Default().Info("start mqtt handler")
		return mqtt_handler.StartHandler(ctx)
	})

	if err := eg.Wait(); err != nil {
		slog.Default().Error("error in sub goroutine at main", err)
	}
	newCtx, srvTimeOutCancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer srvTimeOutCancel()
	srv.Shutdown(newCtx)
  <-newCtx.Done()
}
