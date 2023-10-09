// Mqtt event handler

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ueckoken/plarail2023/backend/mqtt-handler/pkg/connect"
	"github.com/ueckoken/plarail2023/backend/mqtt-handler/pkg/handler"
	mqttclient "github.com/ueckoken/plarail2023/backend/mqtt-handler/pkg/mqtt"
	"github.com/ueckoken/plarail2023/backend/proto/train/v1/trainv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	msgCh := make(chan mqtt.Message)
	var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		msgCh <- msg
	}
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	// Subscribe topics
	cc := mqttclient.MakeClient()
	mqttclient.Subscribe(cc, "plarail2023/#", f)

	// Make Connect handler
	server := &connect.TrainServer{}
	mux := http.NewServeMux()
	path, trainServiceHandler := trainv1connect.NewTrainServiceHandler(server)
	mux.Handle(path, trainServiceHandler)
	http.ListenAndServe(
		os.Getenv("CONNECT_LISTEN_ADDR"),
		h2c.NewHandler(mux, &http2.Server{}),
	)

	// Wait for messages
	for {
		select {
		case m := <-msgCh:
			handler.HandleMqttMessage(m.Topic(), string(m.Payload()))
		case <-signalCh:
			fmt.Printf("interrupted")
			cc.Disconnect(1000)
			return
		}
	}

}
