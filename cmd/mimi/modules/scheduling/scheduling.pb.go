// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.23.4
// source: scheduling.proto

package scheduling

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

type ConversationMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role  string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *ConversationMember) Reset() {
	*x = ConversationMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduling_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversationMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversationMember) ProtoMessage() {}

func (x *ConversationMember) ProtoReflect() protoreflect.Message {
	mi := &file_scheduling_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversationMember.ProtoReflect.Descriptor instead.
func (*ConversationMember) Descriptor() ([]byte, []int) {
	return file_scheduling_proto_rawDescGZIP(), []int{0}
}

func (x *ConversationMember) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *ConversationMember) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConversationMember) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ParseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Month               string                `protobuf:"bytes,1,opt,name=month,proto3" json:"month,omitempty"`
	ConversationMembers []*ConversationMember `protobuf:"bytes,2,rep,name=conversation_members,json=conversationMembers,proto3" json:"conversation_members,omitempty"`
	Message             string                `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Date                string                `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *ParseReq) Reset() {
	*x = ParseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduling_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseReq) ProtoMessage() {}

func (x *ParseReq) ProtoReflect() protoreflect.Message {
	mi := &file_scheduling_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseReq.ProtoReflect.Descriptor instead.
func (*ParseReq) Descriptor() ([]byte, []int) {
	return file_scheduling_proto_rawDescGZIP(), []int{1}
}

func (x *ParseReq) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *ParseReq) GetConversationMembers() []*ConversationMember {
	if x != nil {
		return x.ConversationMembers
	}
	return nil
}

func (x *ParseReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ParseReq) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type ParseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime string                `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Duration  string                `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	Summary   string                `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`
	Attendees []*ConversationMember `protobuf:"bytes,4,rep,name=attendees,proto3" json:"attendees,omitempty"`
	Location  string                `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *ParseResp) Reset() {
	*x = ParseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduling_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseResp) ProtoMessage() {}

func (x *ParseResp) ProtoReflect() protoreflect.Message {
	mi := &file_scheduling_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseResp.ProtoReflect.Descriptor instead.
func (*ParseResp) Descriptor() ([]byte, []int) {
	return file_scheduling_proto_rawDescGZIP(), []int{2}
}

func (x *ParseResp) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *ParseResp) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *ParseResp) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *ParseResp) GetAttendees() []*ConversationMember {
	if x != nil {
		return x.Attendees
	}
	return nil
}

func (x *ParseResp) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

var File_scheduling_proto protoreflect.FileDescriptor

var file_scheduling_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x20, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x2e, 0x78, 0x2e, 0x6d, 0x69, 0x6d, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x69, 0x6e, 0x67, 0x22, 0x52, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xb7, 0x01, 0x0a, 0x08, 0x50, 0x61, 0x72,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x67, 0x0a, 0x14, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x77, 0x69, 0x74, 0x68,
	0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x78, 0x2e, 0x6d, 0x69, 0x6d,
	0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x13, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x22, 0xd0, 0x01, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x52, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x69,
	0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x78, 0x2e, 0x6d, 0x69, 0x6d, 0x69,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x09,
	0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x73, 0x0a, 0x0a, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x69, 0x6e, 0x67, 0x12, 0x65, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x73, 0x65, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x2a, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x2e, 0x78, 0x2e, 0x6d, 0x69, 0x6d, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x2b, 0x2e,
	0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x78,
	0x2e, 0x6d, 0x69, 0x6d, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x2e, 0x5a, 0x2c, 0x77, 0x69,
	0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2f, 0x78, 0x2f, 0x63,
	0x6d, 0x64, 0x2f, 0x6d, 0x69, 0x6d, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_scheduling_proto_rawDescOnce sync.Once
	file_scheduling_proto_rawDescData = file_scheduling_proto_rawDesc
)

func file_scheduling_proto_rawDescGZIP() []byte {
	file_scheduling_proto_rawDescOnce.Do(func() {
		file_scheduling_proto_rawDescData = protoimpl.X.CompressGZIP(file_scheduling_proto_rawDescData)
	})
	return file_scheduling_proto_rawDescData
}

var file_scheduling_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_scheduling_proto_goTypes = []interface{}{
	(*ConversationMember)(nil), // 0: within.website.x.mimi.scheduling.ConversationMember
	(*ParseReq)(nil),           // 1: within.website.x.mimi.scheduling.ParseReq
	(*ParseResp)(nil),          // 2: within.website.x.mimi.scheduling.ParseResp
}
var file_scheduling_proto_depIdxs = []int32{
	0, // 0: within.website.x.mimi.scheduling.ParseReq.conversation_members:type_name -> within.website.x.mimi.scheduling.ConversationMember
	0, // 1: within.website.x.mimi.scheduling.ParseResp.attendees:type_name -> within.website.x.mimi.scheduling.ConversationMember
	1, // 2: within.website.x.mimi.scheduling.Scheduling.ParseEmail:input_type -> within.website.x.mimi.scheduling.ParseReq
	2, // 3: within.website.x.mimi.scheduling.Scheduling.ParseEmail:output_type -> within.website.x.mimi.scheduling.ParseResp
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_scheduling_proto_init() }
func file_scheduling_proto_init() {
	if File_scheduling_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scheduling_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversationMember); i {
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
		file_scheduling_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseReq); i {
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
		file_scheduling_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseResp); i {
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
			RawDescriptor: file_scheduling_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scheduling_proto_goTypes,
		DependencyIndexes: file_scheduling_proto_depIdxs,
		MessageInfos:      file_scheduling_proto_msgTypes,
	}.Build()
	File_scheduling_proto = out.File
	file_scheduling_proto_rawDesc = nil
	file_scheduling_proto_goTypes = nil
	file_scheduling_proto_depIdxs = nil
}
