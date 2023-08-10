// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: pro/ring_modifying_api.proto

package ring_modifying_api

import (
	response_model "github.com/Alexander272/sealur_proto/api/pro/models/response_model"
	ring_modifying_model "github.com/Alexander272/sealur_proto/api/pro/models/ring_modifying_model"
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

type GetRingModifying struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRingModifying) Reset() {
	*x = GetRingModifying{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_modifying_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRingModifying) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRingModifying) ProtoMessage() {}

func (x *GetRingModifying) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_modifying_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRingModifying.ProtoReflect.Descriptor instead.
func (*GetRingModifying) Descriptor() ([]byte, []int) {
	return file_pro_ring_modifying_api_proto_rawDescGZIP(), []int{0}
}

type CreateRingModifying struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateRingModifying) Reset() {
	*x = CreateRingModifying{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_modifying_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRingModifying) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRingModifying) ProtoMessage() {}

func (x *CreateRingModifying) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_modifying_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRingModifying.ProtoReflect.Descriptor instead.
func (*CreateRingModifying) Descriptor() ([]byte, []int) {
	return file_pro_ring_modifying_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRingModifying) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateRingModifying) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type UpdateRingModifying struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code        string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *UpdateRingModifying) Reset() {
	*x = UpdateRingModifying{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_modifying_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRingModifying) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRingModifying) ProtoMessage() {}

func (x *UpdateRingModifying) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_modifying_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRingModifying.ProtoReflect.Descriptor instead.
func (*UpdateRingModifying) Descriptor() ([]byte, []int) {
	return file_pro_ring_modifying_api_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRingModifying) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRingModifying) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateRingModifying) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type DeleteRingModifying struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRingModifying) Reset() {
	*x = DeleteRingModifying{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_modifying_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRingModifying) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRingModifying) ProtoMessage() {}

func (x *DeleteRingModifying) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_modifying_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRingModifying.ProtoReflect.Descriptor instead.
func (*DeleteRingModifying) Descriptor() ([]byte, []int) {
	return file_pro_ring_modifying_api_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRingModifying) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RingModifying struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Modifying []*ring_modifying_model.RingModifying `protobuf:"bytes,1,rep,name=modifying,proto3" json:"modifying,omitempty"`
}

func (x *RingModifying) Reset() {
	*x = RingModifying{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_modifying_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RingModifying) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RingModifying) ProtoMessage() {}

func (x *RingModifying) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_modifying_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RingModifying.ProtoReflect.Descriptor instead.
func (*RingModifying) Descriptor() ([]byte, []int) {
	return file_pro_ring_modifying_api_proto_rawDescGZIP(), []int{4}
}

func (x *RingModifying) GetModifying() []*ring_modifying_model.RingModifying {
	if x != nil {
		return x.Modifying
	}
	return nil
}

var File_pro_ring_modifying_api_proto protoreflect.FileDescriptor

var file_pro_ring_modifying_api_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x2f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x61,
	0x70, 0x69, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x52, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x69, 0x6e, 0x67, 0x22, 0x4b,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5b, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x69,
	0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x69, 0x6e, 0x67, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x52, 0x0a, 0x0d, 0x52, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x69, 0x6e, 0x67,
	0x12, 0x41, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x69, 0x6e, 0x67, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x69, 0x6e, 0x67, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x69, 0x6e, 0x67, 0x32, 0xd0, 0x02, 0x0a, 0x14, 0x52, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x06,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x24, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x69, 0x6e, 0x67, 0x1a, 0x21, 0x2e, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x69, 0x6e, 0x67, 0x12,
	0x4b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x69,
	0x6e, 0x67, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x69, 0x6e, 0x67, 0x1a,
	0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x69, 0x6e, 0x67, 0x1a, 0x18, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x2f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x69, 0x6e, 0x67,
	0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro_ring_modifying_api_proto_rawDescOnce sync.Once
	file_pro_ring_modifying_api_proto_rawDescData = file_pro_ring_modifying_api_proto_rawDesc
)

func file_pro_ring_modifying_api_proto_rawDescGZIP() []byte {
	file_pro_ring_modifying_api_proto_rawDescOnce.Do(func() {
		file_pro_ring_modifying_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_ring_modifying_api_proto_rawDescData)
	})
	return file_pro_ring_modifying_api_proto_rawDescData
}

var file_pro_ring_modifying_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pro_ring_modifying_api_proto_goTypes = []interface{}{
	(*GetRingModifying)(nil),                   // 0: ring_modifying_api.GetRingModifying
	(*CreateRingModifying)(nil),                // 1: ring_modifying_api.CreateRingModifying
	(*UpdateRingModifying)(nil),                // 2: ring_modifying_api.UpdateRingModifying
	(*DeleteRingModifying)(nil),                // 3: ring_modifying_api.DeleteRingModifying
	(*RingModifying)(nil),                      // 4: ring_modifying_api.RingModifying
	(*ring_modifying_model.RingModifying)(nil), // 5: ring_modifying_model.RingModifying
	(*response_model.Response)(nil),            // 6: response_model.Response
}
var file_pro_ring_modifying_api_proto_depIdxs = []int32{
	5, // 0: ring_modifying_api.RingModifying.modifying:type_name -> ring_modifying_model.RingModifying
	0, // 1: ring_modifying_api.RingModifyingService.GetAll:input_type -> ring_modifying_api.GetRingModifying
	1, // 2: ring_modifying_api.RingModifyingService.Create:input_type -> ring_modifying_api.CreateRingModifying
	2, // 3: ring_modifying_api.RingModifyingService.Update:input_type -> ring_modifying_api.UpdateRingModifying
	3, // 4: ring_modifying_api.RingModifyingService.Delete:input_type -> ring_modifying_api.DeleteRingModifying
	4, // 5: ring_modifying_api.RingModifyingService.GetAll:output_type -> ring_modifying_api.RingModifying
	6, // 6: ring_modifying_api.RingModifyingService.Create:output_type -> response_model.Response
	6, // 7: ring_modifying_api.RingModifyingService.Update:output_type -> response_model.Response
	6, // 8: ring_modifying_api.RingModifyingService.Delete:output_type -> response_model.Response
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pro_ring_modifying_api_proto_init() }
func file_pro_ring_modifying_api_proto_init() {
	if File_pro_ring_modifying_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_ring_modifying_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRingModifying); i {
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
		file_pro_ring_modifying_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRingModifying); i {
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
		file_pro_ring_modifying_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRingModifying); i {
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
		file_pro_ring_modifying_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRingModifying); i {
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
		file_pro_ring_modifying_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RingModifying); i {
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
			RawDescriptor: file_pro_ring_modifying_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pro_ring_modifying_api_proto_goTypes,
		DependencyIndexes: file_pro_ring_modifying_api_proto_depIdxs,
		MessageInfos:      file_pro_ring_modifying_api_proto_msgTypes,
	}.Build()
	File_pro_ring_modifying_api_proto = out.File
	file_pro_ring_modifying_api_proto_rawDesc = nil
	file_pro_ring_modifying_api_proto_goTypes = nil
	file_pro_ring_modifying_api_proto_depIdxs = nil
}