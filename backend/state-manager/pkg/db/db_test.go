package db

import (
	"testing"

	"github.com/joho/godotenv"
	statev1 "github.com/ueckoken/plarail2023/backend/spec/state/v1"
)

func Test_ConnectDB(t *testing.T) {
	defer C()
	err := godotenv.Load("../../cmd/.env")
	if err != nil {
		panic(err)
	}
	Open()
}

func Test_SetPoint(t *testing.T) {
	defer C()
	err := godotenv.Load("../../cmd/.env")
	if err != nil {
		panic(err)
	}
	Open()
	SetPoint(&statev1.PointAndState{
		Id:    "test",
		State: statev1.PointStateEnum_POINT_STATE_REVERSE,
	})
	point, err := GetPoint("test")
	if err != nil {
		t.Fatal("error")
	}
	if point.State != statev1.PointStateEnum_POINT_STATE_REVERSE {
		t.Fatal("point state is not reverse")
	}
}

func Test_SetStop(t *testing.T) {
	defer C()
	err := godotenv.Load("../../cmd/.env")
	if err != nil {
		panic(err)
	}
	Open()
	SetStop(&statev1.StopAndState{
		Id:    "test",
		State: statev1.StopStateEnum_STOP_STATE_GO,
	})
	stop, err := GetStop("test")
	if err != nil {
		t.Fatal("error")
	}
	if stop.State != statev1.StopStateEnum_STOP_STATE_GO {
		t.Fatal("point state is not stop")
	}
}

func Test_GetPoints(t *testing.T) {
	defer C()
	err := godotenv.Load("../../cmd/.env")
	if err != nil {
		panic(err)
	}
	Open()
	SetPoint(&statev1.PointAndState{
		Id:    "test",
		State: statev1.PointStateEnum_POINT_STATE_REVERSE,
	})
	SetPoint(&statev1.PointAndState{
		Id:    "test2",
		State: statev1.PointStateEnum_POINT_STATE_NORMAL,
	})
	points := GetPoints()
	if len(points) < 2 {
		t.Fatal("points length is not larger than 2")
	}
}

func Test_GetStops(t *testing.T) {
	defer C()
	err := godotenv.Load("../../cmd/.env")
	if err != nil {
		panic(err)
	}
	Open()
	SetStop(&statev1.StopAndState{
		Id:    "test",
		State: statev1.StopStateEnum_STOP_STATE_GO,
	})
	SetStop(&statev1.StopAndState{
		Id:    "test2",
		State: statev1.StopStateEnum_STOP_STATE_STOP,
	})
	stops := GetStops()
	if len(stops) < 2 {
		t.Fatal("stops length is not larger than 2")
	}
}

func Test_GetBlocks(t *testing.T) {
	defer C()
	err := godotenv.Load("../../cmd/.env")
	if err != nil {
		panic(err)
	}
	Open()
	err = SetBlock(&statev1.BlockState{
		BlockId: "test",
		State:   statev1.BlockStateEnum_BLOCK_STATE_OPEN,
	})
	if err != nil {
		t.Fatal("error")
	}
	err = SetBlock(&statev1.BlockState{
		BlockId: "test2",
		State:   statev1.BlockStateEnum_BLOCK_STATE_CLOSE,
	})
	if err != nil {
		t.Fatal("error")
	}
	block, err := GetBlock("test")
	if err != nil {
		t.Fatal("error")
	}
	if block.State != statev1.BlockStateEnum_BLOCK_STATE_OPEN {
		t.Fatal("block state is not open")
	}
	blocks, err := GetBlocks()
	if err != nil {
		t.Fatal("error")
	}
	if len(blocks) < 2 {
		t.Fatal("blocks length is not larger than 2")
	}
}
