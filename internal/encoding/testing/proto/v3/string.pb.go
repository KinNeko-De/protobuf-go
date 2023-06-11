// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: v3/string.proto

package v3

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

type SingleString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fut string `protobuf:"bytes,1,opt,name=fut,proto3" json:"fut,omitempty"`
}

func (x *SingleString) Reset() {
	*x = SingleString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v3_string_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleString) ProtoMessage() {}

func (x *SingleString) ProtoReflect() protoreflect.Message {
	mi := &file_v3_string_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleString.ProtoReflect.Descriptor instead.
func (*SingleString) Descriptor() ([]byte, []int) {
	return file_v3_string_proto_rawDescGZIP(), []int{0}
}

func (x *SingleString) GetFut() string {
	if x != nil {
		return x.Fut
	}
	return ""
}

type MultipleStrings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Street string `protobuf:"bytes,2,opt,name=street,proto3" json:"street,omitempty"`
}

func (x *MultipleStrings) Reset() {
	*x = MultipleStrings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v3_string_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultipleStrings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultipleStrings) ProtoMessage() {}

func (x *MultipleStrings) ProtoReflect() protoreflect.Message {
	mi := &file_v3_string_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultipleStrings.ProtoReflect.Descriptor instead.
func (*MultipleStrings) Descriptor() ([]byte, []int) {
	return file_v3_string_proto_rawDescGZIP(), []int{1}
}

func (x *MultipleStrings) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MultipleStrings) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

var File_v3_string_proto protoreflect.FileDescriptor

var file_v3_string_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x67, 0x6f, 0x2e, 0x65,
	0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x33, 0x22, 0x20, 0x0a, 0x0c, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x75, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x75, 0x74, 0x22, 0x3d, 0x0a, 0x0f, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x42, 0x27, 0x5a, 0x25, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x67, 0x6f, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v3_string_proto_rawDescOnce sync.Once
	file_v3_string_proto_rawDescData = file_v3_string_proto_rawDesc
)

func file_v3_string_proto_rawDescGZIP() []byte {
	file_v3_string_proto_rawDescOnce.Do(func() {
		file_v3_string_proto_rawDescData = protoimpl.X.CompressGZIP(file_v3_string_proto_rawDescData)
	})
	return file_v3_string_proto_rawDescData
}

var file_v3_string_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v3_string_proto_goTypes = []interface{}{
	(*SingleString)(nil),    // 0: protobuf_go.encoding.testing.proto.v3.SingleString
	(*MultipleStrings)(nil), // 1: protobuf_go.encoding.testing.proto.v3.MultipleStrings
}
var file_v3_string_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v3_string_proto_init() }
func file_v3_string_proto_init() {
	if File_v3_string_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v3_string_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleString); i {
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
		file_v3_string_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultipleStrings); i {
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
			RawDescriptor: file_v3_string_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v3_string_proto_goTypes,
		DependencyIndexes: file_v3_string_proto_depIdxs,
		MessageInfos:      file_v3_string_proto_msgTypes,
	}.Build()
	File_v3_string_proto = out.File
	file_v3_string_proto_rawDesc = nil
	file_v3_string_proto_goTypes = nil
	file_v3_string_proto_depIdxs = nil
}
