//
//Train Proto
//駅に停車している列車の情報を扱うためのプロトコル

// @generated by protoc-gen-es v1.4.2 with parameter "target=ts"
// @generated from file state/v1/train.proto (package state.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum state.v1.Priority
 */
export enum Priority {
  /**
   * @generated from enum value: PRIORITY_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: PRIORITY_LOW = 1;
   */
  LOW = 1,

  /**
   * @generated from enum value: PRIORITY_HIGH = 2;
   */
  HIGH = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(Priority)
proto3.util.setEnumType(Priority, "state.v1.Priority", [
  { no: 0, name: "PRIORITY_UNSPECIFIED" },
  { no: 1, name: "PRIORITY_LOW" },
  { no: 2, name: "PRIORITY_HIGH" },
]);

/**
 * @generated from message state.v1.Train
 */
export class Train extends Message<Train> {
  /**
   * 列車ID(NFCのUUIDと一意に対応している)
   *
   * @generated from field: string train_id = 1;
   */
  trainId = "";

  /**
   * 駅のID
   *
   * @generated from field: string station_id = 2;
   */
  stationId = "";

  constructor(data?: PartialMessage<Train>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "state.v1.Train";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "train_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "station_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Train {
    return new Train().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Train {
    return new Train().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Train {
    return new Train().fromJsonString(jsonString, options);
  }

  static equals(a: Train | PlainMessage<Train> | undefined, b: Train | PlainMessage<Train> | undefined): boolean {
    return proto3.util.equals(Train, a, b);
  }
}

/**
 *
 * GetTrains : 列車の状態を取得するAPI
 *
 * @generated from message state.v1.GetTrainsRequest
 */
export class GetTrainsRequest extends Message<GetTrainsRequest> {
  constructor(data?: PartialMessage<GetTrainsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "state.v1.GetTrainsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTrainsRequest {
    return new GetTrainsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTrainsRequest {
    return new GetTrainsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTrainsRequest {
    return new GetTrainsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetTrainsRequest | PlainMessage<GetTrainsRequest> | undefined, b: GetTrainsRequest | PlainMessage<GetTrainsRequest> | undefined): boolean {
    return proto3.util.equals(GetTrainsRequest, a, b);
  }
}

/**
 * @generated from message state.v1.GetTrainsResponse
 */
export class GetTrainsResponse extends Message<GetTrainsResponse> {
  /**
   * @generated from field: repeated state.v1.Train trains = 1;
   */
  trains: Train[] = [];

  constructor(data?: PartialMessage<GetTrainsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "state.v1.GetTrainsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "trains", kind: "message", T: Train, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTrainsResponse {
    return new GetTrainsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTrainsResponse {
    return new GetTrainsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTrainsResponse {
    return new GetTrainsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetTrainsResponse | PlainMessage<GetTrainsResponse> | undefined, b: GetTrainsResponse | PlainMessage<GetTrainsResponse> | undefined): boolean {
    return proto3.util.equals(GetTrainsResponse, a, b);
  }
}

/**
 *
 * UpdateTrainUUID : NFCのUUID紐付けを更新するAPI
 *
 * @generated from message state.v1.UpdateTrainUUIDRequest
 */
export class UpdateTrainUUIDRequest extends Message<UpdateTrainUUIDRequest> {
  /**
   * 列車ID
   *
   * @generated from field: string train_id = 1;
   */
  trainId = "";

  /**
   * NFCのUUID
   *
   * @generated from field: string uuid = 2;
   */
  uuid = "";

  constructor(data?: PartialMessage<UpdateTrainUUIDRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "state.v1.UpdateTrainUUIDRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "train_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "uuid", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateTrainUUIDRequest {
    return new UpdateTrainUUIDRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateTrainUUIDRequest {
    return new UpdateTrainUUIDRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateTrainUUIDRequest {
    return new UpdateTrainUUIDRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateTrainUUIDRequest | PlainMessage<UpdateTrainUUIDRequest> | undefined, b: UpdateTrainUUIDRequest | PlainMessage<UpdateTrainUUIDRequest> | undefined): boolean {
    return proto3.util.equals(UpdateTrainUUIDRequest, a, b);
  }
}

/**
 * @generated from message state.v1.UpdateTrainUUIDResponse
 */
export class UpdateTrainUUIDResponse extends Message<UpdateTrainUUIDResponse> {
  constructor(data?: PartialMessage<UpdateTrainUUIDResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "state.v1.UpdateTrainUUIDResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateTrainUUIDResponse {
    return new UpdateTrainUUIDResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateTrainUUIDResponse {
    return new UpdateTrainUUIDResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateTrainUUIDResponse {
    return new UpdateTrainUUIDResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateTrainUUIDResponse | PlainMessage<UpdateTrainUUIDResponse> | undefined, b: UpdateTrainUUIDResponse | PlainMessage<UpdateTrainUUIDResponse> | undefined): boolean {
    return proto3.util.equals(UpdateTrainUUIDResponse, a, b);
  }
}
