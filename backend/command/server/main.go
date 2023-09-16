package main

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	trainv1 "github.com/ueckoken/plarail2023/backend/proto/train/v1"
	"github.com/ueckoken/plarail2023/backend/proto/train/v1/trainv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type TrainServer struct{}

func (s *TrainServer) NotifyTrainArrival(
	ctx context.Context,
	req *connect.Request[trainv1.NotifyTrainArrivalRequest],
) (*connect.Response[trainv1.NotifyTrainArrivalResponse], error) {
	res := connect.NewResponse(&trainv1.NotifyTrainArrivalResponse{
		Success: true,
	})
	return res, nil
}

func main() {
	server := &TrainServer{}
	mux := http.NewServeMux()
	path, handler := trainv1connect.NewTrainServiceHandler(server)
	mux.Handle(path, handler)
	http.ListenAndServe("localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
