// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: bstream.proto

package stream

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MsgReply_MsgType int32

const (
	MsgReply_CONNECT_SUCCESS MsgReply_MsgType = 0
	MsgReply_CONNECT_FAILED  MsgReply_MsgType = 1
	MsgReply_NORMAL_MESSAGE  MsgReply_MsgType = 2
)

// Enum value maps for MsgReply_MsgType.
var (
	MsgReply_MsgType_name = map[int32]string{
		0: "CONNECT_SUCCESS",
		1: "CONNECT_FAILED",
		2: "NORMAL_MESSAGE",
	}
	MsgReply_MsgType_value = map[string]int32{
		"CONNECT_SUCCESS": 0,
		"CONNECT_FAILED":  1,
		"NORMAL_MESSAGE":  2,
	}
)

func (x MsgReply_MsgType) Enum() *MsgReply_MsgType {
	p := new(MsgReply_MsgType)
	*p = x
	return p
}

func (x MsgReply_MsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgReply_MsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_bstream_proto_enumTypes[0].Descriptor()
}

func (MsgReply_MsgType) Type() protoreflect.EnumType {
	return &file_bstream_proto_enumTypes[0]
}

func (x MsgReply_MsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgReply_MsgType.Descriptor instead.
func (MsgReply_MsgType) EnumDescriptor() ([]byte, []int) {
	return file_bstream_proto_rawDescGZIP(), []int{1, 0}
}

type MsgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MsgRequest) Reset() {
	*x = MsgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bstream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgRequest) ProtoMessage() {}

func (x *MsgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bstream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgRequest.ProtoReflect.Descriptor instead.
func (*MsgRequest) Descriptor() ([]byte, []int) {
	return file_bstream_proto_rawDescGZIP(), []int{0}
}

func (x *MsgRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MsgReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message     string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	TS          *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=TS,proto3" json:"TS,omitempty"`
	MessageType MsgReply_MsgType       `protobuf:"varint,3,opt,name=message_type,json=messageType,proto3,enum=bstream.MsgReply_MsgType" json:"message_type,omitempty"`
}

func (x *MsgReply) Reset() {
	*x = MsgReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bstream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgReply) ProtoMessage() {}

func (x *MsgReply) ProtoReflect() protoreflect.Message {
	mi := &file_bstream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgReply.ProtoReflect.Descriptor instead.
func (*MsgReply) Descriptor() ([]byte, []int) {
	return file_bstream_proto_rawDescGZIP(), []int{1}
}

func (x *MsgReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MsgReply) GetTS() *timestamppb.Timestamp {
	if x != nil {
		return x.TS
	}
	return nil
}

func (x *MsgReply) GetMessageType() MsgReply_MsgType {
	if x != nil {
		return x.MessageType
	}
	return MsgReply_CONNECT_SUCCESS
}

var File_bstream_proto protoreflect.FileDescriptor

var file_bstream_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0a, 0x4d, 0x73, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0xd6, 0x01, 0x0a, 0x08, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x02, 0x54, 0x53, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x02, 0x54, 0x53, 0x12, 0x3c, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x62, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x4d, 0x73,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x46, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a,
	0x0f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c,
	0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x02, 0x32, 0x49, 0x0a, 0x0e, 0x42, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x07,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x12, 0x13, 0x2e, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x62,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2e, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bstream_proto_rawDescOnce sync.Once
	file_bstream_proto_rawDescData = file_bstream_proto_rawDesc
)

func file_bstream_proto_rawDescGZIP() []byte {
	file_bstream_proto_rawDescOnce.Do(func() {
		file_bstream_proto_rawDescData = protoimpl.X.CompressGZIP(file_bstream_proto_rawDescData)
	})
	return file_bstream_proto_rawDescData
}

var file_bstream_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bstream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bstream_proto_goTypes = []interface{}{
	(MsgReply_MsgType)(0),         // 0: bstream.MsgReply.MsgType
	(*MsgRequest)(nil),            // 1: bstream.MsgRequest
	(*MsgReply)(nil),              // 2: bstream.MsgReply
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_bstream_proto_depIdxs = []int32{
	3, // 0: bstream.MsgReply.TS:type_name -> google.protobuf.Timestamp
	0, // 1: bstream.MsgReply.message_type:type_name -> bstream.MsgReply.MsgType
	1, // 2: bstream.BStreamService.SendMsg:input_type -> bstream.MsgRequest
	2, // 3: bstream.BStreamService.SendMsg:output_type -> bstream.MsgReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_bstream_proto_init() }
func file_bstream_proto_init() {
	if File_bstream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bstream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgRequest); i {
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
		file_bstream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgReply); i {
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
			RawDescriptor: file_bstream_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bstream_proto_goTypes,
		DependencyIndexes: file_bstream_proto_depIdxs,
		EnumInfos:         file_bstream_proto_enumTypes,
		MessageInfos:      file_bstream_proto_msgTypes,
	}.Build()
	File_bstream_proto = out.File
	file_bstream_proto_rawDesc = nil
	file_bstream_proto_goTypes = nil
	file_bstream_proto_depIdxs = nil
}
