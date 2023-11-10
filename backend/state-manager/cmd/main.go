package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	connectHandler "github.com/ueckoken/plarail2023/backend/state-manager/pkg/connect"
	"github.com/ueckoken/plarail2023/backend/state-manager/pkg/mqtt_handler"
	"golang.org/x/sync/errgroup"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	baseCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx, stop := signal.NotifyContext(baseCtx, os.Interrupt)
	defer stop()

	go func() {
		<-ctx.Done()
		log.Println("signal received")
	}()

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		log.Println("start connect server")
		return connectHandler.StartHandler(ctx)
	})
	//go operation.Handler()
	eg.Go(func() error {
		log.Println("start mqtt handler")
		return mqtt_handler.StartHandler(ctx)
	})

	if err := eg.Wait(); err != nil {
		log.Fatalln("error in sub goroutine at main: ", err)
	}
}
