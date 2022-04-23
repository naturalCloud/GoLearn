// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: oneof.proto

package pb

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

type Boy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Age:
	//	*Boy_A
	//	*Boy_B
	//	*Boy_C
	Age isBoy_Age `protobuf_oneof:"Age"`
}

func (x *Boy) Reset() {
	*x = Boy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oneof_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Boy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Boy) ProtoMessage() {}

func (x *Boy) ProtoReflect() protoreflect.Message {
	mi := &file_oneof_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Boy.ProtoReflect.Descriptor instead.
func (*Boy) Descriptor() ([]byte, []int) {
	return file_oneof_proto_rawDescGZIP(), []int{0}
}

func (m *Boy) GetAge() isBoy_Age {
	if m != nil {
		return m.Age
	}
	return nil
}

func (x *Boy) GetA() int32 {
	if x, ok := x.GetAge().(*Boy_A); ok {
		return x.A
	}
	return 0
}

func (x *Boy) GetB() int32 {
	if x, ok := x.GetAge().(*Boy_B); ok {
		return x.B
	}
	return 0
}

func (x *Boy) GetC() int32 {
	if x, ok := x.GetAge().(*Boy_C); ok {
		return x.C
	}
	return 0
}

type isBoy_Age interface {
	isBoy_Age()
}

type Boy_A struct {
	A int32 `protobuf:"varint,1,opt,name=a,proto3,oneof"`
}

type Boy_B struct {
	B int32 `protobuf:"varint,2,opt,name=b,proto3,oneof"`
}

type Boy_C struct {
	C int32 `protobuf:"varint,3,opt,name=c,proto3,oneof"`
}

func (*Boy_A) isBoy_Age() {}

func (*Boy_B) isBoy_Age() {}

func (*Boy_C) isBoy_Age() {}

type Girl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Age int32 `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *Girl) Reset() {
	*x = Girl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oneof_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Girl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Girl) ProtoMessage() {}

func (x *Girl) ProtoReflect() protoreflect.Message {
	mi := &file_oneof_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Girl.ProtoReflect.Descriptor instead.
func (*Girl) Descriptor() ([]byte, []int) {
	return file_oneof_proto_rawDescGZIP(), []int{1}
}

func (x *Girl) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

var File_oneof_proto protoreflect.FileDescriptor

var file_oneof_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x3c, 0x0a, 0x03, 0x42, 0x6f, 0x79, 0x12, 0x0e, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x01, 0x61, 0x12, 0x0e, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x01, 0x62, 0x12, 0x0e, 0x0a, 0x01, 0x63, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x01, 0x63, 0x42, 0x05, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x22,
	0x18, 0x0a, 0x04, 0x47, 0x69, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x32, 0x29, 0x0a, 0x06, 0x43, 0x68, 0x6f,
	0x6f, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x65, 0x73, 0x74, 0x43, 0x68, 0x6f, 0x6f, 0x73,
	0x65, 0x12, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x6f, 0x79, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x69, 0x72, 0x6c, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oneof_proto_rawDescOnce sync.Once
	file_oneof_proto_rawDescData = file_oneof_proto_rawDesc
)

func file_oneof_proto_rawDescGZIP() []byte {
	file_oneof_proto_rawDescOnce.Do(func() {
		file_oneof_proto_rawDescData = protoimpl.X.CompressGZIP(file_oneof_proto_rawDescData)
	})
	return file_oneof_proto_rawDescData
}

var file_oneof_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_oneof_proto_goTypes = []interface{}{
	(*Boy)(nil),  // 0: pb.Boy
	(*Girl)(nil), // 1: pb.Girl
}
var file_oneof_proto_depIdxs = []int32{
	0, // 0: pb.Choose.bestChoose:input_type -> pb.Boy
	1, // 1: pb.Choose.bestChoose:output_type -> pb.Girl
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_oneof_proto_init() }
func file_oneof_proto_init() {
	if File_oneof_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oneof_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Boy); i {
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
		file_oneof_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Girl); i {
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
	file_oneof_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Boy_A)(nil),
		(*Boy_B)(nil),
		(*Boy_C)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_oneof_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_oneof_proto_goTypes,
		DependencyIndexes: file_oneof_proto_depIdxs,
		MessageInfos:      file_oneof_proto_msgTypes,
	}.Build()
	File_oneof_proto = out.File
	file_oneof_proto_rawDesc = nil
	file_oneof_proto_goTypes = nil
	file_oneof_proto_depIdxs = nil
}