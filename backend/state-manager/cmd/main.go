package main

import (
	"github.com/joho/godotenv"
	connectHandler "github.com/ueckoken/plarail2023/backend/state-manager/pkg/connect"
	"github.com/ueckoken/plarail2023/backend/state-manager/pkg/mqtt_handler"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	go connectHandler.StartHandler()
	//go operation.Handler()
	mqtt_handler.StartHandler()
}
