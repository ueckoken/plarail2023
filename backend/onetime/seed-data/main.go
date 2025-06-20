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
	// .envファイルが存在する場合のみ読み込む（開発環境用）
	_ = godotenv.Load(".env")
	
	// 環境変数から設定を取得
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		panic("MONGODB_URI is not set")
	}
	
	dataFile := os.Getenv("SEED_DATA_FILE")
	if dataFile == "" {
		dataFile = "./data/mfk-2024.yaml" // デフォルト値
	}
	
	db, err := dbhandler.Open(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}
	defer db.Close(context.TODO())
	
	data := &Seed{}
	b, err := os.ReadFile(dataFile)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(b, data); err != nil {
		panic(err)
	}

	for _, stop := range data.StopRails {
		println("Adding stop rail:", stop)
		err := db.AddStop(&statev1.StopAndState{
			Id:    string(stop),
			State: statev1.StopStateEnum_STOP_STATE_GO,
		})
		if err != nil {
			println("Error adding stop rail", stop, ":", err.Error())
			// エラーが発生しても処理を続行
		}
	}

	for _, point := range data.PointRails {
		println("Adding point rail:", point)
		err := db.AddPoint(&statev1.PointAndState{
			Id:    string(point),
			State: statev1.PointStateEnum_POINT_STATE_NORMAL,
		})
		if err != nil {
			println("Error adding point rail", point, ":", err.Error())
			// エラーが発生しても処理を続行
		}
	}
	for _, block := range data.Blocks {
		println("Adding block:", block)
		err := db.AddBlock(&statev1.BlockState{
			BlockId: string(block),
			State:   statev1.BlockStateEnum_BLOCK_STATE_OPEN,
		})
		if err != nil {
			println("Error adding block", block, ":", err.Error())
			// エラーが発生しても処理を続行
		}
	}
	
	println("Seed data import completed")
}
