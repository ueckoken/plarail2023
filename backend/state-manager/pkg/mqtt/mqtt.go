package mqtt

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func MakeClient() mqtt.Client {
	var opts = mqtt.NewClientOptions()
	opts.AddBroker(os.Getenv("MQTT_BROKER_ADDR"))
	opts.Username = os.Getenv("MQTT_USERNAME")
	opts.Password = os.Getenv("MQTT_PASSWORD")
	opts.ClientID = os.Getenv("MQTT_CLIENT_ID")

	cc := mqtt.NewClient(opts)

	if token := cc.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Mqtt error: %s", token.Error())
	}

	return cc
}

func Subscribe(cc mqtt.Client, topic string, f mqtt.MessageHandler) {
	subscribeToken := cc.Subscribe(topic, 0, f)
	if subscribeToken.Wait() && subscribeToken.Error() != nil {
		log.Fatal(subscribeToken.Error())
	}
}

func Send(cc mqtt.Client, topic string, payload string) {
	token := cc.Publish(topic, 0, false, payload)
	token.Wait()
}

func StartHandler() {
	msgCh := make(chan mqtt.Message)
	var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		msgCh <- msg
	}
	cc := MakeClient()
	Subscribe(cc, "point/#", f)
	Subscribe(cc, "stop/#", f)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	for {
		select {
		case msg := <-msgCh:
			// if topic start with "point/"
			if msg.Topic()[0:6] == "point/" {
				pointHandler(msg)
			}
			// if topic start with "stop/"
			if msg.Topic()[0:5] == "stop/" {
				stopHandler(msg)
			}
		case <-signalCh:
			fmt.Println("Interrupted")
			cc.Disconnect(1000)
			return
		}
	}
}

func pointHandler(msg mqtt.Message) {
	log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

func stopHandler(msg mqtt.Message) {
	log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}
