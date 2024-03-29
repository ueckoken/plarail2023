//
//Stop Proto
//ストップレールの状態を扱うプロトコル

// @generated by protoc-gen-es v1.4.2 with parameter "target=dts+js"
// @generated from file state/v1/stop.proto (package state.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum state.v1.StopStateEnum
 */
export const StopStateEnum = proto3.makeEnum(
  "state.v1.StopStateEnum",
  [
    {no: 0, name: "STOP_STATE_UNKNOWN"},
    {no: 1, name: "STOP_STATE_GO"},
    {no: 2, name: "STOP_STATE_STOP"},
  ],
);

/**
 * @generated from message state.v1.StopAndState
 */
export const StopAndState = proto3.makeMessageType(
  "state.v1.StopAndState",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "state", kind: "enum", T: proto3.getEnumType(StopStateEnum) },
  ],
);

/**
 *
 * UpdateStopState : ストップレールの状態を更新するAPI
 *
 * @generated from message state.v1.UpdateStopStateRequest
 */
export const UpdateStopStateRequest = proto3.makeMessageType(
  "state.v1.UpdateStopStateRequest",
  () => [
    { no: 1, name: "state", kind: "message", T: StopAndState },
  ],
);

/**
 * @generated from message state.v1.UpdateStopStateResponse
 */
export const UpdateStopStateResponse = proto3.makeMessageType(
  "state.v1.UpdateStopStateResponse",
  [],
);

/**
 *
 * GetStopStates : 全てのストップレールの状態を取得するAPI
 *
 * @generated from message state.v1.GetStopStatesRequest
 */
export const GetStopStatesRequest = proto3.makeMessageType(
  "state.v1.GetStopStatesRequest",
  [],
);

/**
 * @generated from message state.v1.GetStopStatesResponse
 */
export const GetStopStatesResponse = proto3.makeMessageType(
  "state.v1.GetStopStatesResponse",
  () => [
    { no: 1, name: "states", kind: "message", T: StopAndState, repeated: true },
  ],
);

