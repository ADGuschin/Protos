// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.1
// source: logger.proto

package Logger

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

// Enum which cosist of services
type Source int32

const (
	Source_SCRAPPING_SERVICE           Source = 0
	Source_DATA_PROCESSING_SERVICE     Source = 1
	Source_PROCESSING_DATABASE_SERVICE Source = 2
	Source_SEARCH_SERVICE              Source = 3
	Source_RECOMMENDATION_SERVICE      Source = 4
	Source_LOGGER_SERVICE              Source = 5
)

// Enum value maps for Source.
var (
	Source_name = map[int32]string{
		0: "SCRAPPING_SERVICE",
		1: "DATA_PROCESSING_SERVICE",
		2: "PROCESSING_DATABASE_SERVICE",
		3: "SEARCH_SERVICE",
		4: "RECOMMENDATION_SERVICE",
		5: "LOGGER_SERVICE",
	}
	Source_value = map[string]int32{
		"SCRAPPING_SERVICE":           0,
		"DATA_PROCESSING_SERVICE":     1,
		"PROCESSING_DATABASE_SERVICE": 2,
		"SEARCH_SERVICE":              3,
		"RECOMMENDATION_SERVICE":      4,
		"LOGGER_SERVICE":              5,
	}
)

func (x Source) Enum() *Source {
	p := new(Source)
	*p = x
	return p
}

func (x Source) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Source) Descriptor() protoreflect.EnumDescriptor {
	return file_logger_proto_enumTypes[0].Descriptor()
}

func (Source) Type() protoreflect.EnumType {
	return &file_logger_proto_enumTypes[0]
}

func (x Source) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Source.Descriptor instead.
func (Source) EnumDescriptor() ([]byte, []int) {
	return file_logger_proto_rawDescGZIP(), []int{0}
}

type MessageType int32

const (
	MessageType_INFO    MessageType = 0
	MessageType_SERVICE MessageType = 1
	MessageType_WARNING MessageType = 2
	MessageType_ERROR   MessageType = 3
	MessageType_DEBUG   MessageType = 4
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "INFO",
		1: "SERVICE",
		2: "WARNING",
		3: "ERROR",
		4: "DEBUG",
	}
	MessageType_value = map[string]int32{
		"INFO":    0,
		"SERVICE": 1,
		"WARNING": 2,
		"ERROR":   3,
		"DEBUG":   4,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_logger_proto_enumTypes[1].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_logger_proto_enumTypes[1]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_logger_proto_rawDescGZIP(), []int{1}
}

type LogFormatRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogFormatRequest) Reset() {
	*x = LogFormatRequest{}
	mi := &file_logger_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogFormatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogFormatRequest) ProtoMessage() {}

func (x *LogFormatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logger_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogFormatRequest.ProtoReflect.Descriptor instead.
func (*LogFormatRequest) Descriptor() ([]byte, []int) {
	return file_logger_proto_rawDescGZIP(), []int{0}
}

type LogFormatResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LogFormat     string                 `protobuf:"bytes,1,opt,name=log_format,json=logFormat,proto3" json:"log_format,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogFormatResponse) Reset() {
	*x = LogFormatResponse{}
	mi := &file_logger_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogFormatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogFormatResponse) ProtoMessage() {}

func (x *LogFormatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logger_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogFormatResponse.ProtoReflect.Descriptor instead.
func (*LogFormatResponse) Descriptor() ([]byte, []int) {
	return file_logger_proto_rawDescGZIP(), []int{1}
}

func (x *LogFormatResponse) GetLogFormat() string {
	if x != nil {
		return x.LogFormat
	}
	return ""
}

type LogMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          MessageType            `protobuf:"varint,1,opt,name=type,proto3,enum=Dressy.Logger.MessageType" json:"type,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Source        Source                 `protobuf:"varint,3,opt,name=source,proto3,enum=Dressy.Logger.Source" json:"source,omitempty"`
	Content       string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogMessage) Reset() {
	*x = LogMessage{}
	mi := &file_logger_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogMessage) ProtoMessage() {}

func (x *LogMessage) ProtoReflect() protoreflect.Message {
	mi := &file_logger_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogMessage.ProtoReflect.Descriptor instead.
func (*LogMessage) Descriptor() ([]byte, []int) {
	return file_logger_proto_rawDescGZIP(), []int{2}
}

func (x *LogMessage) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_INFO
}

func (x *LogMessage) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *LogMessage) GetSource() Source {
	if x != nil {
		return x.Source
	}
	return Source_SCRAPPING_SERVICE
}

func (x *LogMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type LogResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsRecorded    bool                   `protobuf:"varint,1,opt,name=is_recorded,json=isRecorded,proto3" json:"is_recorded,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogResult) Reset() {
	*x = LogResult{}
	mi := &file_logger_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogResult) ProtoMessage() {}

func (x *LogResult) ProtoReflect() protoreflect.Message {
	mi := &file_logger_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogResult.ProtoReflect.Descriptor instead.
func (*LogResult) Descriptor() ([]byte, []int) {
	return file_logger_proto_rawDescGZIP(), []int{3}
}

func (x *LogResult) GetIsRecorded() bool {
	if x != nil {
		return x.IsRecorded
	}
	return false
}

var File_logger_proto protoreflect.FileDescriptor

var file_logger_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x44, 0x72, 0x65, 0x73, 0x73, 0x79, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x12,
	0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x32, 0x0a, 0x11, 0x4c, 0x6f, 0x67, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x5f, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0xbf, 0x01, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x44, 0x72, 0x65, 0x73, 0x73, 0x79, 0x2e, 0x4c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x2d, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x44, 0x72, 0x65, 0x73, 0x73, 0x79, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x2a, 0xa1, 0x01, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x43, 0x52, 0x41, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x41, 0x54, 0x41,
	0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x52, 0x56,
	0x49, 0x43, 0x45, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53,
	0x49, 0x4e, 0x47, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48,
	0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45,
	0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x4f, 0x47, 0x47, 0x45, 0x52,
	0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x05, 0x2a, 0x47, 0x0a, 0x0b, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46,
	0x4f, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42, 0x55,
	0x47, 0x10, 0x04, 0x32, 0x97, 0x01, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x12, 0x51,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x1f,
	0x2e, 0x44, 0x72, 0x65, 0x73, 0x73, 0x79, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x4c,
	0x6f, 0x67, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x44, 0x72, 0x65, 0x73, 0x73, 0x79, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e,
	0x4c, 0x6f, 0x67, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3a, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x19, 0x2e, 0x44, 0x72, 0x65, 0x73, 0x73,
	0x79, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x18, 0x2e, 0x44, 0x72, 0x65, 0x73, 0x73, 0x79, 0x2e, 0x4c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x0f, 0x5a,
	0x0d, 0x44, 0x72, 0x65, 0x73, 0x73, 0x79, 0x2f, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_logger_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_logger_proto_goTypes = []any{
	(Source)(0),                   // 0: Dressy.Logger.Source
	(MessageType)(0),              // 1: Dressy.Logger.MessageType
	(*LogFormatRequest)(nil),      // 2: Dressy.Logger.LogFormatRequest
	(*LogFormatResponse)(nil),     // 3: Dressy.Logger.LogFormatResponse
	(*LogMessage)(nil),            // 4: Dressy.Logger.LogMessage
	(*LogResult)(nil),             // 5: Dressy.Logger.LogResult
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_logger_proto_depIdxs = []int32{
	1, // 0: Dressy.Logger.LogMessage.type:type_name -> Dressy.Logger.MessageType
	6, // 1: Dressy.Logger.LogMessage.timestamp:type_name -> google.protobuf.Timestamp
	0, // 2: Dressy.Logger.LogMessage.source:type_name -> Dressy.Logger.Source
	2, // 3: Dressy.Logger.Logger.GetLogFormat:input_type -> Dressy.Logger.LogFormatRequest
	4, // 4: Dressy.Logger.Logger.Log:input_type -> Dressy.Logger.LogMessage
	3, // 5: Dressy.Logger.Logger.GetLogFormat:output_type -> Dressy.Logger.LogFormatResponse
	5, // 6: Dressy.Logger.Logger.Log:output_type -> Dressy.Logger.LogResult
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_logger_proto_init() }
func file_logger_proto_init() {
	if File_logger_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_logger_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
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
