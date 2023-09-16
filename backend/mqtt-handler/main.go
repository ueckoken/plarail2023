// Mqtt event handler

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	fmt.Fprintf(os.Stderr, "*** 開始 ***\n")
	msgCh := make(chan mqtt.Message)
	var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		msgCh <- msg
	}
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")
	cc := mqtt.NewClient(opts)

	if token := cc.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Mqtt error: %s", token.Error())
	}

	if subscribeToken := cc.Subscribe("go-mqtt/sample", 0, f); subscribeToken.Wait() && subscribeToken.Error() != nil {
		log.Fatal(subscribeToken.Error())
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	for {
		select {
		case m := <-msgCh:
			fmt.Printf("topic: %v, payload: %v\n", m.Topic(), string(m.Payload()))
		case <-signalCh:
			fmt.Printf("Interrupt detected.\n")
			cc.Disconnect(1000)
			return
		}
	}
}
