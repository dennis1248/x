// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: mimi/announce/announce.proto

package announce

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	jsonfeed "within.website/x/buf/external/jsonfeed"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_mimi_announce_announce_proto protoreflect.FileDescriptor

var file_mimi_announce_announce_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x69, 0x6d, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x78,
	0x2e, 0x6d, 0x69, 0x6d, 0x69, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x40, 0x0a,
	0x08, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x41, 0x6e, 0x6e,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x66, 0x65, 0x65, 0x64,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42,
	0x24, 0x5a, 0x22, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x2f, 0x78, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x69, 0x6d, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_mimi_announce_announce_proto_goTypes = []any{
	(*jsonfeed.Item)(nil), // 0: jsonfeed.Item
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
}
var file_mimi_announce_announce_proto_depIdxs = []int32{
	0, // 0: within.website.x.mimi.announce.Announce.Announce:input_type -> jsonfeed.Item
	1, // 1: within.website.x.mimi.announce.Announce.Announce:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mimi_announce_announce_proto_init() }
func file_mimi_announce_announce_proto_init() {
	if File_mimi_announce_announce_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mimi_announce_announce_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mimi_announce_announce_proto_goTypes,
		DependencyIndexes: file_mimi_announce_announce_proto_depIdxs,
	}.Build()
	File_mimi_announce_announce_proto = out.File
	file_mimi_announce_announce_proto_rawDesc = nil
	file_mimi_announce_announce_proto_goTypes = nil
	file_mimi_announce_announce_proto_depIdxs = nil
}
