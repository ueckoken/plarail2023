//
//Train Proto
//駅に停車している列車の情報を扱うためのプロトコル

// @generated by protoc-gen-es v1.4.2 with parameter "target=dts+js"
// @generated from file state/v1/train.proto (package state.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum state.v1.Priority
 */
export declare enum Priority {
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

/**
 * @generated from message state.v1.Train
 */
export declare class Train extends Message<Train> {
  /**
   * 列車ID(NFCのUUIDと一意に対応している)
   *
   * @generated from field: string train_id = 1;
   */
  trainId: string;

  /**
   * 駅 or 閉塞のID
   *
   * @generated from field: string position_id = 2;
   */
  positionId: string;

  /**
   * 列車の優先度
   *
   * @generated from field: state.v1.Priority priority = 3;
   */
  priority: Priority;

  /**
   * NFCのUUID
   *
   * @generated from field: string uuid = 4;
   */
  uuid: string;

  /**
   * 行き先
   *
   * @generated from field: string destination = 5;
   */
  destination: string;

  constructor(data?: PartialMessage<Train>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "state.v1.Train";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Train;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Train;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Train;

  static equals(a: Train | PlainMessage<Train> | undefined, b: Train | PlainMessage<Train> | undefined): boolean;
}

/**
 *
 * GetTrains : 列車の状態を取得するAPI
 *
 * @generated from message state.v1.GetTrainsRequest
 */
export declare class GetTrainsRequest extends Message<GetTrainsRequest> {
  constructor(data?: PartialMessage<GetTrainsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "state.v1.GetTrainsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTrainsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTrainsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTrainsRequest;

  static equals(a: GetTrainsRequest | PlainMessage<GetTrainsRequest> | undefined, b: GetTrainsRequest | PlainMessage<GetTrainsRequest> | undefined): boolean;
}

/**
 * @generated from message state.v1.GetTrainsResponse
 */
export declare class GetTrainsResponse extends Message<GetTrainsResponse> {
  /**
   * @generated from field: repeated state.v1.Train trains = 1;
   */
  trains: Train[];

  constructor(data?: PartialMessage<GetTrainsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "state.v1.GetTrainsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTrainsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTrainsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTrainsResponse;

  static equals(a: GetTrainsResponse | PlainMessage<GetTrainsResponse> | undefined, b: GetTrainsResponse | PlainMessage<GetTrainsResponse> | undefined): boolean;
}

/**
 *
 * Add Train : 列車を追加するAPI
 *
 * @generated from message state.v1.AddTrainRequest
 */
export declare class AddTrainRequest extends Message<AddTrainRequest> {
  /**
   * @generated from field: state.v1.Train train = 1;
   */
  train?: Train;

  constructor(data?: PartialMessage<AddTrainRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "state.v1.AddTrainRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddTrainRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddTrainRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddTrainRequest;

  static equals(a: AddTrainRequest | PlainMessage<AddTrainRequest> | undefined, b: AddTrainRequest | PlainMessage<AddTrainRequest> | undefined): boolean;
}

/**
 * @generated from message state.v1.AddTrainResponse
 */
export declare class AddTrainResponse extends Message<AddTrainResponse> {
  /**
   * @generated from field: state.v1.Train train = 1;
   */
  train?: Train;

  constructor(data?: PartialMessage<AddTrainResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "state.v1.AddTrainResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddTrainResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddTrainResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddTrainResponse;

  static equals(a: AddTrainResponse | PlainMessage<AddTrainResponse> | undefined, b: AddTrainResponse | PlainMessage<AddTrainResponse> | undefined): boolean;
}

/**
 * @generated from message state.v1.UpdateTrainRequest
 */
export declare class UpdateTrainRequest extends Message<UpdateTrainRequest> {
  /**
   * @generated from field: state.v1.Train train = 1;
   */
  train?: Train;

  constructor(data?: PartialMessage<UpdateTrainRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "state.v1.UpdateTrainRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateTrainRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateTrainRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateTrainRequest;

  static equals(a: UpdateTrainRequest | PlainMessage<UpdateTrainRequest> | undefined, b: UpdateTrainRequest | PlainMessage<UpdateTrainRequest> | undefined): boolean;
}

/**
 * @generated from message state.v1.UpdateTrainResponse
 */
export declare class UpdateTrainResponse extends Message<UpdateTrainResponse> {
  /**
   * @generated from field: state.v1.Train train = 1;
   */
  train?: Train;

  constructor(data?: PartialMessage<UpdateTrainResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "state.v1.UpdateTrainResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateTrainResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateTrainResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateTrainResponse;

  static equals(a: UpdateTrainResponse | PlainMessage<UpdateTrainResponse> | undefined, b: UpdateTrainResponse | PlainMessage<UpdateTrainResponse> | undefined): boolean;
}

