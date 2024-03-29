// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: stroeer/curation/v1/core_curation_service.proto

//*
// @FileArticle ⚙︎ CurationService

package core

import (
	v1 "github.com/stroeer/go-tapir/core/v1"
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

// *
// # `⚙︎ GetCuration`
//
// Fetch a curation by its id and return the [`repeated stroeer.core.v1.Article`](Article.html) this
// curation contains. The response may be empty in case the curation does not contain any items.
//
// a `NOT_FOUND` status code will indicate the curation `id` does not exist.
//
// ## GetCurationRequest
//
// | Field name       | Type     | Description                                                 |
// |------------------|----------|-------------------------------------------------------------|
// | `id`             | `int64`  | [required] id of the list to be fetched                     |
//
// @CodeBlockStart protobuf
type GetCurationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCurationRequest) Reset() {
	*x = GetCurationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_curation_v1_core_curation_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurationRequest) ProtoMessage() {}

func (x *GetCurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_curation_v1_core_curation_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurationRequest.ProtoReflect.Descriptor instead.
func (*GetCurationRequest) Descriptor() ([]byte, []int) {
	return file_stroeer_curation_v1_core_curation_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetCurationRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCurationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Label      string                 `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	Articles   []*v1.Article          `protobuf:"bytes,4,rep,name=articles,proto3" json:"articles,omitempty"`
}

func (x *GetCurationResponse) Reset() {
	*x = GetCurationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_curation_v1_core_curation_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurationResponse) ProtoMessage() {}

func (x *GetCurationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_curation_v1_core_curation_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurationResponse.ProtoReflect.Descriptor instead.
func (*GetCurationResponse) Descriptor() ([]byte, []int) {
	return file_stroeer_curation_v1_core_curation_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetCurationResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetCurationResponse) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *GetCurationResponse) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *GetCurationResponse) GetArticles() []*v1.Article {
	if x != nil {
		return x.Articles
	}
	return nil
}

// *
// # `⚙︎ BatchGetCuration`
//
// Fetch multiple curations by their id and return the [`repeated stroeer.core.v1.Article`](Article.html) those
// curations contain. The response may be empty in case the curation does not contain any items.
// The ordering of items will the same ordering as the `ids` requested.
//
// ## BatchGetCurationRequest
//
// | Field name       | Type                                | Description                          |
// |------------------|-------------------------------------|--------------------------------------|
// | `ids`            | `repeated int64`                    | the _ids_ of the lists to be fetched |
//
// @CodeBlockStart protobuf
type BatchGetCurationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *BatchGetCurationRequest) Reset() {
	*x = BatchGetCurationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_curation_v1_core_curation_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetCurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetCurationRequest) ProtoMessage() {}

func (x *BatchGetCurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_curation_v1_core_curation_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetCurationRequest.ProtoReflect.Descriptor instead.
func (*BatchGetCurationRequest) Descriptor() ([]byte, []int) {
	return file_stroeer_curation_v1_core_curation_service_proto_rawDescGZIP(), []int{2}
}

func (x *BatchGetCurationRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

// *
// ## BatchGetCurationResponse
//
// | Field name       | Type                                | Description                                                                    |
// |------------------|-------------------------------------|--------------------------------------------------------------------------------|
// | `curations`      | [`GetCurationResponse`][cr]         | a single response item that corresponds to _ids_ this service was called with. |
//
// [cr]: #getcurationresponse
//
// @CodeBlockStart protobuf
type BatchGetCurationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Curations []*GetCurationResponse `protobuf:"bytes,1,rep,name=curations,proto3" json:"curations,omitempty"`
}

func (x *BatchGetCurationResponse) Reset() {
	*x = BatchGetCurationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_curation_v1_core_curation_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetCurationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetCurationResponse) ProtoMessage() {}

func (x *BatchGetCurationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_curation_v1_core_curation_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetCurationResponse.ProtoReflect.Descriptor instead.
func (*BatchGetCurationResponse) Descriptor() ([]byte, []int) {
	return file_stroeer_curation_v1_core_curation_service_proto_rawDescGZIP(), []int{3}
}

func (x *BatchGetCurationResponse) GetCurations() []*GetCurationResponse {
	if x != nil {
		return x.Curations
	}
	return nil
}

var File_stroeer_curation_v1_core_curation_service_proto protoreflect.FileDescriptor

var file_stroeer_curation_v1_core_curation_service_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x63, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xae, 0x01, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x72, 0x6f,
	0x65, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x22, 0x2b, 0x0a,
	0x17, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x62, 0x0a, 0x18, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x74, 0x72, 0x6f,
	0x65, 0x65, 0x72, 0x2e, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x09, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xe8,
	0x01, 0x0a, 0x0f, 0x43, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x62, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x27, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x63, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x74, 0x72,
	0x6f, 0x65, 0x65, 0x72, 0x2e, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x71, 0x0a, 0x10, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47,
	0x65, 0x74, 0x43, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x73, 0x74, 0x72,
	0x6f, 0x65, 0x65, 0x72, 0x2e, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65,
	0x65, 0x72, 0x2e, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f,
	0x67, 0x6f, 0x2d, 0x74, 0x61, 0x70, 0x69, 0x72, 0x2f, 0x63, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_stroeer_curation_v1_core_curation_service_proto_rawDescOnce sync.Once
	file_stroeer_curation_v1_core_curation_service_proto_rawDescData = file_stroeer_curation_v1_core_curation_service_proto_rawDesc
)

func file_stroeer_curation_v1_core_curation_service_proto_rawDescGZIP() []byte {
	file_stroeer_curation_v1_core_curation_service_proto_rawDescOnce.Do(func() {
		file_stroeer_curation_v1_core_curation_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_stroeer_curation_v1_core_curation_service_proto_rawDescData)
	})
	return file_stroeer_curation_v1_core_curation_service_proto_rawDescData
}

var file_stroeer_curation_v1_core_curation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_stroeer_curation_v1_core_curation_service_proto_goTypes = []interface{}{
	(*GetCurationRequest)(nil),       // 0: stroeer.curation.v1.GetCurationRequest
	(*GetCurationResponse)(nil),      // 1: stroeer.curation.v1.GetCurationResponse
	(*BatchGetCurationRequest)(nil),  // 2: stroeer.curation.v1.BatchGetCurationRequest
	(*BatchGetCurationResponse)(nil), // 3: stroeer.curation.v1.BatchGetCurationResponse
	(*timestamppb.Timestamp)(nil),    // 4: google.protobuf.Timestamp
	(*v1.Article)(nil),               // 5: stroeer.core.v1.Article
}
var file_stroeer_curation_v1_core_curation_service_proto_depIdxs = []int32{
	4, // 0: stroeer.curation.v1.GetCurationResponse.update_time:type_name -> google.protobuf.Timestamp
	5, // 1: stroeer.curation.v1.GetCurationResponse.articles:type_name -> stroeer.core.v1.Article
	1, // 2: stroeer.curation.v1.BatchGetCurationResponse.curations:type_name -> stroeer.curation.v1.GetCurationResponse
	0, // 3: stroeer.curation.v1.CurationService.GetCuration:input_type -> stroeer.curation.v1.GetCurationRequest
	2, // 4: stroeer.curation.v1.CurationService.BatchGetCuration:input_type -> stroeer.curation.v1.BatchGetCurationRequest
	1, // 5: stroeer.curation.v1.CurationService.GetCuration:output_type -> stroeer.curation.v1.GetCurationResponse
	3, // 6: stroeer.curation.v1.CurationService.BatchGetCuration:output_type -> stroeer.curation.v1.BatchGetCurationResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_stroeer_curation_v1_core_curation_service_proto_init() }
func file_stroeer_curation_v1_core_curation_service_proto_init() {
	if File_stroeer_curation_v1_core_curation_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stroeer_curation_v1_core_curation_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurationRequest); i {
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
		file_stroeer_curation_v1_core_curation_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurationResponse); i {
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
		file_stroeer_curation_v1_core_curation_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetCurationRequest); i {
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
		file_stroeer_curation_v1_core_curation_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetCurationResponse); i {
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
			RawDescriptor: file_stroeer_curation_v1_core_curation_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stroeer_curation_v1_core_curation_service_proto_goTypes,
		DependencyIndexes: file_stroeer_curation_v1_core_curation_service_proto_depIdxs,
		MessageInfos:      file_stroeer_curation_v1_core_curation_service_proto_msgTypes,
	}.Build()
	File_stroeer_curation_v1_core_curation_service_proto = out.File
	file_stroeer_curation_v1_core_curation_service_proto_rawDesc = nil
	file_stroeer_curation_v1_core_curation_service_proto_goTypes = nil
	file_stroeer_curation_v1_core_curation_service_proto_depIdxs = nil
}
