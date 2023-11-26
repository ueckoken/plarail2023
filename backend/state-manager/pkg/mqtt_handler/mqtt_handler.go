package mqtt_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	statev1 "github.com/ueckoken/plarail2023/backend/spec/state/v1"
	"github.com/ueckoken/plarail2023/backend/state-manager/pkg/db"
)

type Handler struct {
	client    mqtt.Client
	dbHandler *db.DBHandler
}

func NewHandler(clientOpts *mqtt.ClientOptions, dbHandler *db.DBHandler) (*Handler, error) {
	cc := mqtt.NewClient(clientOpts)

	if token := cc.Connect(); token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("mqtt error: %w", token.Error())
	}
	return &Handler{client: cc, dbHandler: dbHandler}, nil
}

func (h *Handler) Start(ctx context.Context) error {
	msgCh := make(chan mqtt.Message)
	var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		msgCh <- msg
	}
	h.Subscribe([]string{"point/#", "stop/#", "block/#", "train/#", "setting/#"}, f)

	for {
		select {
		case msg := <-msgCh:
			// if topic start with "point/"
			log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
			h.topicHandler(msg)
		case <-ctx.Done():
			slog.Default().Info("Interrupted at mqtt_handler")
			h.client.Disconnect(1000)
			slog.Default().Info("Disconnected from mqtt broker")
			return nil
		}
	}
}

func (h *Handler) Subscribe(topic []string, f mqtt.MessageHandler) {
	qos := byte(1)

	filters := make(map[string]byte)
	for _, t := range topic {
		filters[t] = qos
	}

	subscribeToken := h.client.SubscribeMultiple(filters, f)
	if subscribeToken.Wait() && subscribeToken.Error() != nil {
		log.Fatal(subscribeToken.Error())
	}
}

func (h *Handler) Send(topic string, payload string) {
	token := h.client.Publish(topic, 0, false, payload)
	token.Wait()
}

/*
	Endpoint
	{target}/{pointId}/get
	{target}/{pointId}/delta
	{target}/{pointId}/update
*/

func (h *Handler) topicHandler(msg mqtt.Message) error {
	// Handle by Path
	arr := strings.Split(msg.Topic(), "/")
	target := arr[0]
	id := arr[1]
	method := arr[2]

	if len(arr) > 3 {
		return fmt.Errorf("message format failed")
	}

	switch method {
	case "get":
		h.getState(target, id)
	case "delta":
		h.getDelta(target, id)
	case "update":
		h.updateState(target, id, msg.Payload())
	}
	return nil
}

func (h *Handler) NotifyStateUpdate(target string, id string, state string) {
	token := h.client.Publish(target+"/"+id+"/delta", 0, false, state)
	token.Wait()
}

func (h *Handler) getState(target string, id string) error {
	switch target {
	case "point":
		point, err := h.dbHandler.GetPoint(id)
		if err != nil {
			return fmt.Errorf("Get point failed, %w", err)
		}
		log.Println(point)
		token := h.client.Publish("point/"+id+"/get/accepted", 0, false, point.State.String())
		token.Wait()

	case "stop":
		stop, err := h.dbHandler.GetStop(id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(stop)
		token := h.client.Publish("stop/"+id+"/get/accepted", 0, false, stop.State.String())
		token.Wait()

	case "block":
		fmt.Println("block")
		block, err := h.dbHandler.GetBlock(id)
		if err != nil {
			return fmt.Errorf("getblock in block failed, %w", err)
		}
		res, err := json.Marshal(block)
		if err != nil {
			slog.Default().Info("invaild json marshaled in mqtt_handler.NotifyStateUpdate", slog.Any("err", err))
		}
		token := h.client.Publish("block/"+id+"/get/accepted", 0, false, res)
		token.Wait()

	case "setting":
		// read from /setting/esp/{id}.json
		// check file exists
		_, err := os.Stat("../settings/esp/" + id + ".json")
		if err != nil {
			log.Println(err.Error())
			// Return error message
			token := h.client.Publish("setting/"+id+"/get/accepted", 0, false, "error")
			token.Wait()
			return fmt.Errorf("settings %w", err)
		}
		raw, err := os.ReadFile("../settings/esp/" + id + ".json")
		if err != nil {
			log.Println(err.Error())
			return nil
		}
		// remove \n code
		raw = []byte(strings.Replace(string(raw), "\n", "", -1))
		raw = []byte(strings.Replace(string(raw), " ", "", -1))
		token := h.client.Publish("setting/"+id+"/get/accepted", 0, false, string(raw))
		token.Wait()

	case "train":
		// TODO: implement
	}
	return nil
}

func (h *Handler) getDelta(target string, id string) {

}

func (h *Handler) updateState(target string, id string, payload []byte) {

	switch target {
	case "block":
		// Check State
		fmt.Print(id)
		newState := string(payload)
		fmt.Print(" change to ")
		fmt.Println(newState)
		if newState == "OPEN" {
			err := h.dbHandler.UpdateBlock(&statev1.BlockState{
				BlockId: id,
				State:   statev1.BlockStateEnum_BLOCK_STATE_OPEN,
			})
			if err != nil {
				slog.Warn("Update block failed", slog.Any("err", err))
			}
		} else if newState == "CLOSE" {
			err := h.dbHandler.UpdateBlock(&statev1.BlockState{
				BlockId: id,
				State:   statev1.BlockStateEnum_BLOCK_STATE_CLOSE,
			})
			if err != nil {
				slog.Warn("Update block failed", slog.Any("err", err))
			}
		}
	}
}
