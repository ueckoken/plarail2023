package mqtt

import (
	"github.com/joho/godotenv"
	"testing"
)

func Test_SendMsg(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	client := MakeClient()
	Send(client, "test", "test")
}
