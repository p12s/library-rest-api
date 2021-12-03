// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: logger.proto

package pb

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

type LoggerRequest_Actions int32

const (
	LoggerRequest_REGISTER LoggerRequest_Actions = 0
	LoggerRequest_LOGIN    LoggerRequest_Actions = 1
	LoggerRequest_CREATE   LoggerRequest_Actions = 2
	LoggerRequest_UPDATE   LoggerRequest_Actions = 3
	LoggerRequest_GET      LoggerRequest_Actions = 4
	LoggerRequest_DELETE   LoggerRequest_Actions = 5
)

// Enum value maps for LoggerRequest_Actions.
var (
	LoggerRequest_Actions_name = map[int32]string{
		0: "REGISTER",
		1: "LOGIN",
		2: "CREATE",
		3: "UPDATE",
		4: "GET",
		5: "DELETE",
	}
	LoggerRequest_Actions_value = map[string]int32{
		"REGISTER": 0,
		"LOGIN":    1,
		"CREATE":   2,
		"UPDATE":   3,
		"GET":      4,
		"DELETE":   5,
	}
)

func (x LoggerRequest_Actions) Enum() *LoggerRequest_Actions {
	p := new(LoggerRequest_Actions)
	*p = x
	return p
}

func (x LoggerRequest_Actions) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoggerRequest_Actions) Descriptor() protoreflect.EnumDescriptor {
	return file_logger_proto_enumTypes[0].Descriptor()
}

func (LoggerRequest_Actions) Type() protoreflect.EnumType {
	return &file_logger_proto_enumTypes[0]
}

func (x LoggerRequest_Actions) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoggerRequest_Actions.Descriptor instead.
func (LoggerRequest_Actions) EnumDescriptor() ([]byte, []int) {
	return file_logger_proto_rawDescGZIP(), []int{0, 0}
}

type LoggerRequest_Entities int32

const (
	LoggerRequest_USER LoggerRequest_Entities = 0
	LoggerRequest_BOOK LoggerRequest_Entities = 1
)

// Enum value maps for LoggerRequest_Entities.
var (
	LoggerRequest_Entities_name = map[int32]string{
		0: "USER",
		1: "BOOK",
	}
	LoggerRequest_Entities_value = map[string]int32{
		"USER": 0,
		"BOOK": 1,
	}
)

func (x LoggerRequest_Entities) Enum() *LoggerRequest_Entities {
	p := new(LoggerRequest_Entities)
	*p = x
	return p
}

func (x LoggerRequest_Entities) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoggerRequest_Entities) Descriptor() protoreflect.EnumDescriptor {
	return file_logger_proto_enumTypes[1].Descriptor()
}

func (LoggerRequest_Entities) Type() protoreflect.EnumType {
	return &file_logger_proto_enumTypes[1]
}

func (x LoggerRequest_Entities) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoggerRequest_Entities.Descriptor instead.
func (LoggerRequest_Entities) EnumDescriptor() ([]byte, []int) {
	return file_logger_proto_rawDescGZIP(), []int{0, 1}
}

type LoggerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action    LoggerRequest_Actions  `protobuf:"varint,1,opt,name=action,proto3,enum=logger.LoggerRequest_Actions" json:"action,omitempty"`
	Entity    LoggerRequest_Entities `protobuf:"varint,2,opt,name=entity,proto3,enum=logger.LoggerRequest_Entities" json:"entity,omitempty"`
	EntityId  int64                  `protobuf:"varint,3,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *LoggerRequest) Reset() {
	*x = LoggerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoggerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoggerRequest) ProtoMessage() {}

func (x *LoggerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoggerRequest.ProtoReflect.Descriptor instead.
func (*LoggerRequest) Descriptor() ([]byte, []int) {
	return file_logger_proto_rawDescGZIP(), []int{0}
}

func (x *LoggerRequest) GetAction() LoggerRequest_Actions {
	if x != nil {
		return x.Action
	}
	return LoggerRequest_REGISTER
}

func (x *LoggerRequest) GetEntity() LoggerRequest_Entities {
	if x != nil {
		return x.Entity
	}
	return LoggerRequest_USER
}

func (x *LoggerRequest) GetEntityId() int64 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *LoggerRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_logger_proto_rawDescGZIP(), []int{1}
}

var File_logger_proto protoreflect.FileDescriptor

var file_logger_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x02, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x36, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1e, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22,
	0x4f, 0x0a, 0x07, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45,
	0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x47, 0x49,
	0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12,
	0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x47,
	0x45, 0x54, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x05,
	0x22, 0x1e, 0x0a, 0x08, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x08, 0x0a, 0x04,
	0x55, 0x53, 0x45, 0x52, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4f, 0x4f, 0x4b, 0x10, 0x01,
	0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x46, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x35, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x15, 0x2e, 0x6c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logger_proto_rawDescOnce sync.Once
	file_logger_proto_rawDescData = file_logger_proto_rawDesc
)

func file_logger_proto_rawDescGZIP() []byte {
	file_logger_proto_rawDescOnce.Do(func() {
		file_logger_proto_rawDescData = protoimpl.X.CompressGZIP(file_logger_proto_rawDescData)
	})
	return file_logger_proto_rawDescData
}

var file_logger_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_logger_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_logger_proto_goTypes = []interface{}{
	(LoggerRequest_Actions)(0),    // 0: logger.LoggerRequest.Actions
	(LoggerRequest_Entities)(0),   // 1: logger.LoggerRequest.Entities
	(*LoggerRequest)(nil),         // 2: logger.LoggerRequest
	(*EmptyResponse)(nil),         // 3: logger.EmptyResponse
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_logger_proto_depIdxs = []int32{
	0, // 0: logger.LoggerRequest.action:type_name -> logger.LoggerRequest.Actions
	1, // 1: logger.LoggerRequest.entity:type_name -> logger.LoggerRequest.Entities
	4, // 2: logger.LoggerRequest.timestamp:type_name -> google.protobuf.Timestamp
	2, // 3: logger.LoggerService.Log:input_type -> logger.LoggerRequest
	3, // 4: logger.LoggerService.Log:output_type -> logger.EmptyResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_logger_proto_init() }
func file_logger_proto_init() {
	if File_logger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoggerRequest); i {
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
		file_logger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
			RawDescriptor: file_logger_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logger_proto_goTypes,
		DependencyIndexes: file_logger_proto_depIdxs,
		EnumInfos:         file_logger_proto_enumTypes,
		MessageInfos:      file_logger_proto_msgTypes,
	}.Build()
	File_logger_proto = out.File
	file_logger_proto_rawDesc = nil
	file_logger_proto_goTypes = nil
	file_logger_proto_depIdxs = nil
}