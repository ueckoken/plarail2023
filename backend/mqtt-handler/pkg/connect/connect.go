package connect

import (
	"context"

	"github.com/bufbuild/connect-go"
	trainv1 "github.com/ueckoken/plarail2023/backend/proto/train/v1"
)

type TrainServer struct {
}

func (s *TrainServer) NotifyTrainArrival(
	ctx context.Context,
	req *connect.Request[trainv1.NotifyTrainArrivalRequest],
) (*connect.Response[trainv1.NotifyTrainArrivalResponse], error) {
	// handler.HandleConnectMessage(req.Target, req.Message)

	res := connect.NewResponse(&trainv1.NotifyTrainArrivalResponse{
		Success: true,
	})
	return res, nil
}
