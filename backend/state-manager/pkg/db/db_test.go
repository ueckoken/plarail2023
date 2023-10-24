package main

import (
	"github.com/joho/godotenv"
	statev1 "github.com/ueckoken/plarail2023/backend/spec/state/v1"
	"testing"
)

func Test_ConnectDB(t *testing.T) {
	defer c()
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	open()
}

func Test_SetPoint(t *testing.T) {
	defer c()
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	open()
	setPoint(&statev1.PointAndState{
		Id:    "test",
		State: statev1.PointStateEnum_POINT_STATE_REVERSE,
	})
	point := getPoint("test")
	if point.State != statev1.PointStateEnum_POINT_STATE_REVERSE {
		t.Fatal("point state is not reverse")
	}
}

func Test_SetStop(t *testing.T) {
	defer c()
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	open()
	setStop(&statev1.StopAndState{
		Id:    "test",
		State: statev1.StopStateEnum_STOP_STATE_GO,
	})
	stop := getStop("test")
	if stop.State != statev1.StopStateEnum_STOP_STATE_GO {
		t.Fatal("point state is not stop")
	}
}

func Test_GetPoints(t *testing.T) {
	defer c()
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	open()
	setPoint(&statev1.PointAndState{
		Id:    "test",
		State: statev1.PointStateEnum_POINT_STATE_REVERSE,
	})
	setPoint(&statev1.PointAndState{
		Id:    "test2",
		State: statev1.PointStateEnum_POINT_STATE_NORMAL,
	})
	points := getPoints()
	if len(points) < 2 {
		t.Fatal("points length is not larger than 2")
	}
}

func Test_GetStops(t *testing.T) {
	defer c()
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	open()
	setStop(&statev1.StopAndState{
		Id:    "test",
		State: statev1.StopStateEnum_STOP_STATE_GO,
	})
	setStop(&statev1.StopAndState{
		Id:    "test2",
		State: statev1.StopStateEnum_STOP_STATE_STOP,
	})
	stops := getStops()
	if len(stops) < 2 {
		t.Fatal("stops length is not larger than 2")
	}
}
