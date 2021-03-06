// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: M1/main/bookStore.proto

package bytesT

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

type UploadBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileChunk    []byte `protobuf:"bytes,1,opt,name=fileChunk,proto3" json:"fileChunk,omitempty"`
	BookName     string `protobuf:"bytes,2,opt,name=bookName,proto3" json:"bookName,omitempty"`
	ProposalType string `protobuf:"bytes,3,opt,name=proposalType,proto3" json:"proposalType,omitempty"`
	BookSize     uint64 `protobuf:"varint,4,opt,name=bookSize,proto3" json:"bookSize,omitempty"`
	NPart        uint64 `protobuf:"varint,5,opt,name=nPart,proto3" json:"nPart,omitempty"`
	APart        uint64 `protobuf:"varint,6,opt,name=aPart,proto3" json:"aPart,omitempty"`
}

func (x *UploadBookRequest) Reset() {
	*x = UploadBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_M1_main_bookStore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadBookRequest) ProtoMessage() {}

func (x *UploadBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_M1_main_bookStore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadBookRequest.ProtoReflect.Descriptor instead.
func (*UploadBookRequest) Descriptor() ([]byte, []int) {
	return file_M1_main_bookStore_proto_rawDescGZIP(), []int{0}
}

func (x *UploadBookRequest) GetFileChunk() []byte {
	if x != nil {
		return x.FileChunk
	}
	return nil
}

func (x *UploadBookRequest) GetBookName() string {
	if x != nil {
		return x.BookName
	}
	return ""
}

func (x *UploadBookRequest) GetProposalType() string {
	if x != nil {
		return x.ProposalType
	}
	return ""
}

func (x *UploadBookRequest) GetBookSize() uint64 {
	if x != nil {
		return x.BookSize
	}
	return 0
}

func (x *UploadBookRequest) GetNPart() uint64 {
	if x != nil {
		return x.NPart
	}
	return 0
}

func (x *UploadBookRequest) GetAPart() uint64 {
	if x != nil {
		return x.APart
	}
	return 0
}

type DistributeBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileChunk []byte `protobuf:"bytes,1,opt,name=fileChunk,proto3" json:"fileChunk,omitempty"`
	ChunkName string `protobuf:"bytes,2,opt,name=chunkName,proto3" json:"chunkName,omitempty"`
}

func (x *DistributeBookRequest) Reset() {
	*x = DistributeBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_M1_main_bookStore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributeBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributeBookRequest) ProtoMessage() {}

func (x *DistributeBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_M1_main_bookStore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributeBookRequest.ProtoReflect.Descriptor instead.
func (*DistributeBookRequest) Descriptor() ([]byte, []int) {
	return file_M1_main_bookStore_proto_rawDescGZIP(), []int{1}
}

func (x *DistributeBookRequest) GetFileChunk() []byte {
	if x != nil {
		return x.FileChunk
	}
	return nil
}

func (x *DistributeBookRequest) GetChunkName() string {
	if x != nil {
		return x.ChunkName
	}
	return ""
}

type RnARequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	T         string  `protobuf:"bytes,1,opt,name=T,proto3" json:"T,omitempty"`
	State     string  `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	FileChunk []byte  `protobuf:"bytes,3,opt,name=fileChunk,proto3,oneof" json:"fileChunk,omitempty"`
	ChunkName *string `protobuf:"bytes,4,opt,name=chunkName,proto3,oneof" json:"chunkName,omitempty"`
}

func (x *RnARequest) Reset() {
	*x = RnARequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_M1_main_bookStore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RnARequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RnARequest) ProtoMessage() {}

func (x *RnARequest) ProtoReflect() protoreflect.Message {
	mi := &file_M1_main_bookStore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RnARequest.ProtoReflect.Descriptor instead.
func (*RnARequest) Descriptor() ([]byte, []int) {
	return file_M1_main_bookStore_proto_rawDescGZIP(), []int{2}
}

func (x *RnARequest) GetT() string {
	if x != nil {
		return x.T
	}
	return ""
}

func (x *RnARequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *RnARequest) GetFileChunk() []byte {
	if x != nil {
		return x.FileChunk
	}
	return nil
}

func (x *RnARequest) GetChunkName() string {
	if x != nil && x.ChunkName != nil {
		return *x.ChunkName
	}
	return ""
}

type LogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookName string   `protobuf:"bytes,1,opt,name=bookName,proto3" json:"bookName,omitempty"`
	NPart    int64    `protobuf:"varint,2,opt,name=nPart,proto3" json:"nPart,omitempty"`
	PartPos  []string `protobuf:"bytes,3,rep,name=partPos,proto3" json:"partPos,omitempty"`
}

func (x *LogRequest) Reset() {
	*x = LogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_M1_main_bookStore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRequest) ProtoMessage() {}

func (x *LogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_M1_main_bookStore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRequest.ProtoReflect.Descriptor instead.
func (*LogRequest) Descriptor() ([]byte, []int) {
	return file_M1_main_bookStore_proto_rawDescGZIP(), []int{3}
}

func (x *LogRequest) GetBookName() string {
	if x != nil {
		return x.BookName
	}
	return ""
}

func (x *LogRequest) GetNPart() int64 {
	if x != nil {
		return x.NPart
	}
	return 0
}

func (x *LogRequest) GetPartPos() []string {
	if x != nil {
		return x.PartPos
	}
	return nil
}

type GetChunkLocRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookName string `protobuf:"bytes,1,opt,name=bookName,proto3" json:"bookName,omitempty"`
}

func (x *GetChunkLocRequest) Reset() {
	*x = GetChunkLocRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_M1_main_bookStore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChunkLocRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChunkLocRequest) ProtoMessage() {}

func (x *GetChunkLocRequest) ProtoReflect() protoreflect.Message {
	mi := &file_M1_main_bookStore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChunkLocRequest.ProtoReflect.Descriptor instead.
func (*GetChunkLocRequest) Descriptor() ([]byte, []int) {
	return file_M1_main_bookStore_proto_rawDescGZIP(), []int{4}
}

func (x *GetChunkLocRequest) GetBookName() string {
	if x != nil {
		return x.BookName
	}
	return ""
}

type GetChunkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkName string `protobuf:"bytes,1,opt,name=chunkName,proto3" json:"chunkName,omitempty"`
}

func (x *GetChunkRequest) Reset() {
	*x = GetChunkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_M1_main_bookStore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChunkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChunkRequest) ProtoMessage() {}

func (x *GetChunkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_M1_main_bookStore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChunkRequest.ProtoReflect.Descriptor instead.
func (*GetChunkRequest) Descriptor() ([]byte, []int) {
	return file_M1_main_bookStore_proto_rawDescGZIP(), []int{5}
}

func (x *GetChunkRequest) GetChunkName() string {
	if x != nil {
		return x.ChunkName
	}
	return ""
}

type UploadBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size uint32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *UploadBookResponse) Reset() {
	*x = UploadBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_M1_main_bookStore_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadBookResponse) ProtoMessage() {}

func (x *UploadBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_M1_main_bookStore_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadBookResponse.ProtoReflect.Descriptor instead.
func (*UploadBookResponse) Descriptor() ([]byte, []int) {
	return file_M1_main_bookStore_proto_rawDescGZIP(), []int{6}
}

func (x *UploadBookResponse) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type DistributeBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *DistributeBookResponse) Reset() {
	*x = DistributeBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_M1_main_bookStore_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributeBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributeBookResponse) ProtoMessage() {}

func (x *DistributeBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_M1_main_bookStore_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributeBookResponse.ProtoReflect.Descriptor instead.
func (*DistributeBookResponse) Descriptor() ([]byte, []int) {
	return file_M1_main_bookStore_proto_rawDescGZIP(), []int{7}
}

func (x *DistributeBookResponse) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type RnAResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alive string  `protobuf:"bytes,1,opt,name=alive,proto3" json:"alive,omitempty"`
	State *string `protobuf:"bytes,2,opt,name=state,proto3,oneof" json:"state,omitempty"`
}

func (x *RnAResponse) Reset() {
	*x = RnAResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_M1_main_bookStore_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RnAResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RnAResponse) ProtoMessage() {}

func (x *RnAResponse) ProtoReflect() protoreflect.Message {
	mi := &file_M1_main_bookStore_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RnAResponse.ProtoReflect.Descriptor instead.
func (*RnAResponse) Descriptor() ([]byte, []int) {
	return file_M1_main_bookStore_proto_rawDescGZIP(), []int{8}
}

func (x *RnAResponse) GetAlive() string {
	if x != nil {
		return x.Alive
	}
	return ""
}

func (x *RnAResponse) GetState() string {
	if x != nil && x.State != nil {
		return *x.State
	}
	return ""
}

type LogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *LogResponse) Reset() {
	*x = LogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_M1_main_bookStore_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogResponse) ProtoMessage() {}

func (x *LogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_M1_main_bookStore_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogResponse.ProtoReflect.Descriptor instead.
func (*LogResponse) Descriptor() ([]byte, []int) {
	return file_M1_main_bookStore_proto_rawDescGZIP(), []int{9}
}

func (x *LogResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type GetChunkLocResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location []string `protobuf:"bytes,1,rep,name=location,proto3" json:"location,omitempty"`
}

func (x *GetChunkLocResponse) Reset() {
	*x = GetChunkLocResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_M1_main_bookStore_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChunkLocResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChunkLocResponse) ProtoMessage() {}

func (x *GetChunkLocResponse) ProtoReflect() protoreflect.Message {
	mi := &file_M1_main_bookStore_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChunkLocResponse.ProtoReflect.Descriptor instead.
func (*GetChunkLocResponse) Descriptor() ([]byte, []int) {
	return file_M1_main_bookStore_proto_rawDescGZIP(), []int{10}
}

func (x *GetChunkLocResponse) GetLocation() []string {
	if x != nil {
		return x.Location
	}
	return nil
}

type GetChunkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileChunk []byte `protobuf:"bytes,1,opt,name=fileChunk,proto3" json:"fileChunk,omitempty"`
}

func (x *GetChunkResponse) Reset() {
	*x = GetChunkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_M1_main_bookStore_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChunkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChunkResponse) ProtoMessage() {}

func (x *GetChunkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_M1_main_bookStore_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChunkResponse.ProtoReflect.Descriptor instead.
func (*GetChunkResponse) Descriptor() ([]byte, []int) {
	return file_M1_main_bookStore_proto_rawDescGZIP(), []int{11}
}

func (x *GetChunkResponse) GetFileChunk() []byte {
	if x != nil {
		return x.FileChunk
	}
	return nil
}

var File_M1_main_bookStore_proto protoreflect.FileDescriptor

var file_M1_main_bookStore_proto_rawDesc = []byte{
	0x0a, 0x17, 0x4d, 0x31, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x54, 0x22, 0xb9, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x50, 0x61, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x6e, 0x50, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x50, 0x61, 0x72, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x61, 0x50, 0x61, 0x72, 0x74, 0x22, 0x53, 0x0a,
	0x15, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x92, 0x01, 0x0a, 0x0a, 0x52, 0x6e, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0c, 0x0a, 0x01, 0x54, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x54, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x58, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x50, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6e, 0x50, 0x61, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x50,
	0x6f, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x50, 0x6f,
	0x73, 0x22, 0x30, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x2e,
	0x0a, 0x16, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x48,
	0x0a, 0x0b, 0x52, 0x6e, 0x41, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c,
	0x69, 0x76, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x29, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x31, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c,
	0x6f, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x30, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x32, 0xac, 0x03, 0x0a, 0x0e, 0x4c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x19, 0x2e, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x54, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x54, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x12, 0x51, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1d, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x54, 0x2e,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x54, 0x2e, 0x44,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0f, 0x52, 0x69, 0x63, 0x61, 0x72,
	0x74, 0x4e, 0x41, 0x67, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x61, 0x12, 0x12, 0x2e, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x54, 0x2e, 0x52, 0x6e, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x54, 0x2e, 0x52, 0x6e, 0x41, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x33, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x4d, 0x6f,
	0x64, 0x12, 0x12, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x54, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x54, 0x2e, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x63, 0x12, 0x1a, 0x2e, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x54, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x54,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x12, 0x17, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x54, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x54, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_M1_main_bookStore_proto_rawDescOnce sync.Once
	file_M1_main_bookStore_proto_rawDescData = file_M1_main_bookStore_proto_rawDesc
)

func file_M1_main_bookStore_proto_rawDescGZIP() []byte {
	file_M1_main_bookStore_proto_rawDescOnce.Do(func() {
		file_M1_main_bookStore_proto_rawDescData = protoimpl.X.CompressGZIP(file_M1_main_bookStore_proto_rawDescData)
	})
	return file_M1_main_bookStore_proto_rawDescData
}

var file_M1_main_bookStore_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_M1_main_bookStore_proto_goTypes = []interface{}{
	(*UploadBookRequest)(nil),      // 0: bytesT.UploadBookRequest
	(*DistributeBookRequest)(nil),  // 1: bytesT.DistributeBookRequest
	(*RnARequest)(nil),             // 2: bytesT.RnARequest
	(*LogRequest)(nil),             // 3: bytesT.LogRequest
	(*GetChunkLocRequest)(nil),     // 4: bytesT.GetChunkLocRequest
	(*GetChunkRequest)(nil),        // 5: bytesT.GetChunkRequest
	(*UploadBookResponse)(nil),     // 6: bytesT.UploadBookResponse
	(*DistributeBookResponse)(nil), // 7: bytesT.DistributeBookResponse
	(*RnAResponse)(nil),            // 8: bytesT.RnAResponse
	(*LogResponse)(nil),            // 9: bytesT.LogResponse
	(*GetChunkLocResponse)(nil),    // 10: bytesT.GetChunkLocResponse
	(*GetChunkResponse)(nil),       // 11: bytesT.GetChunkResponse
}
var file_M1_main_bookStore_proto_depIdxs = []int32{
	0,  // 0: bytesT.LibraryService.UploadBook:input_type -> bytesT.UploadBookRequest
	1,  // 1: bytesT.LibraryService.DistributeBook:input_type -> bytesT.DistributeBookRequest
	2,  // 2: bytesT.LibraryService.RicartNAgrawala:input_type -> bytesT.RnARequest
	3,  // 3: bytesT.LibraryService.LogMod:input_type -> bytesT.LogRequest
	4,  // 4: bytesT.LibraryService.GetChunkLoc:input_type -> bytesT.GetChunkLocRequest
	5,  // 5: bytesT.LibraryService.GetChunk:input_type -> bytesT.GetChunkRequest
	6,  // 6: bytesT.LibraryService.UploadBook:output_type -> bytesT.UploadBookResponse
	7,  // 7: bytesT.LibraryService.DistributeBook:output_type -> bytesT.DistributeBookResponse
	8,  // 8: bytesT.LibraryService.RicartNAgrawala:output_type -> bytesT.RnAResponse
	9,  // 9: bytesT.LibraryService.LogMod:output_type -> bytesT.LogResponse
	10, // 10: bytesT.LibraryService.GetChunkLoc:output_type -> bytesT.GetChunkLocResponse
	11, // 11: bytesT.LibraryService.GetChunk:output_type -> bytesT.GetChunkResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_M1_main_bookStore_proto_init() }
func file_M1_main_bookStore_proto_init() {
	if File_M1_main_bookStore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_M1_main_bookStore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadBookRequest); i {
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
		file_M1_main_bookStore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributeBookRequest); i {
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
		file_M1_main_bookStore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RnARequest); i {
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
		file_M1_main_bookStore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogRequest); i {
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
		file_M1_main_bookStore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChunkLocRequest); i {
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
		file_M1_main_bookStore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChunkRequest); i {
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
		file_M1_main_bookStore_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadBookResponse); i {
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
		file_M1_main_bookStore_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributeBookResponse); i {
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
		file_M1_main_bookStore_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RnAResponse); i {
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
		file_M1_main_bookStore_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogResponse); i {
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
		file_M1_main_bookStore_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChunkLocResponse); i {
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
		file_M1_main_bookStore_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChunkResponse); i {
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
	file_M1_main_bookStore_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_M1_main_bookStore_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_M1_main_bookStore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_M1_main_bookStore_proto_goTypes,
		DependencyIndexes: file_M1_main_bookStore_proto_depIdxs,
		MessageInfos:      file_M1_main_bookStore_proto_msgTypes,
	}.Build()
	File_M1_main_bookStore_proto = out.File
	file_M1_main_bookStore_proto_rawDesc = nil
	file_M1_main_bookStore_proto_goTypes = nil
	file_M1_main_bookStore_proto_depIdxs = nil
}
