// ---------------------------------------------------------------
//
//	publish.go
//
//					Jan/25/2021
//
// ---------------------------------------------------------------
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// ---------------------------------------------------------------
func main() {
	fmt.Fprintf(os.Stderr, "*** 開始 ***\n")
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")
	cc := mqtt.NewClient(opts)

	if token := cc.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Mqtt error: %s", token.Error())
	}

	for it := 0; it < 5; it++ {
		now := time.Now()
		text := fmt.Sprintf("こんにちは %d %s", it, now)
		token := cc.Publish("go-mqtt/sample", 0, false, text)
		token.Wait()
	}

	cc.Disconnect(250)

	fmt.Println("Complete publish")
	fmt.Fprintf(os.Stderr, "*** 終了 ***\n")
}

// ---------------------------------------------------------------
