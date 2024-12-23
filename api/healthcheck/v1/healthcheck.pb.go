// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.27.0
// source: healthcheck/v1/healthcheck.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type LivenessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LivenessRequest) Reset() {
	*x = LivenessRequest{}
	mi := &file_healthcheck_v1_healthcheck_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LivenessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LivenessRequest) ProtoMessage() {}

func (x *LivenessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_healthcheck_v1_healthcheck_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LivenessRequest.ProtoReflect.Descriptor instead.
func (*LivenessRequest) Descriptor() ([]byte, []int) {
	return file_healthcheck_v1_healthcheck_proto_rawDescGZIP(), []int{0}
}

type LivenessReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LivenessReply) Reset() {
	*x = LivenessReply{}
	mi := &file_healthcheck_v1_healthcheck_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LivenessReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LivenessReply) ProtoMessage() {}

func (x *LivenessReply) ProtoReflect() protoreflect.Message {
	mi := &file_healthcheck_v1_healthcheck_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LivenessReply.ProtoReflect.Descriptor instead.
func (*LivenessReply) Descriptor() ([]byte, []int) {
	return file_healthcheck_v1_healthcheck_proto_rawDescGZIP(), []int{1}
}

type ReadinessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReadinessRequest) Reset() {
	*x = ReadinessRequest{}
	mi := &file_healthcheck_v1_healthcheck_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadinessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadinessRequest) ProtoMessage() {}

func (x *ReadinessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_healthcheck_v1_healthcheck_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadinessRequest.ProtoReflect.Descriptor instead.
func (*ReadinessRequest) Descriptor() ([]byte, []int) {
	return file_healthcheck_v1_healthcheck_proto_rawDescGZIP(), []int{2}
}

type ReadinessReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReadinessReply) Reset() {
	*x = ReadinessReply{}
	mi := &file_healthcheck_v1_healthcheck_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadinessReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadinessReply) ProtoMessage() {}

func (x *ReadinessReply) ProtoReflect() protoreflect.Message {
	mi := &file_healthcheck_v1_healthcheck_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadinessReply.ProtoReflect.Descriptor instead.
func (*ReadinessReply) Descriptor() ([]byte, []int) {
	return file_healthcheck_v1_healthcheck_proto_rawDescGZIP(), []int{3}
}

var File_healthcheck_v1_healthcheck_proto protoreflect.FileDescriptor

var file_healthcheck_v1_healthcheck_proto_rawDesc = []byte{
	0x0a, 0x20, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x76, 0x31,
	0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x52, 0x65, 0x61, 0x64,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xdf, 0x01, 0x0a, 0x11, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x62, 0x0a, 0x0d, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x62,
	0x65, 0x12, 0x1f, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x6c, 0x69, 0x76, 0x65,
	0x6e, 0x65, 0x73, 0x73, 0x12, 0x66, 0x0a, 0x0e, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x20, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c,
	0x12, 0x0a, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x4c, 0x0a, 0x1d,
	0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56,
	0x31, 0x50, 0x01, 0x5a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_healthcheck_v1_healthcheck_proto_rawDescOnce sync.Once
	file_healthcheck_v1_healthcheck_proto_rawDescData = file_healthcheck_v1_healthcheck_proto_rawDesc
)

func file_healthcheck_v1_healthcheck_proto_rawDescGZIP() []byte {
	file_healthcheck_v1_healthcheck_proto_rawDescOnce.Do(func() {
		file_healthcheck_v1_healthcheck_proto_rawDescData = protoimpl.X.CompressGZIP(file_healthcheck_v1_healthcheck_proto_rawDescData)
	})
	return file_healthcheck_v1_healthcheck_proto_rawDescData
}

var file_healthcheck_v1_healthcheck_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_healthcheck_v1_healthcheck_proto_goTypes = []any{
	(*LivenessRequest)(nil),  // 0: healthcheck.v1.LivenessRequest
	(*LivenessReply)(nil),    // 1: healthcheck.v1.LivenessReply
	(*ReadinessRequest)(nil), // 2: healthcheck.v1.ReadinessRequest
	(*ReadinessReply)(nil),   // 3: healthcheck.v1.ReadinessReply
}
var file_healthcheck_v1_healthcheck_proto_depIdxs = []int32{
	0, // 0: healthcheck.v1.HealthcheckServer.LivenessProbe:input_type -> healthcheck.v1.LivenessRequest
	2, // 1: healthcheck.v1.HealthcheckServer.ReadinessProbe:input_type -> healthcheck.v1.ReadinessRequest
	1, // 2: healthcheck.v1.HealthcheckServer.LivenessProbe:output_type -> healthcheck.v1.LivenessReply
	3, // 3: healthcheck.v1.HealthcheckServer.ReadinessProbe:output_type -> healthcheck.v1.ReadinessReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_healthcheck_v1_healthcheck_proto_init() }
func file_healthcheck_v1_healthcheck_proto_init() {
	if File_healthcheck_v1_healthcheck_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_healthcheck_v1_healthcheck_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_healthcheck_v1_healthcheck_proto_goTypes,
		DependencyIndexes: file_healthcheck_v1_healthcheck_proto_depIdxs,
		MessageInfos:      file_healthcheck_v1_healthcheck_proto_msgTypes,
	}.Build()
	File_healthcheck_v1_healthcheck_proto = out.File
	file_healthcheck_v1_healthcheck_proto_rawDesc = nil
	file_healthcheck_v1_healthcheck_proto_goTypes = nil
	file_healthcheck_v1_healthcheck_proto_depIdxs = nil
}
