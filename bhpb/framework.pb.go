// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: framework.proto

package bhpb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FrameworkType int32

const (
	FrameworkType_FRAMEWORKTYPE_UNKNOWN         FrameworkType = 0
	FrameworkType_FRAMEWORKTYPE_UNIT_TEST       FrameworkType = 1
	FrameworkType_FRAMEWORKTYPE_MICRO_BENCHMARK FrameworkType = 2
)

// Enum value maps for FrameworkType.
var (
	FrameworkType_name = map[int32]string{
		0: "FRAMEWORKTYPE_UNKNOWN",
		1: "FRAMEWORKTYPE_UNIT_TEST",
		2: "FRAMEWORKTYPE_MICRO_BENCHMARK",
	}
	FrameworkType_value = map[string]int32{
		"FRAMEWORKTYPE_UNKNOWN":         0,
		"FRAMEWORKTYPE_UNIT_TEST":       1,
		"FRAMEWORKTYPE_MICRO_BENCHMARK": 2,
	}
)

func (x FrameworkType) Enum() *FrameworkType {
	p := new(FrameworkType)
	*p = x
	return p
}

func (x FrameworkType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FrameworkType) Descriptor() protoreflect.EnumDescriptor {
	return file_framework_proto_enumTypes[0].Descriptor()
}

func (FrameworkType) Type() protoreflect.EnumType {
	return &file_framework_proto_enumTypes[0]
}

func (x FrameworkType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FrameworkType.Descriptor instead.
func (FrameworkType) EnumDescriptor() ([]byte, []int) {
	return file_framework_proto_rawDescGZIP(), []int{0}
}

var File_framework_proto protoreflect.FileDescriptor

var file_framework_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x62, 0x68, 0x70, 0x62, 0x2a, 0x6a, 0x0a, 0x0d, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x46, 0x52, 0x41, 0x4d,
	0x45, 0x57, 0x4f, 0x52, 0x4b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x57, 0x4f, 0x52, 0x4b,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x10, 0x01,
	0x12, 0x21, 0x0a, 0x1d, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x57, 0x4f, 0x52, 0x4b, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4d, 0x49, 0x43, 0x52, 0x4f, 0x5f, 0x42, 0x45, 0x4e, 0x43, 0x48, 0x4d, 0x41, 0x52,
	0x4b, 0x10, 0x02, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x68, 0x75, 0x62, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68,
	0x68, 0x75, 0x62, 0x2f, 0x62, 0x68, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_framework_proto_rawDescOnce sync.Once
	file_framework_proto_rawDescData = file_framework_proto_rawDesc
)

func file_framework_proto_rawDescGZIP() []byte {
	file_framework_proto_rawDescOnce.Do(func() {
		file_framework_proto_rawDescData = protoimpl.X.CompressGZIP(file_framework_proto_rawDescData)
	})
	return file_framework_proto_rawDescData
}

var file_framework_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_framework_proto_goTypes = []interface{}{
	(FrameworkType)(0), // 0: bhpb.FrameworkType
}
var file_framework_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_framework_proto_init() }
func file_framework_proto_init() {
	if File_framework_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_framework_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_framework_proto_goTypes,
		DependencyIndexes: file_framework_proto_depIdxs,
		EnumInfos:         file_framework_proto_enumTypes,
	}.Build()
	File_framework_proto = out.File
	file_framework_proto_rawDesc = nil
	file_framework_proto_goTypes = nil
	file_framework_proto_depIdxs = nil
}