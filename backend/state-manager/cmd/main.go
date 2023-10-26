package main

import (
	"github.com/joho/godotenv"
	mqtt "github.com/ueckoken/plarail2023/backend/state-manager/pkg/mqtt"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	mqtt.StartHandler()
}
