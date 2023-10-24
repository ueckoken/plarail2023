package mqtt

import (
	"log"
	"os"

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
