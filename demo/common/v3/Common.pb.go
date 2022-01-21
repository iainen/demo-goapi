// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.0
// source: common/Common.proto

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

type ErrCode int32

const (
	ErrCode_ErrCode_None         ErrCode = 0
	ErrCode_ErrCode_SUCCESS      ErrCode = 1
	ErrCode_ErrCode_UnknownError ErrCode = 2
	ErrCode_ErrCode_ParamError   ErrCode = 3
	ErrCode_ErrCode_LogicError   ErrCode = 4
)

// Enum value maps for ErrCode.
var (
	ErrCode_name = map[int32]string{
		0: "ErrCode_None",
		1: "ErrCode_SUCCESS",
		2: "ErrCode_UnknownError",
		3: "ErrCode_ParamError",
		4: "ErrCode_LogicError",
	}
	ErrCode_value = map[string]int32{
		"ErrCode_None":         0,
		"ErrCode_SUCCESS":      1,
		"ErrCode_UnknownError": 2,
		"ErrCode_ParamError":   3,
		"ErrCode_LogicError":   4,
	}
)

func (x ErrCode) Enum() *ErrCode {
	p := new(ErrCode)
	*p = x
	return p
}

func (x ErrCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrCode) Descriptor() protoreflect.EnumDescriptor {
	return file_common_Common_proto_enumTypes[0].Descriptor()
}

func (ErrCode) Type() protoreflect.EnumType {
	return &file_common_Common_proto_enumTypes[0]
}

func (x ErrCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrCode.Descriptor instead.
func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return file_common_Common_proto_rawDescGZIP(), []int{0}
}

var File_common_Common_proto protoreflect.FileDescriptor

var file_common_Common_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x76, 0x33, 0x2a, 0x7a,
	0x0a, 0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x72, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x5f, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45,
	0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01,
	0x12, 0x18, 0x0a, 0x14, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x5f, 0x55, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x72,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x5f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x5f, 0x4c, 0x6f,
	0x67, 0x69, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x04, 0x42, 0x16, 0x5a, 0x14, 0x67, 0x6f,
	0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_Common_proto_rawDescOnce sync.Once
	file_common_Common_proto_rawDescData = file_common_Common_proto_rawDesc
)

func file_common_Common_proto_rawDescGZIP() []byte {
	file_common_Common_proto_rawDescOnce.Do(func() {
		file_common_Common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_Common_proto_rawDescData)
	})
	return file_common_Common_proto_rawDescData
}

var file_common_Common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_Common_proto_goTypes = []interface{}{
	(ErrCode)(0), // 0: demo.v3.ErrCode
}
var file_common_Common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_Common_proto_init() }
func file_common_Common_proto_init() {
	if File_common_Common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_Common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_Common_proto_goTypes,
		DependencyIndexes: file_common_Common_proto_depIdxs,
		EnumInfos:         file_common_Common_proto_enumTypes,
	}.Build()
	File_common_Common_proto = out.File
	file_common_Common_proto_rawDesc = nil
	file_common_Common_proto_goTypes = nil
	file_common_Common_proto_depIdxs = nil
}
