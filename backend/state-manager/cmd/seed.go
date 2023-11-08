// Seed DB Data FROM YAML

package main

import (
	"github.com/joho/godotenv"
	statev1 "github.com/ueckoken/plarail2023/backend/spec/state/v1"
	"github.com/ueckoken/plarail2023/backend/state-manager/pkg/db"
	"gopkg.in/yaml.v3"
	"os"
)

type Station string
type StopRail string
type PointRail string
type Block string

type Seed struct {
	Stations   []Station   `yaml:"stations"`
	StopRails  []StopRail  `yaml:"stop_rails"`
	PointRails []PointRail `yaml:"point_rails"`
	Blocks     []Block     `yaml:"blocks"`
}

func main() {
	err := godotenv.Load(".env")
	defer db.C()
	db.Open()
	data := &Seed{}
	b, _ := os.ReadFile("./data/nt-tokyo.yaml")
	err = yaml.Unmarshal(b, data)
	if err != nil {
		panic(err)
	}

	//for _, stop := range data.StopRails {
	//	println(stop)
	//	err := db.AddStop(&trainv1.StopAndState{
	//		Id:    string(stop),
	//		State: trainv1.StopStateEnum_STOP_STATE_GO,
	//	})
	//	if err != nil {
	//		return
	//	}
	//}
	//
	//for _, point := range data.PointRails {
	//	println(point)
	//	err := db.AddPoint(&trainv1.PointAndState{
	//		Id:    string(point),
	//		State: trainv1.PointStateEnum_POINT_STATE_NORMAL,
	//	})
	//	if err != nil {
	//		return
	//	}
	//}
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
