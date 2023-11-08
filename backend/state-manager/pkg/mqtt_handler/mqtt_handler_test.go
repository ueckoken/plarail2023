package mqtt_handler

import (
	"github.com/joho/godotenv"
	"testing"
)

func Test_SendMsg(t *testing.T) {
	err := godotenv.Load("../../cmd/.env")
	if err != nil {
		panic(err)
	}
	client := MakeClient()
	Send(client, "test", "test")
}
