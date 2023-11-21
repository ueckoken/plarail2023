package connect_handler

import (
	"context"
	"errors"
	"log/slog"

	connect "github.com/bufbuild/connect-go"

	statev1 "github.com/ueckoken/plarail2023/backend/spec/state/v1"
	"github.com/ueckoken/plarail2023/backend/state-manager/pkg/db"
	"github.com/ueckoken/plarail2023/backend/state-manager/pkg/mqtt_handler"
)

type StateManagerServer struct {
	DBHandler   *db.DBHandler
	MqttHandler *mqtt_handler.Handler
}

/*
Block
*/

// GetBlockStates Blockの状態を取得する
func (s *StateManagerServer) GetBlockStates(
	ctx context.Context,
	req *connect.Request[statev1.GetBlockStatesRequest],
) (*connect.Response[statev1.GetBlockStatesResponse], error) {
	blockStates, err := s.DBHandler.GetBlocks()
	if err != nil {
		err = connect.NewError(
			connect.CodeUnknown,
			errors.New("db error"),
		)
		return nil, err
	}

	var response []*statev1.BlockState

	for _, blockState := range blockStates {
		response = append(response, &statev1.BlockState{
			BlockId: blockState.BlockId,
			State:   blockState.State,
		})
	}

	res := connect.NewResponse(&statev1.GetBlockStatesResponse{
		States: response,
	})

	return res, nil
}

// UpdateBlockState Blockの状態を更新する
func (s *StateManagerServer) UpdateBlockState(
	ctx context.Context,
	req *connect.Request[statev1.UpdateBlockStateRequest],
) (*connect.Response[statev1.UpdateBlockStateResponse], error) {
	err := s.DBHandler.UpdateBlock(req.Msg.State)
	if err != nil {
		err = connect.NewError(
			connect.CodeUnknown,
			errors.New("db error"),
		)
		return nil, err
	}
	res := connect.NewResponse(&statev1.UpdateBlockStateResponse{})
	return res, nil
}

/*
Point
*/

func (s *StateManagerServer) UpdatePointState(
	ctx context.Context,
	req *connect.Request[statev1.UpdatePointStateRequest],
) (*connect.Response[statev1.UpdatePointStateResponse], error) {
	err := s.DBHandler.UpdatePoint(req.Msg.State)
	if err != nil {
		err = connect.NewError(
			connect.CodeUnknown,
			errors.New("db error"),
		)
		slog.Default().Error("db error", err)
		return nil, err
	}
	s.MqttHandler.NotifyStateUpdate("point", req.Msg.State.Id, req.Msg.State.State.String())

	return connect.NewResponse(&statev1.UpdatePointStateResponse{}), nil
}

func (s *StateManagerServer) GetPointStates(
	ctx context.Context,
	req *connect.Request[statev1.GetPointStatesRequest],
) (*connect.Response[statev1.GetPointStatesResponse], error) {
	err := connect.NewError(
		connect.CodeUnknown,
		errors.New("not implemented"),
	)
	return nil, err
}

/*
Stop
*/

func (s *StateManagerServer) UpdateStopState(
	ctx context.Context,
	req *connect.Request[statev1.UpdateStopStateRequest],
) (*connect.Response[statev1.UpdateStopStateResponse], error) {
	err := s.DBHandler.UpdateStop(req.Msg.State)
	if err != nil {
		err = connect.NewError(
			connect.CodeUnknown,
			errors.New("db error"),
		)
		slog.Default().Error("db connection error", err)
		return nil, err
	}
	s.MqttHandler.NotifyStateUpdate("stop", req.Msg.State.Id, req.Msg.State.State.String())
	return connect.NewResponse(&statev1.UpdateStopStateResponse{}), nil
}

func (s *StateManagerServer) GetStopStates(
	ctx context.Context,
	req *connect.Request[statev1.GetStopStatesRequest],
) (*connect.Response[statev1.GetStopStatesResponse], error) {
	err := connect.NewError(
		connect.CodeUnknown,
		errors.New("not implemented"),
	)
	return nil, err
}

/*
Train
*/

func (s *StateManagerServer) GetTrains(
	ctx context.Context,
	req *connect.Request[statev1.GetTrainsRequest],
) (*connect.Response[statev1.GetTrainsResponse], error) {
	err := connect.NewError(
		connect.CodeUnknown,
		errors.New("not implemented"),
	)
	return nil, err
}

func (s *StateManagerServer) UpdateTrainUUID(
	ctx context.Context,
	req *connect.Request[statev1.UpdateTrainUUIDRequest],
) (*connect.Response[statev1.UpdateTrainUUIDResponse], error) {
	err := connect.NewError(
		connect.CodeUnknown,
		errors.New("not implemented"),
	)
	return nil, err
}
