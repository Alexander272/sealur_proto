// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: pro/standard_api.proto

package standard_api

import (
	response_model "github.com/Alexander272/sealur_proto/api/pro/models/response_model"
	standard_model "github.com/Alexander272/sealur_proto/api/pro/models/standard_model"
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

type GetAllStandards struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllStandards) Reset() {
	*x = GetAllStandards{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_standard_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllStandards) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllStandards) ProtoMessage() {}

func (x *GetAllStandards) ProtoReflect() protoreflect.Message {
	mi := &file_pro_standard_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllStandards.ProtoReflect.Descriptor instead.
func (*GetAllStandards) Descriptor() ([]byte, []int) {
	return file_pro_standard_api_proto_rawDescGZIP(), []int{0}
}

type CreateStandard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Format []string `protobuf:"bytes,2,rep,name=format,proto3" json:"format,omitempty"`
}

func (x *CreateStandard) Reset() {
	*x = CreateStandard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_standard_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStandard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStandard) ProtoMessage() {}

func (x *CreateStandard) ProtoReflect() protoreflect.Message {
	mi := &file_pro_standard_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStandard.ProtoReflect.Descriptor instead.
func (*CreateStandard) Descriptor() ([]byte, []int) {
	return file_pro_standard_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateStandard) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateStandard) GetFormat() []string {
	if x != nil {
		return x.Format
	}
	return nil
}

type CreateSeveralStandard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Standards []*CreateStandard `protobuf:"bytes,1,rep,name=standards,proto3" json:"standards,omitempty"`
}

func (x *CreateSeveralStandard) Reset() {
	*x = CreateSeveralStandard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_standard_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSeveralStandard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSeveralStandard) ProtoMessage() {}

func (x *CreateSeveralStandard) ProtoReflect() protoreflect.Message {
	mi := &file_pro_standard_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSeveralStandard.ProtoReflect.Descriptor instead.
func (*CreateSeveralStandard) Descriptor() ([]byte, []int) {
	return file_pro_standard_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSeveralStandard) GetStandards() []*CreateStandard {
	if x != nil {
		return x.Standards
	}
	return nil
}

type UpdateStandard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title  string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Format []string `protobuf:"bytes,3,rep,name=format,proto3" json:"format,omitempty"`
}

func (x *UpdateStandard) Reset() {
	*x = UpdateStandard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_standard_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStandard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStandard) ProtoMessage() {}

func (x *UpdateStandard) ProtoReflect() protoreflect.Message {
	mi := &file_pro_standard_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStandard.ProtoReflect.Descriptor instead.
func (*UpdateStandard) Descriptor() ([]byte, []int) {
	return file_pro_standard_api_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateStandard) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateStandard) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateStandard) GetFormat() []string {
	if x != nil {
		return x.Format
	}
	return nil
}

type DeleteStandard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteStandard) Reset() {
	*x = DeleteStandard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_standard_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStandard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStandard) ProtoMessage() {}

func (x *DeleteStandard) ProtoReflect() protoreflect.Message {
	mi := &file_pro_standard_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStandard.ProtoReflect.Descriptor instead.
func (*DeleteStandard) Descriptor() ([]byte, []int) {
	return file_pro_standard_api_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteStandard) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Standards struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Standards []*standard_model.Standard `protobuf:"bytes,1,rep,name=standards,proto3" json:"standards,omitempty"`
}

func (x *Standards) Reset() {
	*x = Standards{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_standard_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Standards) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Standards) ProtoMessage() {}

func (x *Standards) ProtoReflect() protoreflect.Message {
	mi := &file_pro_standard_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Standards.ProtoReflect.Descriptor instead.
func (*Standards) Descriptor() ([]byte, []int) {
	return file_pro_standard_api_proto_rawDescGZIP(), []int{5}
}

func (x *Standards) GetStandards() []*standard_model.Standard {
	if x != nil {
		return x.Standards
	}
	return nil
}

var File_pro_standard_api_proto protoreflect.FileDescriptor

var file_pro_standard_api_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61,
	0x72, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x22, 0x3e, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x53, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x6e,
	0x64, 0x61, 0x72, 0x64, 0x12, 0x3a, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61,
	0x72, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x6e, 0x64, 0x61, 0x72, 0x64, 0x52, 0x09, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73,
	0x22, 0x4e, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61,
	0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x22, 0x20, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61,
	0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x43, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x12,
	0x36, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x32, 0xe9, 0x02, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x6e,
	0x64, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1d, 0x2e, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x6e, 0x64,
	0x61, 0x72, 0x64, 0x73, 0x1a, 0x17, 0x2e, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x12, 0x40, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61,
	0x72, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x6e, 0x64, 0x61, 0x72, 0x64, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x76, 0x65, 0x72, 0x61, 0x6c,
	0x12, 0x23, 0x2e, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x74, 0x61,
	0x6e, 0x64, 0x61, 0x72, 0x64, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x40, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x74, 0x61, 0x6e,
	0x64, 0x61, 0x72, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x40, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x74,
	0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x2f, 0x73,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pro_standard_api_proto_rawDescOnce sync.Once
	file_pro_standard_api_proto_rawDescData = file_pro_standard_api_proto_rawDesc
)

func file_pro_standard_api_proto_rawDescGZIP() []byte {
	file_pro_standard_api_proto_rawDescOnce.Do(func() {
		file_pro_standard_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_standard_api_proto_rawDescData)
	})
	return file_pro_standard_api_proto_rawDescData
}

var file_pro_standard_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pro_standard_api_proto_goTypes = []interface{}{
	(*GetAllStandards)(nil),         // 0: standard_api.GetAllStandards
	(*CreateStandard)(nil),          // 1: standard_api.CreateStandard
	(*CreateSeveralStandard)(nil),   // 2: standard_api.CreateSeveralStandard
	(*UpdateStandard)(nil),          // 3: standard_api.UpdateStandard
	(*DeleteStandard)(nil),          // 4: standard_api.DeleteStandard
	(*Standards)(nil),               // 5: standard_api.Standards
	(*standard_model.Standard)(nil), // 6: standard_model.Standard
	(*response_model.Response)(nil), // 7: response_model.Response
}
var file_pro_standard_api_proto_depIdxs = []int32{
	1, // 0: standard_api.CreateSeveralStandard.standards:type_name -> standard_api.CreateStandard
	6, // 1: standard_api.Standards.standards:type_name -> standard_model.Standard
	0, // 2: standard_api.StandardService.GetAll:input_type -> standard_api.GetAllStandards
	1, // 3: standard_api.StandardService.Create:input_type -> standard_api.CreateStandard
	2, // 4: standard_api.StandardService.CreateSeveral:input_type -> standard_api.CreateSeveralStandard
	3, // 5: standard_api.StandardService.Update:input_type -> standard_api.UpdateStandard
	4, // 6: standard_api.StandardService.Delete:input_type -> standard_api.DeleteStandard
	5, // 7: standard_api.StandardService.GetAll:output_type -> standard_api.Standards
	7, // 8: standard_api.StandardService.Create:output_type -> response_model.Response
	7, // 9: standard_api.StandardService.CreateSeveral:output_type -> response_model.Response
	7, // 10: standard_api.StandardService.Update:output_type -> response_model.Response
	7, // 11: standard_api.StandardService.Delete:output_type -> response_model.Response
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pro_standard_api_proto_init() }
func file_pro_standard_api_proto_init() {
	if File_pro_standard_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_standard_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllStandards); i {
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
		file_pro_standard_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStandard); i {
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
		file_pro_standard_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSeveralStandard); i {
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
		file_pro_standard_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStandard); i {
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
		file_pro_standard_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStandard); i {
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
		file_pro_standard_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Standards); i {
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
			RawDescriptor: file_pro_standard_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pro_standard_api_proto_goTypes,
		DependencyIndexes: file_pro_standard_api_proto_depIdxs,
		MessageInfos:      file_pro_standard_api_proto_msgTypes,
	}.Build()
	File_pro_standard_api_proto = out.File
	file_pro_standard_api_proto_rawDesc = nil
	file_pro_standard_api_proto_goTypes = nil
	file_pro_standard_api_proto_depIdxs = nil
}