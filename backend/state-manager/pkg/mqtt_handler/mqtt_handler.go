package mqtt_handler

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	statev1 "github.com/ueckoken/plarail2023/backend/state-manager/spec/state/v1"
	"github.com/ueckoken/plarail2023/backend/state-manager/pkg/db"
	"log"
	"os"
	"strings"
)

import (
	"fmt"
	"os/signal"
)

var cc mqtt.Client

func MakeClient() mqtt.Client {
	var opts = mqtt.NewClientOptions()
	opts.AddBroker(os.Getenv("MQTT_BROKER_ADDR"))
	opts.Username = os.Getenv("MQTT_USERNAME")
	opts.Password = os.Getenv("MQTT_PASSWORD")
	opts.ClientID = os.Getenv("MQTT_CLIENT_ID")

	cc = mqtt.NewClient(opts)

	if token := cc.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Mqtt error: %s", token.Error())
	}

	return cc
}

func Subscribe(cc mqtt.Client, topic []string, f mqtt.MessageHandler) {
	qos := byte(1)

	filters := make(map[string]byte)
	for _, t := range topic {
		filters[t] = qos
	}

	subscribeToken := cc.SubscribeMultiple(filters, f)
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
	Subscribe(cc, []string{"point/#", "stop/#", "block/#", "train/#"}, f)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	for {
		select {
		case msg := <-msgCh:
			// if topic start with "point/"
			log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
			topicHandler(cc, msg)
		case <-signalCh:
			fmt.Println("Interrupted")
			cc.Disconnect(1000)
			return
		}
	}
}

/*
	Endpoint
	{target}/{pointId}/get
	{target}/{pointId}/delta
	{target}/{pointId}/update
*/

func topicHandler(cc mqtt.Client, msg mqtt.Message) {
	// Handle by Path
	arr := strings.Split(msg.Topic(), "/")
	target := arr[0]
	id := arr[1]
	method := arr[2]

	if len(arr) > 3 {
		return
	}

	switch method {
	case "get":
		getState(cc, target, id)
	case "delta":
		getDelta(cc, target, id)
	case "update":
		updateState(cc, target, id, msg.Payload())
	}
}

func NotifyStateUpdate(target string, id string, state string) {
	token := cc.Publish(target+"/"+id+"/delta", 0, false, state)
	token.Wait()
}

func getState(cc mqtt.Client, target string, id string) {
	defer db.C()
	db.Open()

	switch target {
	case "point":
		point, err := db.GetPoint(id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(point)
		token := cc.Publish("point/"+id+"/get/accepted", 0, false, point.State.String())
		token.Wait()

	case "stop":
		stop, err := db.GetStop(id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(stop)
		token := cc.Publish("stop/"+id+"/get/accepted", 0, false, stop.State.String())
		token.Wait()

	case "block":
		block, err := db.GetBlock(id)
		if err != nil {
			log.Fatal(err)
		}
		res, err := json.Marshal(block)
		token := cc.Publish("block/"+id+"/get/accepted", 0, false, res)
		token.Wait()

	case "train":
		// TODO: implement
	}
}

func getDelta(cc mqtt.Client, target string, id string) {

}

func updateState(cc mqtt.Client, target string, id string, payload []byte) {
	defer db.C()
	db.Open()

	switch target {
	case "block":
		// Check State
		newState := string(payload)
		fmt.Print("newState: ")
		fmt.Println(newState)
		if newState == "OPEN" {
			err := db.UpdateBlock(&statev1.BlockState{
				BlockId: id,
				State:   statev1.BlockStateEnum_BLOCK_STATE_OPEN,
			})
			if err != nil {
				log.Fatal(err)
			}
			// NT Tokyo
			if id == "yamashita_b1" {
				err := db.UpdateStop(&statev1.StopAndState{
					Id:    "yamashita_s1",
					State: statev1.StopStateEnum_STOP_STATE_GO,
				})
				if err != nil {
					log.Fatal(err)
				}
				NotifyStateUpdate("stop", "yamashita_s1", statev1.StopStateEnum_STOP_STATE_GO.String())
				err = db.UpdateStop(&statev1.StopAndState{
					Id:    "yamashita_s2",
					State: statev1.StopStateEnum_STOP_STATE_GO,
				})
				if err != nil {
					log.Fatal(err)
				}
				NotifyStateUpdate("stop", "yamashita_s2", statev1.StopStateEnum_STOP_STATE_GO.String())

				// 今と逆にする
				now, err := db.GetPoint("yamashita_p1")
				if err != nil {
					log.Fatal(err)
				}
				var newS statev1.PointStateEnum
				if now.State == statev1.PointStateEnum_POINT_STATE_NORMAL {
					newS = statev1.PointStateEnum_POINT_STATE_REVERSE
				} else {
					newS = statev1.PointStateEnum_POINT_STATE_NORMAL
				}
				err = db.UpdatePoint(&statev1.PointAndState{
					Id:    "yamashita_p1",
					State: newS,
				})

				if err != nil {
					log.Fatal(err)
				}

				NotifyStateUpdate("point", "yamashita_p1", newS.String())

			}
		} else if newState == "CLOSE" {
			err := db.UpdateBlock(&statev1.BlockState{
				BlockId: id,
				State:   statev1.BlockStateEnum_BLOCK_STATE_CLOSE,
			})
			if err != nil {
				log.Fatal(err)
			}
			// NT Tokyo
			if id == "yamashita_b1" {
				err := db.UpdateStop(&statev1.StopAndState{
					Id:    "yamashita_s1",
					State: statev1.StopStateEnum_STOP_STATE_STOP,
				})
				if err != nil {
					log.Fatal(err)
				}
				NotifyStateUpdate("stop", "yamashita_s1", statev1.StopStateEnum_STOP_STATE_STOP.String())
				err = db.UpdateStop(&statev1.StopAndState{
					Id:    "yamashita_s2",
					State: statev1.StopStateEnum_STOP_STATE_STOP,
				})
				if err != nil {
					log.Fatal(err)
				}
				NotifyStateUpdate("stop", "yamashita_s2", statev1.StopStateEnum_STOP_STATE_STOP.String())
			}
		}

	}
}
