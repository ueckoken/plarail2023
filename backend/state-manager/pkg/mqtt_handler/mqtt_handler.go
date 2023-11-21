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
			h.topicHandler(ctx, msg)
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

func (h *Handler) topicHandler(ctx context.Context, msg mqtt.Message) {
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
		h.getState(ctx, target, id)
	case "delta":
		h.getDelta(ctx, target, id)
	case "update":
		h.updateState(ctx, target, id, msg.Payload())
	}
}

func (h *Handler) NotifyStateUpdate(ctx context.Context, target string, id string, state string) {
	token := h.client.Publish(target+"/"+id+"/delta", 0, false, state)
	select {
	case <-token.Done():
		slog.Default().Info("token done in mqtt_handler.NotifyStateUpdate", slog.Any("err", token.Error()))
	case <-ctx.Done():
		slog.Default().Info("context done in mqtt_handler.NotifyStateUpdate", slog.Any("err", ctx.Err()))
	}
}

func (h *Handler) getState(ctx context.Context, target string, id string) error {
	switch target {
	case "point":
		point, err := h.dbHandler.GetPoint(ctx, id)
		if err != nil {
			slog.Default().Info("db error in mqtt_handler.getState", slog.Any("err", err))
		}
		log.Println(point)
		token := h.client.Publish("point/"+id+"/get/accepted", 0, false, point.State.String())
		token.Wait()

	case "stop":
		stop, err := h.dbHandler.GetStop(ctx, id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(stop)
		token := h.client.Publish("stop/"+id+"/get/accepted", 0, false, stop.State.String())
		token.Wait()

	case "block":
		block, err := h.dbHandler.GetBlock(ctx, id)
		if err != nil {
			log.Fatal(err)
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
			return fmt.Errorf("setting file not found: %w", err)
		}
		raw, err := os.ReadFile("../settings/esp/" + id + ".json")
		if err != nil {
			slog.Default().Error("setting file read error", slog.Any("err", err))
			return fmt.Errorf("setting file read error: %w", err)
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

func (h *Handler) getDelta(ctx context.Context, target string, id string) {

}

func (h *Handler) updateState(ctx context.Context, target string, id string, payload []byte) {

	switch target {
	case "block":
		// Check State
		newState := string(payload)
		fmt.Print("newState: ")
		fmt.Println(newState)
		if newState == "OPEN" {
			err := h.dbHandler.UpdateBlock(ctx, &statev1.BlockState{
				BlockId: id,
				State:   statev1.BlockStateEnum_BLOCK_STATE_OPEN,
			})
			if err != nil {
				log.Fatal(err)
			}
			// NT Tokyo
			if id == "yamashita_b1" {
				err := h.dbHandler.UpdateStop(ctx, &statev1.StopAndState{
					Id:    "yamashita_s1",
					State: statev1.StopStateEnum_STOP_STATE_GO,
				})
				if err != nil {
					log.Fatal(err)
				}
				h.NotifyStateUpdate(ctx, "stop", "yamashita_s1", statev1.StopStateEnum_STOP_STATE_GO.String())
				err = h.dbHandler.UpdateStop(ctx, &statev1.StopAndState{
					Id:    "yamashita_s2",
					State: statev1.StopStateEnum_STOP_STATE_GO,
				})
				if err != nil {
					log.Fatal(err)
				}
				h.NotifyStateUpdate(ctx, "stop", "yamashita_s2", statev1.StopStateEnum_STOP_STATE_GO.String())

				// 今と逆にする
				now, err := h.dbHandler.GetPoint(ctx, "yamashita_p1")
				if err != nil {
					log.Fatal(err)
				}
				var newS statev1.PointStateEnum
				if now.State == statev1.PointStateEnum_POINT_STATE_NORMAL {
					newS = statev1.PointStateEnum_POINT_STATE_REVERSE
				} else {
					newS = statev1.PointStateEnum_POINT_STATE_NORMAL
				}
				err = h.dbHandler.UpdatePoint(ctx, &statev1.PointAndState{
					Id:    "yamashita_p1",
					State: newS,
				})

				if err != nil {
					log.Fatal(err)
				}

				h.NotifyStateUpdate(ctx, "point", "yamashita_p1", newS.String())

			}
		} else if newState == "CLOSE" {
			err := h.dbHandler.UpdateBlock(ctx, &statev1.BlockState{
				BlockId: id,
				State:   statev1.BlockStateEnum_BLOCK_STATE_CLOSE,
			})
			if err != nil {
				log.Fatal(err)
			}
			// NT Tokyo
			if id == "yamashita_b1" {
				err := h.dbHandler.UpdateStop(ctx, &statev1.StopAndState{
					Id:    "yamashita_s1",
					State: statev1.StopStateEnum_STOP_STATE_STOP,
				})
				if err != nil {
					log.Fatal(err)
				}
				h.NotifyStateUpdate(ctx, "stop", "yamashita_s1", statev1.StopStateEnum_STOP_STATE_STOP.String())
				err = h.dbHandler.UpdateStop(ctx, &statev1.StopAndState{
					Id:    "yamashita_s2",
					State: statev1.StopStateEnum_STOP_STATE_STOP,
				})
				if err != nil {
					log.Fatal(err)
				}
				h.NotifyStateUpdate(ctx, "stop", "yamashita_s2", statev1.StopStateEnum_STOP_STATE_STOP.String())
			}
		}

	}
}
