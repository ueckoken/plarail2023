//
//Train Proto
//駅に停車している列車の情報を扱うためのプロトコル

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: state/v1/train.proto

package statev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Priority int32

const (
	Priority_PRIORITY_UNSPECIFIED Priority = 0
	Priority_PRIORITY_LOW         Priority = 1
	Priority_PRIORITY_HIGH        Priority = 2
)

// Enum value maps for Priority.
var (
	Priority_name = map[int32]string{
		0: "PRIORITY_UNSPECIFIED",
		1: "PRIORITY_LOW",
		2: "PRIORITY_HIGH",
	}
	Priority_value = map[string]int32{
		"PRIORITY_UNSPECIFIED": 0,
		"PRIORITY_LOW":         1,
		"PRIORITY_HIGH":        2,
	}
)

func (x Priority) Enum() *Priority {
	p := new(Priority)
	*p = x
	return p
}

func (x Priority) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Priority) Descriptor() protoreflect.EnumDescriptor {
	return file_state_v1_train_proto_enumTypes[0].Descriptor()
}

func (Priority) Type() protoreflect.EnumType {
	return &file_state_v1_train_proto_enumTypes[0]
}

func (x Priority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Priority.Descriptor instead.
func (Priority) EnumDescriptor() ([]byte, []int) {
	return file_state_v1_train_proto_rawDescGZIP(), []int{0}
}

type Train struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrainId   string `protobuf:"bytes,1,opt,name=train_id,json=trainId,proto3" json:"train_id,omitempty"`       // 列車ID(NFCのUUIDと一意に対応している)
	StationId string `protobuf:"bytes,2,opt,name=station_id,json=stationId,proto3" json:"station_id,omitempty"` // 駅のID
}

func (x *Train) Reset() {
	*x = Train{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_v1_train_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Train) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Train) ProtoMessage() {}

func (x *Train) ProtoReflect() protoreflect.Message {
	mi := &file_state_v1_train_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Train.ProtoReflect.Descriptor instead.
func (*Train) Descriptor() ([]byte, []int) {
	return file_state_v1_train_proto_rawDescGZIP(), []int{0}
}

func (x *Train) GetTrainId() string {
	if x != nil {
		return x.TrainId
	}
	return ""
}

func (x *Train) GetStationId() string {
	if x != nil {
		return x.StationId
	}
	return ""
}

// GetTrains : 列車の状態を取得するAPI
type GetTrainsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTrainsRequest) Reset() {
	*x = GetTrainsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_v1_train_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrainsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrainsRequest) ProtoMessage() {}

func (x *GetTrainsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_state_v1_train_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrainsRequest.ProtoReflect.Descriptor instead.
func (*GetTrainsRequest) Descriptor() ([]byte, []int) {
	return file_state_v1_train_proto_rawDescGZIP(), []int{1}
}

type GetTrainsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trains []*Train `protobuf:"bytes,1,rep,name=trains,proto3" json:"trains,omitempty"`
}

func (x *GetTrainsResponse) Reset() {
	*x = GetTrainsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_v1_train_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrainsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrainsResponse) ProtoMessage() {}

func (x *GetTrainsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_state_v1_train_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrainsResponse.ProtoReflect.Descriptor instead.
func (*GetTrainsResponse) Descriptor() ([]byte, []int) {
	return file_state_v1_train_proto_rawDescGZIP(), []int{2}
}

func (x *GetTrainsResponse) GetTrains() []*Train {
	if x != nil {
		return x.Trains
	}
	return nil
}

// UpdateTrainUUID : NFCのUUID紐付けを更新するAPI
type UpdateTrainUUIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrainId string `protobuf:"bytes,1,opt,name=train_id,json=trainId,proto3" json:"train_id,omitempty"` // 列車ID
	Uuid    string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`                      // NFCのUUID
}

func (x *UpdateTrainUUIDRequest) Reset() {
	*x = UpdateTrainUUIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_v1_train_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTrainUUIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTrainUUIDRequest) ProtoMessage() {}

func (x *UpdateTrainUUIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_state_v1_train_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTrainUUIDRequest.ProtoReflect.Descriptor instead.
func (*UpdateTrainUUIDRequest) Descriptor() ([]byte, []int) {
	return file_state_v1_train_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateTrainUUIDRequest) GetTrainId() string {
	if x != nil {
		return x.TrainId
	}
	return ""
}

func (x *UpdateTrainUUIDRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type UpdateTrainUUIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateTrainUUIDResponse) Reset() {
	*x = UpdateTrainUUIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_v1_train_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTrainUUIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTrainUUIDResponse) ProtoMessage() {}

func (x *UpdateTrainUUIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_state_v1_train_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTrainUUIDResponse.ProtoReflect.Descriptor instead.
func (*UpdateTrainUUIDResponse) Descriptor() ([]byte, []int) {
	return file_state_v1_train_proto_rawDescGZIP(), []int{4}
}

var File_state_v1_train_proto protoreflect.FileDescriptor

var file_state_v1_train_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x41, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x61, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a,
	0x06, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x06,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x22, 0x47, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x69, 0x6e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22,
	0x19, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x55, 0x55,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x49, 0x0a, 0x08, 0x50, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49,
	0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4c, 0x4f, 0x57,
	0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x48,
	0x49, 0x47, 0x48, 0x10, 0x02, 0x42, 0x8d, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x75, 0x65, 0x63, 0x6b, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x70, 0x6c, 0x61, 0x72, 0x61, 0x69,
	0x6c, 0x32, 0x30, 0x32, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x53, 0x74, 0x61, 0x74, 0x65, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x14, 0x53, 0x74, 0x61, 0x74, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_state_v1_train_proto_rawDescOnce sync.Once
	file_state_v1_train_proto_rawDescData = file_state_v1_train_proto_rawDesc
)

func file_state_v1_train_proto_rawDescGZIP() []byte {
	file_state_v1_train_proto_rawDescOnce.Do(func() {
		file_state_v1_train_proto_rawDescData = protoimpl.X.CompressGZIP(file_state_v1_train_proto_rawDescData)
	})
	return file_state_v1_train_proto_rawDescData
}

var file_state_v1_train_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_state_v1_train_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_state_v1_train_proto_goTypes = []interface{}{
	(Priority)(0),                   // 0: state.v1.Priority
	(*Train)(nil),                   // 1: state.v1.Train
	(*GetTrainsRequest)(nil),        // 2: state.v1.GetTrainsRequest
	(*GetTrainsResponse)(nil),       // 3: state.v1.GetTrainsResponse
	(*UpdateTrainUUIDRequest)(nil),  // 4: state.v1.UpdateTrainUUIDRequest
	(*UpdateTrainUUIDResponse)(nil), // 5: state.v1.UpdateTrainUUIDResponse
}
var file_state_v1_train_proto_depIdxs = []int32{
	1, // 0: state.v1.GetTrainsResponse.trains:type_name -> state.v1.Train
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_state_v1_train_proto_init() }
func file_state_v1_train_proto_init() {
	if File_state_v1_train_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_state_v1_train_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Train); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_state_v1_train_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrainsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_state_v1_train_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrainsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_state_v1_train_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTrainUUIDRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_state_v1_train_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTrainUUIDResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_state_v1_train_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_state_v1_train_proto_goTypes,
		DependencyIndexes: file_state_v1_train_proto_depIdxs,
		EnumInfos:         file_state_v1_train_proto_enumTypes,
		MessageInfos:      file_state_v1_train_proto_msgTypes,
	}.Build()
	File_state_v1_train_proto = out.File
	file_state_v1_train_proto_rawDesc = nil
	file_state_v1_train_proto_goTypes = nil
	file_state_v1_train_proto_depIdxs = nil
}
