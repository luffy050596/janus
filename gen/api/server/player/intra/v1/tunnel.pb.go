// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: player/intra/v1/tunnel.proto

package intrav1

import (
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

type Message struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Mod           int32                  `protobuf:"varint,1,opt,name=mod,proto3" json:"mod,omitempty"`                                    // Module ID, globally unique
	Seq           int32                  `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`                                    // Module message ID, unique within the module
	Obj           int64                  `protobuf:"varint,3,opt,name=obj,proto3" json:"obj,omitempty"`                                    // Module object ID, according to the business agreement to pass the corresponding object ID
	Data          []byte                 `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`                                   // Serialized bytes of the cs/sc protocol in the message
	DataVersion   uint64                 `protobuf:"varint,5,opt,name=data_version,json=dataVersion,proto3" json:"data_version,omitempty"` // Data version number
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_player_intra_v1_tunnel_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_player_intra_v1_tunnel_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_player_intra_v1_tunnel_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetMod() int32 {
	if x != nil {
		return x.Mod
	}
	return 0
}

func (x *Message) GetSeq() int32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *Message) GetObj() int64 {
	if x != nil {
		return x.Obj
	}
	return 0
}

func (x *Message) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Message) GetDataVersion() uint64 {
	if x != nil {
		return x.DataVersion
	}
	return 0
}

var File_player_intra_v1_tunnel_proto protoreflect.FileDescriptor

var file_player_intra_v1_tunnel_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x69, 0x6e,
	0x74, 0x72, 0x61, 0x2e, 0x76, 0x31, 0x22, 0x76, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x6d, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x62, 0x6a, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x6f, 0x62, 0x6a, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x61,
	0x0a, 0x0d, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x50, 0x0a, 0x06, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x74, 0x72, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x74, 0x72, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x24, 0x5a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x72, 0x61, 0x2f, 0x76, 0x31, 0x3b,
	0x69, 0x6e, 0x74, 0x72, 0x61, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_player_intra_v1_tunnel_proto_rawDescOnce sync.Once
	file_player_intra_v1_tunnel_proto_rawDescData []byte
)

func file_player_intra_v1_tunnel_proto_rawDescGZIP() []byte {
	file_player_intra_v1_tunnel_proto_rawDescOnce.Do(func() {
		file_player_intra_v1_tunnel_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_player_intra_v1_tunnel_proto_rawDesc), len(file_player_intra_v1_tunnel_proto_rawDesc)))
	})
	return file_player_intra_v1_tunnel_proto_rawDescData
}

var file_player_intra_v1_tunnel_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_player_intra_v1_tunnel_proto_goTypes = []any{
	(*Message)(nil), // 0: server.player.intra.v1.Message
}
var file_player_intra_v1_tunnel_proto_depIdxs = []int32{
	0, // 0: server.player.intra.v1.TunnelService.Tunnel:input_type -> server.player.intra.v1.Message
	0, // 1: server.player.intra.v1.TunnelService.Tunnel:output_type -> server.player.intra.v1.Message
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_player_intra_v1_tunnel_proto_init() }
func file_player_intra_v1_tunnel_proto_init() {
	if File_player_intra_v1_tunnel_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_player_intra_v1_tunnel_proto_rawDesc), len(file_player_intra_v1_tunnel_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_player_intra_v1_tunnel_proto_goTypes,
		DependencyIndexes: file_player_intra_v1_tunnel_proto_depIdxs,
		MessageInfos:      file_player_intra_v1_tunnel_proto_msgTypes,
	}.Build()
	File_player_intra_v1_tunnel_proto = out.File
	file_player_intra_v1_tunnel_proto_goTypes = nil
	file_player_intra_v1_tunnel_proto_depIdxs = nil
}
