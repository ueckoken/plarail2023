//
//Block Proto
//閉塞の状態を扱うプロトコル

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: state/v1/block.proto

package statev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BlockStateEnum int32

const (
	BlockStateEnum_BLOCK_STATE_UNKNOWN BlockStateEnum = 0
	BlockStateEnum_BLOCK_STATE_OPEN    BlockStateEnum = 1 // 閉塞が開の状態(列車がいない)
	BlockStateEnum_BLOCK_STATE_CLOSE   BlockStateEnum = 2 // 閉塞が閉の状態(列車がいない)
)

// Enum value maps for BlockStateEnum.
var (
	BlockStateEnum_name = map[int32]string{
		0: "BLOCK_STATE_UNKNOWN",
		1: "BLOCK_STATE_OPEN",
		2: "BLOCK_STATE_CLOSE",
	}
	BlockStateEnum_value = map[string]int32{
		"BLOCK_STATE_UNKNOWN": 0,
		"BLOCK_STATE_OPEN":    1,
		"BLOCK_STATE_CLOSE":   2,
	}
)

func (x BlockStateEnum) Enum() *BlockStateEnum {
	p := new(BlockStateEnum)
	*p = x
	return p
}

func (x BlockStateEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockStateEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_state_v1_block_proto_enumTypes[0].Descriptor()
}

func (BlockStateEnum) Type() protoreflect.EnumType {
	return &file_state_v1_block_proto_enumTypes[0]
}

func (x BlockStateEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlockStateEnum.Descriptor instead.
func (BlockStateEnum) EnumDescriptor() ([]byte, []int) {
	return file_state_v1_block_proto_rawDescGZIP(), []int{0}
}

// 閉塞の状態
type BlockState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId string         `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`            // 閉塞のID
	State   BlockStateEnum `protobuf:"varint,2,opt,name=state,proto3,enum=state.v1.BlockStateEnum" json:"state,omitempty"` // 閉塞の状態
}

func (x *BlockState) Reset() {
	*x = BlockState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_v1_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockState) ProtoMessage() {}

func (x *BlockState) ProtoReflect() protoreflect.Message {
	mi := &file_state_v1_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockState.ProtoReflect.Descriptor instead.
func (*BlockState) Descriptor() ([]byte, []int) {
	return file_state_v1_block_proto_rawDescGZIP(), []int{0}
}

func (x *BlockState) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

func (x *BlockState) GetState() BlockStateEnum {
	if x != nil {
		return x.State
	}
	return BlockStateEnum_BLOCK_STATE_UNKNOWN
}

// GetBlockStates : 閉塞の状態を取得するAPI
type GetBlockStatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBlockStatesRequest) Reset() {
	*x = GetBlockStatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_v1_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockStatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockStatesRequest) ProtoMessage() {}

func (x *GetBlockStatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_state_v1_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockStatesRequest.ProtoReflect.Descriptor instead.
func (*GetBlockStatesRequest) Descriptor() ([]byte, []int) {
	return file_state_v1_block_proto_rawDescGZIP(), []int{1}
}

type GetBlockStatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	States []*BlockState `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty"`
}

func (x *GetBlockStatesResponse) Reset() {
	*x = GetBlockStatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_v1_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockStatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockStatesResponse) ProtoMessage() {}

func (x *GetBlockStatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_state_v1_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockStatesResponse.ProtoReflect.Descriptor instead.
func (*GetBlockStatesResponse) Descriptor() ([]byte, []int) {
	return file_state_v1_block_proto_rawDescGZIP(), []int{2}
}

func (x *GetBlockStatesResponse) GetStates() []*BlockState {
	if x != nil {
		return x.States
	}
	return nil
}

// UpdateBLockState: 閉塞の状態を更新するAPI
type UpdateBlockStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State *BlockState `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *UpdateBlockStateRequest) Reset() {
	*x = UpdateBlockStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_v1_block_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBlockStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlockStateRequest) ProtoMessage() {}

func (x *UpdateBlockStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_state_v1_block_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlockStateRequest.ProtoReflect.Descriptor instead.
func (*UpdateBlockStateRequest) Descriptor() ([]byte, []int) {
	return file_state_v1_block_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateBlockStateRequest) GetState() *BlockState {
	if x != nil {
		return x.State
	}
	return nil
}

type UpdateBlockStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateBlockStateResponse) Reset() {
	*x = UpdateBlockStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_v1_block_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBlockStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlockStateResponse) ProtoMessage() {}

func (x *UpdateBlockStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_state_v1_block_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlockStateResponse.ProtoReflect.Descriptor instead.
func (*UpdateBlockStateResponse) Descriptor() ([]byte, []int) {
	return file_state_v1_block_proto_rawDescGZIP(), []int{4}
}

var File_state_v1_block_proto protoreflect.FileDescriptor

var file_state_v1_block_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x22, 0x57, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e,
	0x75, 0x6d, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x46, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x22, 0x45, 0x0a, 0x17, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x1a, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x56, 0x0a,
	0x0e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12,
	0x17, 0x0a, 0x13, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x42, 0x4c, 0x4f, 0x43,
	0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x15,
	0x0a, 0x11, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4c,
	0x4f, 0x53, 0x45, 0x10, 0x02, 0x42, 0x9a, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x75, 0x65, 0x63, 0x6b, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x70, 0x6c, 0x61, 0x72, 0x61, 0x69,
	0x6c, 0x32, 0x30, 0x32, 0x33, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x73, 0x70,
	0x65, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x53, 0x74, 0x61, 0x74, 0x65, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x14, 0x53, 0x74, 0x61, 0x74, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_state_v1_block_proto_rawDescOnce sync.Once
	file_state_v1_block_proto_rawDescData = file_state_v1_block_proto_rawDesc
)

func file_state_v1_block_proto_rawDescGZIP() []byte {
	file_state_v1_block_proto_rawDescOnce.Do(func() {
		file_state_v1_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_state_v1_block_proto_rawDescData)
	})
	return file_state_v1_block_proto_rawDescData
}

var file_state_v1_block_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_state_v1_block_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_state_v1_block_proto_goTypes = []interface{}{
	(BlockStateEnum)(0),              // 0: state.v1.BlockStateEnum
	(*BlockState)(nil),               // 1: state.v1.BlockState
	(*GetBlockStatesRequest)(nil),    // 2: state.v1.GetBlockStatesRequest
	(*GetBlockStatesResponse)(nil),   // 3: state.v1.GetBlockStatesResponse
	(*UpdateBlockStateRequest)(nil),  // 4: state.v1.UpdateBlockStateRequest
	(*UpdateBlockStateResponse)(nil), // 5: state.v1.UpdateBlockStateResponse
}
var file_state_v1_block_proto_depIdxs = []int32{
	0, // 0: state.v1.BlockState.state:type_name -> state.v1.BlockStateEnum
	1, // 1: state.v1.GetBlockStatesResponse.states:type_name -> state.v1.BlockState
	1, // 2: state.v1.UpdateBlockStateRequest.state:type_name -> state.v1.BlockState
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_state_v1_block_proto_init() }
func file_state_v1_block_proto_init() {
	if File_state_v1_block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_state_v1_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockState); i {
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
		file_state_v1_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockStatesRequest); i {
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
		file_state_v1_block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockStatesResponse); i {
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
		file_state_v1_block_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBlockStateRequest); i {
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
		file_state_v1_block_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBlockStateResponse); i {
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
			RawDescriptor: file_state_v1_block_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_state_v1_block_proto_goTypes,
		DependencyIndexes: file_state_v1_block_proto_depIdxs,
		EnumInfos:         file_state_v1_block_proto_enumTypes,
		MessageInfos:      file_state_v1_block_proto_msgTypes,
	}.Build()
	File_state_v1_block_proto = out.File
	file_state_v1_block_proto_rawDesc = nil
	file_state_v1_block_proto_goTypes = nil
	file_state_v1_block_proto_depIdxs = nil
}
