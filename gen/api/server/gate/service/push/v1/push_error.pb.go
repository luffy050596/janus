// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: gate/service/push/v1/push_error.proto

package servicev1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PushServiceErrorReason int32

const (
	PushServiceErrorReason_PUSH_SERVICE_ERROR_REASON_UNSPECIFIED PushServiceErrorReason = 0
	PushServiceErrorReason_PUSH_SERVICE_ERROR_REASON_SERVER      PushServiceErrorReason = 1
	PushServiceErrorReason_PUSH_SERVICE_ERROR_REASON_SERVER_ID   PushServiceErrorReason = 2
)

// Enum value maps for PushServiceErrorReason.
var (
	PushServiceErrorReason_name = map[int32]string{
		0: "PUSH_SERVICE_ERROR_REASON_UNSPECIFIED",
		1: "PUSH_SERVICE_ERROR_REASON_SERVER",
		2: "PUSH_SERVICE_ERROR_REASON_SERVER_ID",
	}
	PushServiceErrorReason_value = map[string]int32{
		"PUSH_SERVICE_ERROR_REASON_UNSPECIFIED": 0,
		"PUSH_SERVICE_ERROR_REASON_SERVER":      1,
		"PUSH_SERVICE_ERROR_REASON_SERVER_ID":   2,
	}
)

func (x PushServiceErrorReason) Enum() *PushServiceErrorReason {
	p := new(PushServiceErrorReason)
	*p = x
	return p
}

func (x PushServiceErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PushServiceErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_gate_service_push_v1_push_error_proto_enumTypes[0].Descriptor()
}

func (PushServiceErrorReason) Type() protoreflect.EnumType {
	return &file_gate_service_push_v1_push_error_proto_enumTypes[0]
}

func (x PushServiceErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PushServiceErrorReason.Descriptor instead.
func (PushServiceErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_gate_service_push_v1_push_error_proto_rawDescGZIP(), []int{0}
}

var File_gate_service_push_v1_push_error_proto protoreflect.FileDescriptor

var file_gate_service_push_v1_push_error_proto_rawDesc = string([]byte{
	0x0a, 0x25, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x75, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x75, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xaa, 0x01, 0x0a, 0x16, 0x50, 0x75,
	0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x25, 0x50, 0x55, 0x53, 0x48, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f,
	0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x1a,
	0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x2a, 0x0a, 0x20, 0x50, 0x55, 0x53, 0x48, 0x5f, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0xf4,
	0x03, 0x12, 0x2d, 0x0a, 0x23, 0x50, 0x55, 0x53, 0x48, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43,
	0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x53,
	0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03,
	0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x2b, 0x5a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_gate_service_push_v1_push_error_proto_rawDescOnce sync.Once
	file_gate_service_push_v1_push_error_proto_rawDescData []byte
)

func file_gate_service_push_v1_push_error_proto_rawDescGZIP() []byte {
	file_gate_service_push_v1_push_error_proto_rawDescOnce.Do(func() {
		file_gate_service_push_v1_push_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_gate_service_push_v1_push_error_proto_rawDesc), len(file_gate_service_push_v1_push_error_proto_rawDesc)))
	})
	return file_gate_service_push_v1_push_error_proto_rawDescData
}

var file_gate_service_push_v1_push_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gate_service_push_v1_push_error_proto_goTypes = []any{
	(PushServiceErrorReason)(0), // 0: server.gate.service.push.v1.PushServiceErrorReason
}
var file_gate_service_push_v1_push_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gate_service_push_v1_push_error_proto_init() }
func file_gate_service_push_v1_push_error_proto_init() {
	if File_gate_service_push_v1_push_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_gate_service_push_v1_push_error_proto_rawDesc), len(file_gate_service_push_v1_push_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gate_service_push_v1_push_error_proto_goTypes,
		DependencyIndexes: file_gate_service_push_v1_push_error_proto_depIdxs,
		EnumInfos:         file_gate_service_push_v1_push_error_proto_enumTypes,
	}.Build()
	File_gate_service_push_v1_push_error_proto = out.File
	file_gate_service_push_v1_push_error_proto_goTypes = nil
	file_gate_service_push_v1_push_error_proto_depIdxs = nil
}
