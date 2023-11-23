//
//Block Proto
//閉塞の状態を扱うプロトコル

// @generated by protoc-gen-es v1.4.2 with parameter "target=dts+js"
// @generated from file state/v1/block.proto (package state.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum state.v1.BlockStateEnum
 */
export const BlockStateEnum = proto3.makeEnum(
  "state.v1.BlockStateEnum",
  [
    {no: 0, name: "BLOCK_STATE_UNKNOWN"},
    {no: 1, name: "BLOCK_STATE_OPEN"},
    {no: 2, name: "BLOCK_STATE_CLOSE"},
  ],
);

/**
 * 閉塞の状態
 *
 * @generated from message state.v1.BlockState
 */
export const BlockState = proto3.makeMessageType(
  "state.v1.BlockState",
  () => [
    { no: 1, name: "block_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "state", kind: "enum", T: proto3.getEnumType(BlockStateEnum) },
  ],
);

/**
 *
 * GetBlockStates : 閉塞の状態を取得するAPI
 *
 * @generated from message state.v1.GetBlockStatesRequest
 */
export const GetBlockStatesRequest = proto3.makeMessageType(
  "state.v1.GetBlockStatesRequest",
  [],
);

/**
 * @generated from message state.v1.GetBlockStatesResponse
 */
export const GetBlockStatesResponse = proto3.makeMessageType(
  "state.v1.GetBlockStatesResponse",
  () => [
    { no: 1, name: "states", kind: "message", T: BlockState, repeated: true },
  ],
);

/**
 *
 * UpdateBLockState: 閉塞の状態を更新するAPI
 *
 * @generated from message state.v1.UpdateBlockStateRequest
 */
export const UpdateBlockStateRequest = proto3.makeMessageType(
  "state.v1.UpdateBlockStateRequest",
  () => [
    { no: 1, name: "state", kind: "message", T: BlockState },
  ],
);

/**
 * @generated from message state.v1.UpdateBlockStateResponse
 */
export const UpdateBlockStateResponse = proto3.makeMessageType(
  "state.v1.UpdateBlockStateResponse",
  [],
);
