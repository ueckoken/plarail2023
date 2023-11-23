//
//Train Proto
//駅に停車している列車の情報を扱うためのプロトコル

// @generated by protoc-gen-es v1.4.2 with parameter "target=dts+js"
// @generated from file state/v1/train.proto (package state.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum state.v1.Priority
 */
export const Priority = proto3.makeEnum(
  "state.v1.Priority",
  [
    {no: 0, name: "PRIORITY_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "PRIORITY_LOW", localName: "LOW"},
    {no: 2, name: "PRIORITY_HIGH", localName: "HIGH"},
  ],
);

/**
 * @generated from message state.v1.Train
 */
export const Train = proto3.makeMessageType(
  "state.v1.Train",
  () => [
    { no: 1, name: "train_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "position_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "priority", kind: "enum", T: proto3.getEnumType(Priority) },
    { no: 4, name: "uuid", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "destination", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 *
 * GetTrains : 列車の状態を取得するAPI
 *
 * @generated from message state.v1.GetTrainsRequest
 */
export const GetTrainsRequest = proto3.makeMessageType(
  "state.v1.GetTrainsRequest",
  [],
);

/**
 * @generated from message state.v1.GetTrainsResponse
 */
export const GetTrainsResponse = proto3.makeMessageType(
  "state.v1.GetTrainsResponse",
  () => [
    { no: 1, name: "trains", kind: "message", T: Train, repeated: true },
  ],
);

/**
 *
 * Add Train : 列車を追加するAPI
 *
 * @generated from message state.v1.AddTrainRequest
 */
export const AddTrainRequest = proto3.makeMessageType(
  "state.v1.AddTrainRequest",
  () => [
    { no: 1, name: "train", kind: "message", T: Train },
  ],
);

/**
 * @generated from message state.v1.AddTrainResponse
 */
export const AddTrainResponse = proto3.makeMessageType(
  "state.v1.AddTrainResponse",
  () => [
    { no: 1, name: "train", kind: "message", T: Train },
  ],
);

/**
 * @generated from message state.v1.UpdateTrainRequest
 */
export const UpdateTrainRequest = proto3.makeMessageType(
  "state.v1.UpdateTrainRequest",
  () => [
    { no: 1, name: "train", kind: "message", T: Train },
  ],
);

/**
 * @generated from message state.v1.UpdateTrainResponse
 */
export const UpdateTrainResponse = proto3.makeMessageType(
  "state.v1.UpdateTrainResponse",
  () => [
    { no: 1, name: "train", kind: "message", T: Train },
  ],
);

