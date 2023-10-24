// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: state/v1/state.proto

package statev1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/ueckoken/plarail2023/backend/spec/state/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// StateManagerServiceName is the fully-qualified name of the StateManagerService service.
	StateManagerServiceName = "state.v1.StateManagerService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// StateManagerServiceGetBlockStatesProcedure is the fully-qualified name of the
	// StateManagerService's GetBlockStates RPC.
	StateManagerServiceGetBlockStatesProcedure = "/state.v1.StateManagerService/GetBlockStates"
	// StateManagerServiceUpdateBLockStatesProcedure is the fully-qualified name of the
	// StateManagerService's UpdateBLockStates RPC.
	StateManagerServiceUpdateBLockStatesProcedure = "/state.v1.StateManagerService/UpdateBLockStates"
	// StateManagerServiceUpdatePointStateProcedure is the fully-qualified name of the
	// StateManagerService's UpdatePointState RPC.
	StateManagerServiceUpdatePointStateProcedure = "/state.v1.StateManagerService/UpdatePointState"
	// StateManagerServiceGetPointStatesProcedure is the fully-qualified name of the
	// StateManagerService's GetPointStates RPC.
	StateManagerServiceGetPointStatesProcedure = "/state.v1.StateManagerService/GetPointStates"
	// StateManagerServiceUpdateStopStateProcedure is the fully-qualified name of the
	// StateManagerService's UpdateStopState RPC.
	StateManagerServiceUpdateStopStateProcedure = "/state.v1.StateManagerService/UpdateStopState"
	// StateManagerServiceGetStopStatesProcedure is the fully-qualified name of the
	// StateManagerService's GetStopStates RPC.
	StateManagerServiceGetStopStatesProcedure = "/state.v1.StateManagerService/GetStopStates"
	// StateManagerServiceGetTrainsProcedure is the fully-qualified name of the StateManagerService's
	// GetTrains RPC.
	StateManagerServiceGetTrainsProcedure = "/state.v1.StateManagerService/GetTrains"
	// StateManagerServiceUpdateTrainUUIDProcedure is the fully-qualified name of the
	// StateManagerService's UpdateTrainUUID RPC.
	StateManagerServiceUpdateTrainUUIDProcedure = "/state.v1.StateManagerService/UpdateTrainUUID"
)

// StateManagerServiceClient is a client for the state.v1.StateManagerService service.
type StateManagerServiceClient interface {
	// Block
	GetBlockStates(context.Context, *connect_go.Request[v1.GetBlockStatesRequest]) (*connect_go.Response[v1.GetBlockStatesResponse], error)
	UpdateBLockStates(context.Context, *connect_go.Request[v1.UpdateBlockStateRequest]) (*connect_go.Response[v1.UpdateBlockStateResponse], error)
	// Point
	UpdatePointState(context.Context, *connect_go.Request[v1.UpdatePointStateRequest]) (*connect_go.Response[v1.UpdatePointStateResponse], error)
	GetPointStates(context.Context, *connect_go.Request[v1.GetPointStatesRequest]) (*connect_go.Response[v1.GetPointStatesResponse], error)
	// Stop
	UpdateStopState(context.Context, *connect_go.Request[v1.UpdateStopStateRequest]) (*connect_go.Response[v1.UpdateStopStateResponse], error)
	GetStopStates(context.Context, *connect_go.Request[v1.GetStopStatesRequest]) (*connect_go.Response[v1.GetStopStatesResponse], error)
	// Train
	GetTrains(context.Context, *connect_go.Request[v1.GetTrainsRequest]) (*connect_go.Response[v1.GetTrainsResponse], error)
	UpdateTrainUUID(context.Context, *connect_go.Request[v1.UpdateTrainUUIDRequest]) (*connect_go.Response[v1.UpdateTrainUUIDResponse], error)
}

// NewStateManagerServiceClient constructs a client for the state.v1.StateManagerService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewStateManagerServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) StateManagerServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &stateManagerServiceClient{
		getBlockStates: connect_go.NewClient[v1.GetBlockStatesRequest, v1.GetBlockStatesResponse](
			httpClient,
			baseURL+StateManagerServiceGetBlockStatesProcedure,
			opts...,
		),
		updateBLockStates: connect_go.NewClient[v1.UpdateBlockStateRequest, v1.UpdateBlockStateResponse](
			httpClient,
			baseURL+StateManagerServiceUpdateBLockStatesProcedure,
			opts...,
		),
		updatePointState: connect_go.NewClient[v1.UpdatePointStateRequest, v1.UpdatePointStateResponse](
			httpClient,
			baseURL+StateManagerServiceUpdatePointStateProcedure,
			opts...,
		),
		getPointStates: connect_go.NewClient[v1.GetPointStatesRequest, v1.GetPointStatesResponse](
			httpClient,
			baseURL+StateManagerServiceGetPointStatesProcedure,
			opts...,
		),
		updateStopState: connect_go.NewClient[v1.UpdateStopStateRequest, v1.UpdateStopStateResponse](
			httpClient,
			baseURL+StateManagerServiceUpdateStopStateProcedure,
			opts...,
		),
		getStopStates: connect_go.NewClient[v1.GetStopStatesRequest, v1.GetStopStatesResponse](
			httpClient,
			baseURL+StateManagerServiceGetStopStatesProcedure,
			opts...,
		),
		getTrains: connect_go.NewClient[v1.GetTrainsRequest, v1.GetTrainsResponse](
			httpClient,
			baseURL+StateManagerServiceGetTrainsProcedure,
			opts...,
		),
		updateTrainUUID: connect_go.NewClient[v1.UpdateTrainUUIDRequest, v1.UpdateTrainUUIDResponse](
			httpClient,
			baseURL+StateManagerServiceUpdateTrainUUIDProcedure,
			opts...,
		),
	}
}

// stateManagerServiceClient implements StateManagerServiceClient.
type stateManagerServiceClient struct {
	getBlockStates    *connect_go.Client[v1.GetBlockStatesRequest, v1.GetBlockStatesResponse]
	updateBLockStates *connect_go.Client[v1.UpdateBlockStateRequest, v1.UpdateBlockStateResponse]
	updatePointState  *connect_go.Client[v1.UpdatePointStateRequest, v1.UpdatePointStateResponse]
	getPointStates    *connect_go.Client[v1.GetPointStatesRequest, v1.GetPointStatesResponse]
	updateStopState   *connect_go.Client[v1.UpdateStopStateRequest, v1.UpdateStopStateResponse]
	getStopStates     *connect_go.Client[v1.GetStopStatesRequest, v1.GetStopStatesResponse]
	getTrains         *connect_go.Client[v1.GetTrainsRequest, v1.GetTrainsResponse]
	updateTrainUUID   *connect_go.Client[v1.UpdateTrainUUIDRequest, v1.UpdateTrainUUIDResponse]
}

// GetBlockStates calls state.v1.StateManagerService.GetBlockStates.
func (c *stateManagerServiceClient) GetBlockStates(ctx context.Context, req *connect_go.Request[v1.GetBlockStatesRequest]) (*connect_go.Response[v1.GetBlockStatesResponse], error) {
	return c.getBlockStates.CallUnary(ctx, req)
}

// UpdateBLockStates calls state.v1.StateManagerService.UpdateBLockStates.
func (c *stateManagerServiceClient) UpdateBLockStates(ctx context.Context, req *connect_go.Request[v1.UpdateBlockStateRequest]) (*connect_go.Response[v1.UpdateBlockStateResponse], error) {
	return c.updateBLockStates.CallUnary(ctx, req)
}

// UpdatePointState calls state.v1.StateManagerService.UpdatePointState.
func (c *stateManagerServiceClient) UpdatePointState(ctx context.Context, req *connect_go.Request[v1.UpdatePointStateRequest]) (*connect_go.Response[v1.UpdatePointStateResponse], error) {
	return c.updatePointState.CallUnary(ctx, req)
}

// GetPointStates calls state.v1.StateManagerService.GetPointStates.
func (c *stateManagerServiceClient) GetPointStates(ctx context.Context, req *connect_go.Request[v1.GetPointStatesRequest]) (*connect_go.Response[v1.GetPointStatesResponse], error) {
	return c.getPointStates.CallUnary(ctx, req)
}

// UpdateStopState calls state.v1.StateManagerService.UpdateStopState.
func (c *stateManagerServiceClient) UpdateStopState(ctx context.Context, req *connect_go.Request[v1.UpdateStopStateRequest]) (*connect_go.Response[v1.UpdateStopStateResponse], error) {
	return c.updateStopState.CallUnary(ctx, req)
}

// GetStopStates calls state.v1.StateManagerService.GetStopStates.
func (c *stateManagerServiceClient) GetStopStates(ctx context.Context, req *connect_go.Request[v1.GetStopStatesRequest]) (*connect_go.Response[v1.GetStopStatesResponse], error) {
	return c.getStopStates.CallUnary(ctx, req)
}

// GetTrains calls state.v1.StateManagerService.GetTrains.
func (c *stateManagerServiceClient) GetTrains(ctx context.Context, req *connect_go.Request[v1.GetTrainsRequest]) (*connect_go.Response[v1.GetTrainsResponse], error) {
	return c.getTrains.CallUnary(ctx, req)
}

// UpdateTrainUUID calls state.v1.StateManagerService.UpdateTrainUUID.
func (c *stateManagerServiceClient) UpdateTrainUUID(ctx context.Context, req *connect_go.Request[v1.UpdateTrainUUIDRequest]) (*connect_go.Response[v1.UpdateTrainUUIDResponse], error) {
	return c.updateTrainUUID.CallUnary(ctx, req)
}

// StateManagerServiceHandler is an implementation of the state.v1.StateManagerService service.
type StateManagerServiceHandler interface {
	// Block
	GetBlockStates(context.Context, *connect_go.Request[v1.GetBlockStatesRequest]) (*connect_go.Response[v1.GetBlockStatesResponse], error)
	UpdateBLockStates(context.Context, *connect_go.Request[v1.UpdateBlockStateRequest]) (*connect_go.Response[v1.UpdateBlockStateResponse], error)
	// Point
	UpdatePointState(context.Context, *connect_go.Request[v1.UpdatePointStateRequest]) (*connect_go.Response[v1.UpdatePointStateResponse], error)
	GetPointStates(context.Context, *connect_go.Request[v1.GetPointStatesRequest]) (*connect_go.Response[v1.GetPointStatesResponse], error)
	// Stop
	UpdateStopState(context.Context, *connect_go.Request[v1.UpdateStopStateRequest]) (*connect_go.Response[v1.UpdateStopStateResponse], error)
	GetStopStates(context.Context, *connect_go.Request[v1.GetStopStatesRequest]) (*connect_go.Response[v1.GetStopStatesResponse], error)
	// Train
	GetTrains(context.Context, *connect_go.Request[v1.GetTrainsRequest]) (*connect_go.Response[v1.GetTrainsResponse], error)
	UpdateTrainUUID(context.Context, *connect_go.Request[v1.UpdateTrainUUIDRequest]) (*connect_go.Response[v1.UpdateTrainUUIDResponse], error)
}

// NewStateManagerServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewStateManagerServiceHandler(svc StateManagerServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	stateManagerServiceGetBlockStatesHandler := connect_go.NewUnaryHandler(
		StateManagerServiceGetBlockStatesProcedure,
		svc.GetBlockStates,
		opts...,
	)
	stateManagerServiceUpdateBLockStatesHandler := connect_go.NewUnaryHandler(
		StateManagerServiceUpdateBLockStatesProcedure,
		svc.UpdateBLockStates,
		opts...,
	)
	stateManagerServiceUpdatePointStateHandler := connect_go.NewUnaryHandler(
		StateManagerServiceUpdatePointStateProcedure,
		svc.UpdatePointState,
		opts...,
	)
	stateManagerServiceGetPointStatesHandler := connect_go.NewUnaryHandler(
		StateManagerServiceGetPointStatesProcedure,
		svc.GetPointStates,
		opts...,
	)
	stateManagerServiceUpdateStopStateHandler := connect_go.NewUnaryHandler(
		StateManagerServiceUpdateStopStateProcedure,
		svc.UpdateStopState,
		opts...,
	)
	stateManagerServiceGetStopStatesHandler := connect_go.NewUnaryHandler(
		StateManagerServiceGetStopStatesProcedure,
		svc.GetStopStates,
		opts...,
	)
	stateManagerServiceGetTrainsHandler := connect_go.NewUnaryHandler(
		StateManagerServiceGetTrainsProcedure,
		svc.GetTrains,
		opts...,
	)
	stateManagerServiceUpdateTrainUUIDHandler := connect_go.NewUnaryHandler(
		StateManagerServiceUpdateTrainUUIDProcedure,
		svc.UpdateTrainUUID,
		opts...,
	)
	return "/state.v1.StateManagerService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case StateManagerServiceGetBlockStatesProcedure:
			stateManagerServiceGetBlockStatesHandler.ServeHTTP(w, r)
		case StateManagerServiceUpdateBLockStatesProcedure:
			stateManagerServiceUpdateBLockStatesHandler.ServeHTTP(w, r)
		case StateManagerServiceUpdatePointStateProcedure:
			stateManagerServiceUpdatePointStateHandler.ServeHTTP(w, r)
		case StateManagerServiceGetPointStatesProcedure:
			stateManagerServiceGetPointStatesHandler.ServeHTTP(w, r)
		case StateManagerServiceUpdateStopStateProcedure:
			stateManagerServiceUpdateStopStateHandler.ServeHTTP(w, r)
		case StateManagerServiceGetStopStatesProcedure:
			stateManagerServiceGetStopStatesHandler.ServeHTTP(w, r)
		case StateManagerServiceGetTrainsProcedure:
			stateManagerServiceGetTrainsHandler.ServeHTTP(w, r)
		case StateManagerServiceUpdateTrainUUIDProcedure:
			stateManagerServiceUpdateTrainUUIDHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedStateManagerServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedStateManagerServiceHandler struct{}

func (UnimplementedStateManagerServiceHandler) GetBlockStates(context.Context, *connect_go.Request[v1.GetBlockStatesRequest]) (*connect_go.Response[v1.GetBlockStatesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("state.v1.StateManagerService.GetBlockStates is not implemented"))
}

func (UnimplementedStateManagerServiceHandler) UpdateBLockStates(context.Context, *connect_go.Request[v1.UpdateBlockStateRequest]) (*connect_go.Response[v1.UpdateBlockStateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("state.v1.StateManagerService.UpdateBLockStates is not implemented"))
}

func (UnimplementedStateManagerServiceHandler) UpdatePointState(context.Context, *connect_go.Request[v1.UpdatePointStateRequest]) (*connect_go.Response[v1.UpdatePointStateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("state.v1.StateManagerService.UpdatePointState is not implemented"))
}

func (UnimplementedStateManagerServiceHandler) GetPointStates(context.Context, *connect_go.Request[v1.GetPointStatesRequest]) (*connect_go.Response[v1.GetPointStatesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("state.v1.StateManagerService.GetPointStates is not implemented"))
}

func (UnimplementedStateManagerServiceHandler) UpdateStopState(context.Context, *connect_go.Request[v1.UpdateStopStateRequest]) (*connect_go.Response[v1.UpdateStopStateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("state.v1.StateManagerService.UpdateStopState is not implemented"))
}

func (UnimplementedStateManagerServiceHandler) GetStopStates(context.Context, *connect_go.Request[v1.GetStopStatesRequest]) (*connect_go.Response[v1.GetStopStatesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("state.v1.StateManagerService.GetStopStates is not implemented"))
}

func (UnimplementedStateManagerServiceHandler) GetTrains(context.Context, *connect_go.Request[v1.GetTrainsRequest]) (*connect_go.Response[v1.GetTrainsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("state.v1.StateManagerService.GetTrains is not implemented"))
}

func (UnimplementedStateManagerServiceHandler) UpdateTrainUUID(context.Context, *connect_go.Request[v1.UpdateTrainUUIDRequest]) (*connect_go.Response[v1.UpdateTrainUUIDResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("state.v1.StateManagerService.UpdateTrainUUID is not implemented"))
}
