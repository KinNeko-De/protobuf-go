// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: v3/google/protobuf/timestamp.proto

package protobuf

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

type SingleTimestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fut *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=fut,proto3" json:"fut,omitempty"`
}

func (x *SingleTimestamp) Reset() {
	*x = SingleTimestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v3_google_protobuf_timestamp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleTimestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleTimestamp) ProtoMessage() {}

func (x *SingleTimestamp) ProtoReflect() protoreflect.Message {
	mi := &file_v3_google_protobuf_timestamp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleTimestamp.ProtoReflect.Descriptor instead.
func (*SingleTimestamp) Descriptor() ([]byte, []int) {
	return file_v3_google_protobuf_timestamp_proto_rawDescGZIP(), []int{0}
}

func (x *SingleTimestamp) GetFut() *timestamppb.Timestamp {
	if x != nil {
		return x.Fut
	}
	return nil
}

type MultipleTimestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	One *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=one,proto3" json:"one,omitempty"`
	Two *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=two,proto3" json:"two,omitempty"`
}

func (x *MultipleTimestamp) Reset() {
	*x = MultipleTimestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v3_google_protobuf_timestamp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultipleTimestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultipleTimestamp) ProtoMessage() {}

func (x *MultipleTimestamp) ProtoReflect() protoreflect.Message {
	mi := &file_v3_google_protobuf_timestamp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultipleTimestamp.ProtoReflect.Descriptor instead.
func (*MultipleTimestamp) Descriptor() ([]byte, []int) {
	return file_v3_google_protobuf_timestamp_proto_rawDescGZIP(), []int{1}
}

func (x *MultipleTimestamp) GetOne() *timestamppb.Timestamp {
	if x != nil {
		return x.One
	}
	return nil
}

func (x *MultipleTimestamp) GetTwo() *timestamppb.Timestamp {
	if x != nil {
		return x.Two
	}
	return nil
}

var File_v3_google_protobuf_timestamp_proto protoreflect.FileDescriptor

var file_v3_google_protobuf_timestamp_proto_rawDesc = []byte{
	0x0a, 0x22, 0x76, 0x33, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x35, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x67,
	0x6f, 0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x0f,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x2c, 0x0a, 0x03, 0x66, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x66, 0x75, 0x74, 0x22, 0x6f, 0x0a,
	0x11, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x2c, 0x0a, 0x03, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x6f, 0x6e, 0x65,
	0x12, 0x2c, 0x0a, 0x03, 0x74, 0x77, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x74, 0x77, 0x6f, 0x42, 0x37,
	0x5a, 0x35, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x67, 0x6f, 0x2f, 0x65, 0x6e,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v3_google_protobuf_timestamp_proto_rawDescOnce sync.Once
	file_v3_google_protobuf_timestamp_proto_rawDescData = file_v3_google_protobuf_timestamp_proto_rawDesc
)

func file_v3_google_protobuf_timestamp_proto_rawDescGZIP() []byte {
	file_v3_google_protobuf_timestamp_proto_rawDescOnce.Do(func() {
		file_v3_google_protobuf_timestamp_proto_rawDescData = protoimpl.X.CompressGZIP(file_v3_google_protobuf_timestamp_proto_rawDescData)
	})
	return file_v3_google_protobuf_timestamp_proto_rawDescData
}

var file_v3_google_protobuf_timestamp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v3_google_protobuf_timestamp_proto_goTypes = []interface{}{
	(*SingleTimestamp)(nil),       // 0: protobuf_go.encoding.testing.proto.v3.google.protobuf.SingleTimestamp
	(*MultipleTimestamp)(nil),     // 1: protobuf_go.encoding.testing.proto.v3.google.protobuf.MultipleTimestamp
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_v3_google_protobuf_timestamp_proto_depIdxs = []int32{
	2, // 0: protobuf_go.encoding.testing.proto.v3.google.protobuf.SingleTimestamp.fut:type_name -> google.protobuf.Timestamp
	2, // 1: protobuf_go.encoding.testing.proto.v3.google.protobuf.MultipleTimestamp.one:type_name -> google.protobuf.Timestamp
	2, // 2: protobuf_go.encoding.testing.proto.v3.google.protobuf.MultipleTimestamp.two:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v3_google_protobuf_timestamp_proto_init() }
func file_v3_google_protobuf_timestamp_proto_init() {
	if File_v3_google_protobuf_timestamp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v3_google_protobuf_timestamp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleTimestamp); i {
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
		file_v3_google_protobuf_timestamp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultipleTimestamp); i {
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
			RawDescriptor: file_v3_google_protobuf_timestamp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v3_google_protobuf_timestamp_proto_goTypes,
		DependencyIndexes: file_v3_google_protobuf_timestamp_proto_depIdxs,
		MessageInfos:      file_v3_google_protobuf_timestamp_proto_msgTypes,
	}.Build()
	File_v3_google_protobuf_timestamp_proto = out.File
	file_v3_google_protobuf_timestamp_proto_rawDesc = nil
	file_v3_google_protobuf_timestamp_proto_goTypes = nil
	file_v3_google_protobuf_timestamp_proto_depIdxs = nil
}
