package main

import (
	"context"

	"github.com/joho/godotenv"
	connectHandler "github.com/ueckoken/plarail2023/backend/state-manager/pkg/connect"
	"github.com/ueckoken/plarail2023/backend/state-manager/pkg/mqtt_handler"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	ctx :=context.Background()
	go connectHandler.StartHandler(ctx)
	//go operation.Handler()
	mqtt_handler.StartHandler()
}
