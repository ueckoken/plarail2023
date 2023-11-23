//
//Stop Proto
//ストップレールの状態を扱うプロトコル

// @generated by protoc-gen-es v1.4.2 with parameter "target=ts"
// @generated from file state/v1/stop.proto (package state.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum state.v1.StopStateEnum
 */
export enum StopStateEnum {
  /**
   * @generated from enum value: STOP_STATE_UNKNOWN = 0;
   */
  STOP_STATE_UNKNOWN = 0,

  /**
   * バーが下がってる状態
   *
   * @generated from enum value: STOP_STATE_GO = 1;
   */
  STOP_STATE_GO = 1,

  /**
   * バーが上がってる状態
   *
   * @generated from enum value: STOP_STATE_STOP = 2;
   */
  STOP_STATE_STOP = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(StopStateEnum)
proto3.util.setEnumType(StopStateEnum, "state.v1.StopStateEnum", [
  { no: 0, name: "STOP_STATE_UNKNOWN" },
  { no: 1, name: "STOP_STATE_GO" },
  { no: 2, name: "STOP_STATE_STOP" },
]);

/**
 * @generated from message state.v1.StopAndState
 */
export class StopAndState extends Message<StopAndState> {
  /**
   * ポイントのid
   *
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * ポイントの状態
   *
   * @generated from field: state.v1.StopStateEnum state = 2;
   */
  state = StopStateEnum.STOP_STATE_UNKNOWN;

  constructor(data?: PartialMessage<StopAndState>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "state.v1.StopAndState";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "state", kind: "enum", T: proto3.getEnumType(StopStateEnum) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StopAndState {
    return new StopAndState().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StopAndState {
    return new StopAndState().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StopAndState {
    return new StopAndState().fromJsonString(jsonString, options);
  }

  static equals(a: StopAndState | PlainMessage<StopAndState> | undefined, b: StopAndState | PlainMessage<StopAndState> | undefined): boolean {
    return proto3.util.equals(StopAndState, a, b);
  }
}

/**
 *
 * UpdateStopState : ストップレールの状態を更新するAPI
 *
 * @generated from message state.v1.UpdateStopStateRequest
 */
export class UpdateStopStateRequest extends Message<UpdateStopStateRequest> {
  /**
   * @generated from field: state.v1.StopAndState state = 1;
   */
  state?: StopAndState;

  constructor(data?: PartialMessage<UpdateStopStateRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "state.v1.UpdateStopStateRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "state", kind: "message", T: StopAndState },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateStopStateRequest {
    return new UpdateStopStateRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateStopStateRequest {
    return new UpdateStopStateRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateStopStateRequest {
    return new UpdateStopStateRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateStopStateRequest | PlainMessage<UpdateStopStateRequest> | undefined, b: UpdateStopStateRequest | PlainMessage<UpdateStopStateRequest> | undefined): boolean {
    return proto3.util.equals(UpdateStopStateRequest, a, b);
  }
}

/**
 * @generated from message state.v1.UpdateStopStateResponse
 */
export class UpdateStopStateResponse extends Message<UpdateStopStateResponse> {
  constructor(data?: PartialMessage<UpdateStopStateResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "state.v1.UpdateStopStateResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateStopStateResponse {
    return new UpdateStopStateResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateStopStateResponse {
    return new UpdateStopStateResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateStopStateResponse {
    return new UpdateStopStateResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateStopStateResponse | PlainMessage<UpdateStopStateResponse> | undefined, b: UpdateStopStateResponse | PlainMessage<UpdateStopStateResponse> | undefined): boolean {
    return proto3.util.equals(UpdateStopStateResponse, a, b);
  }
}

/**
 *
 * GetStopStates : 全てのストップレールの状態を取得するAPI
 *
 * @generated from message state.v1.GetStopStatesRequest
 */
export class GetStopStatesRequest extends Message<GetStopStatesRequest> {
  constructor(data?: PartialMessage<GetStopStatesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "state.v1.GetStopStatesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetStopStatesRequest {
    return new GetStopStatesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetStopStatesRequest {
    return new GetStopStatesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetStopStatesRequest {
    return new GetStopStatesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetStopStatesRequest | PlainMessage<GetStopStatesRequest> | undefined, b: GetStopStatesRequest | PlainMessage<GetStopStatesRequest> | undefined): boolean {
    return proto3.util.equals(GetStopStatesRequest, a, b);
  }
}

/**
 * @generated from message state.v1.GetStopStatesResponse
 */
export class GetStopStatesResponse extends Message<GetStopStatesResponse> {
  /**
   * @generated from field: repeated state.v1.StopAndState states = 1;
   */
  states: StopAndState[] = [];

  constructor(data?: PartialMessage<GetStopStatesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "state.v1.GetStopStatesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "states", kind: "message", T: StopAndState, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetStopStatesResponse {
    return new GetStopStatesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetStopStatesResponse {
    return new GetStopStatesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetStopStatesResponse {
    return new GetStopStatesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetStopStatesResponse | PlainMessage<GetStopStatesResponse> | undefined, b: GetStopStatesResponse | PlainMessage<GetStopStatesResponse> | undefined): boolean {
    return proto3.util.equals(GetStopStatesResponse, a, b);
  }
}

