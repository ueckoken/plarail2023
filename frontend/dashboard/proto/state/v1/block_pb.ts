//
//Block Proto
//閉塞の状態を扱うプロトコル

// @generated by protoc-gen-es v1.4.2 with parameter "target=ts"
// @generated from file state/v1/block.proto (package state.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum state.v1.BlockStateEnum
 */
export enum BlockStateEnum {
  /**
   * @generated from enum value: BLOCK_STATE_UNKNOWN = 0;
   */
  BLOCK_STATE_UNKNOWN = 0,

  /**
   * 閉塞が開の状態(列車がいない)
   *
   * @generated from enum value: BLOCK_STATE_OPEN = 1;
   */
  BLOCK_STATE_OPEN = 1,

  /**
   * 閉塞が閉の状態(列車がいない)
   *
   * @generated from enum value: BLOCK_STATE_CLOSE = 2;
   */
  BLOCK_STATE_CLOSE = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(BlockStateEnum)
proto3.util.setEnumType(BlockStateEnum, "state.v1.BlockStateEnum", [
  { no: 0, name: "BLOCK_STATE_UNKNOWN" },
  { no: 1, name: "BLOCK_STATE_OPEN" },
  { no: 2, name: "BLOCK_STATE_CLOSE" },
]);

/**
 * 閉塞の状態
 *
 * @generated from message state.v1.BlockState
 */
export class BlockState extends Message<BlockState> {
  /**
   * 閉塞のID
   *
   * @generated from field: string block_id = 1;
   */
  blockId = "";

  /**
   * 閉塞の状態
   *
   * @generated from field: state.v1.BlockStateEnum state = 2;
   */
  state = BlockStateEnum.BLOCK_STATE_UNKNOWN;

  constructor(data?: PartialMessage<BlockState>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "state.v1.BlockState";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "block_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "state", kind: "enum", T: proto3.getEnumType(BlockStateEnum) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BlockState {
    return new BlockState().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BlockState {
    return new BlockState().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BlockState {
    return new BlockState().fromJsonString(jsonString, options);
  }

  static equals(a: BlockState | PlainMessage<BlockState> | undefined, b: BlockState | PlainMessage<BlockState> | undefined): boolean {
    return proto3.util.equals(BlockState, a, b);
  }
}

/**
 *
 * GetBlockStates : 閉塞の状態を取得するAPI
 *
 * @generated from message state.v1.GetBlockStatesRequest
 */
export class GetBlockStatesRequest extends Message<GetBlockStatesRequest> {
  constructor(data?: PartialMessage<GetBlockStatesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "state.v1.GetBlockStatesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetBlockStatesRequest {
    return new GetBlockStatesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetBlockStatesRequest {
    return new GetBlockStatesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetBlockStatesRequest {
    return new GetBlockStatesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetBlockStatesRequest | PlainMessage<GetBlockStatesRequest> | undefined, b: GetBlockStatesRequest | PlainMessage<GetBlockStatesRequest> | undefined): boolean {
    return proto3.util.equals(GetBlockStatesRequest, a, b);
  }
}

/**
 * @generated from message state.v1.GetBlockStatesResponse
 */
export class GetBlockStatesResponse extends Message<GetBlockStatesResponse> {
  /**
   * @generated from field: repeated state.v1.BlockState states = 1;
   */
  states: BlockState[] = [];

  constructor(data?: PartialMessage<GetBlockStatesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "state.v1.GetBlockStatesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "states", kind: "message", T: BlockState, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetBlockStatesResponse {
    return new GetBlockStatesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetBlockStatesResponse {
    return new GetBlockStatesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetBlockStatesResponse {
    return new GetBlockStatesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetBlockStatesResponse | PlainMessage<GetBlockStatesResponse> | undefined, b: GetBlockStatesResponse | PlainMessage<GetBlockStatesResponse> | undefined): boolean {
    return proto3.util.equals(GetBlockStatesResponse, a, b);
  }
}

/**
 *
 * UpdateBLockState: 閉塞の状態を更新するAPI
 *
 * @generated from message state.v1.UpdateBlockStateRequest
 */
export class UpdateBlockStateRequest extends Message<UpdateBlockStateRequest> {
  /**
   * @generated from field: state.v1.BlockState state = 1;
   */
  state?: BlockState;

  constructor(data?: PartialMessage<UpdateBlockStateRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "state.v1.UpdateBlockStateRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "state", kind: "message", T: BlockState },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateBlockStateRequest {
    return new UpdateBlockStateRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateBlockStateRequest {
    return new UpdateBlockStateRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateBlockStateRequest {
    return new UpdateBlockStateRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateBlockStateRequest | PlainMessage<UpdateBlockStateRequest> | undefined, b: UpdateBlockStateRequest | PlainMessage<UpdateBlockStateRequest> | undefined): boolean {
    return proto3.util.equals(UpdateBlockStateRequest, a, b);
  }
}

/**
 * @generated from message state.v1.UpdateBlockStateResponse
 */
export class UpdateBlockStateResponse extends Message<UpdateBlockStateResponse> {
  constructor(data?: PartialMessage<UpdateBlockStateResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "state.v1.UpdateBlockStateResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateBlockStateResponse {
    return new UpdateBlockStateResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateBlockStateResponse {
    return new UpdateBlockStateResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateBlockStateResponse {
    return new UpdateBlockStateResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateBlockStateResponse | PlainMessage<UpdateBlockStateResponse> | undefined, b: UpdateBlockStateResponse | PlainMessage<UpdateBlockStateResponse> | undefined): boolean {
    return proto3.util.equals(UpdateBlockStateResponse, a, b);
  }
}

