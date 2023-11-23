// Seed DB Data FROM YAML

package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	statev1 "github.com/ueckoken/plarail2023/backend/spec/state/v1"
	dbhandler "github.com/ueckoken/plarail2023/backend/state-manager/pkg/db"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v3"
)

type StopRail string
type PointRail string
type Block string

type Seed struct {
	StopRails  []StopRail  `yaml:"stop_rails"`
	PointRails []PointRail `yaml:"point_rails"`
	Blocks     []Block     `yaml:"blocks"`
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	db, err := dbhandler.Open(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		return
	}
	defer db.Close()
	data := &Seed{}
	b, _ := os.ReadFile("./data/chofufes-2023.yaml")
	if err := yaml.Unmarshal(b, data); err != nil {
		panic(err)
	}

	for _, stop := range data.StopRails {
		println(stop)
		err := db.AddStop(&statev1.StopAndState{
			Id:    string(stop),
			State: statev1.StopStateEnum_STOP_STATE_GO,
		})
		if err != nil {
			return
		}
	}

	for _, point := range data.PointRails {
		println(point)
		err := db.AddPoint(&statev1.PointAndState{
			Id:    string(point),
			State: statev1.PointStateEnum_POINT_STATE_NORMAL,
		})
		if err != nil {
			return
		}
	}
	for _, block := range data.Blocks {
		println(block)
		err := db.AddBlock(&statev1.BlockState{
			BlockId: string(block),
			State:   statev1.BlockStateEnum_BLOCK_STATE_OPEN,
		})
		if err != nil {
			return
		}
	}
}
