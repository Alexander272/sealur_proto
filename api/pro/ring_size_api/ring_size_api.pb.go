// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: pro/ring_size_api.proto

package ring_size_api

import (
	response_model "github.com/Alexander272/sealur_proto/api/pro/models/response_model"
	ring_size_model "github.com/Alexander272/sealur_proto/api/pro/models/ring_size_model"
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

type GetRingSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRingSize) Reset() {
	*x = GetRingSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_size_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRingSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRingSize) ProtoMessage() {}

func (x *GetRingSize) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_size_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRingSize.ProtoReflect.Descriptor instead.
func (*GetRingSize) Descriptor() ([]byte, []int) {
	return file_pro_ring_size_api_proto_rawDescGZIP(), []int{0}
}

type CreateRingSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Outer float64 `protobuf:"fixed64,1,opt,name=outer,proto3" json:"outer,omitempty"`
	Inner float64 `protobuf:"fixed64,2,opt,name=inner,proto3" json:"inner,omitempty"`
	Type  string  `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *CreateRingSize) Reset() {
	*x = CreateRingSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_size_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRingSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRingSize) ProtoMessage() {}

func (x *CreateRingSize) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_size_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRingSize.ProtoReflect.Descriptor instead.
func (*CreateRingSize) Descriptor() ([]byte, []int) {
	return file_pro_ring_size_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRingSize) GetOuter() float64 {
	if x != nil {
		return x.Outer
	}
	return 0
}

func (x *CreateRingSize) GetInner() float64 {
	if x != nil {
		return x.Inner
	}
	return 0
}

func (x *CreateRingSize) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type UpdateRingSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Outer float64 `protobuf:"fixed64,2,opt,name=outer,proto3" json:"outer,omitempty"`
	Inner float64 `protobuf:"fixed64,3,opt,name=inner,proto3" json:"inner,omitempty"`
	Type  string  `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *UpdateRingSize) Reset() {
	*x = UpdateRingSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_size_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRingSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRingSize) ProtoMessage() {}

func (x *UpdateRingSize) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_size_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRingSize.ProtoReflect.Descriptor instead.
func (*UpdateRingSize) Descriptor() ([]byte, []int) {
	return file_pro_ring_size_api_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRingSize) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRingSize) GetOuter() float64 {
	if x != nil {
		return x.Outer
	}
	return 0
}

func (x *UpdateRingSize) GetInner() float64 {
	if x != nil {
		return x.Inner
	}
	return 0
}

func (x *UpdateRingSize) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type DeleteRingSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRingSize) Reset() {
	*x = DeleteRingSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_size_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRingSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRingSize) ProtoMessage() {}

func (x *DeleteRingSize) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_size_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRingSize.ProtoReflect.Descriptor instead.
func (*DeleteRingSize) Descriptor() ([]byte, []int) {
	return file_pro_ring_size_api_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRingSize) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RingSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sizes []*ring_size_model.RingSize `protobuf:"bytes,1,rep,name=sizes,proto3" json:"sizes,omitempty"`
}

func (x *RingSize) Reset() {
	*x = RingSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_ring_size_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RingSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RingSize) ProtoMessage() {}

func (x *RingSize) ProtoReflect() protoreflect.Message {
	mi := &file_pro_ring_size_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RingSize.ProtoReflect.Descriptor instead.
func (*RingSize) Descriptor() ([]byte, []int) {
	return file_pro_ring_size_api_proto_rawDescGZIP(), []int{4}
}

func (x *RingSize) GetSizes() []*ring_size_model.RingSize {
	if x != nil {
		return x.Sizes
	}
	return nil
}

var File_pro_ring_size_api_proto protoreflect.FileDescriptor

var file_pro_ring_size_api_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x2f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x52, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x50, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x60, 0x0a, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x20,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x3b, 0x0a, 0x08, 0x52, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2f, 0x0a, 0x05,
	0x73, 0x69, 0x7a, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x69,
	0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x05, 0x73, 0x69, 0x7a, 0x65, 0x73, 0x32, 0x99, 0x02,
	0x0a, 0x0f, 0x52, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3d, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1a, 0x2e, 0x72, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x1a, 0x17, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x41, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x1a, 0x18, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x1d, 0x2e, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x1a,
	0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x2f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro_ring_size_api_proto_rawDescOnce sync.Once
	file_pro_ring_size_api_proto_rawDescData = file_pro_ring_size_api_proto_rawDesc
)

func file_pro_ring_size_api_proto_rawDescGZIP() []byte {
	file_pro_ring_size_api_proto_rawDescOnce.Do(func() {
		file_pro_ring_size_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_ring_size_api_proto_rawDescData)
	})
	return file_pro_ring_size_api_proto_rawDescData
}

var file_pro_ring_size_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pro_ring_size_api_proto_goTypes = []interface{}{
	(*GetRingSize)(nil),              // 0: ring_size_api.GetRingSize
	(*CreateRingSize)(nil),           // 1: ring_size_api.CreateRingSize
	(*UpdateRingSize)(nil),           // 2: ring_size_api.UpdateRingSize
	(*DeleteRingSize)(nil),           // 3: ring_size_api.DeleteRingSize
	(*RingSize)(nil),                 // 4: ring_size_api.RingSize
	(*ring_size_model.RingSize)(nil), // 5: ring_size_model.RingSize
	(*response_model.Response)(nil),  // 6: response_model.Response
}
var file_pro_ring_size_api_proto_depIdxs = []int32{
	5, // 0: ring_size_api.RingSize.sizes:type_name -> ring_size_model.RingSize
	0, // 1: ring_size_api.RingSizeService.GetAll:input_type -> ring_size_api.GetRingSize
	1, // 2: ring_size_api.RingSizeService.Create:input_type -> ring_size_api.CreateRingSize
	2, // 3: ring_size_api.RingSizeService.Update:input_type -> ring_size_api.UpdateRingSize
	3, // 4: ring_size_api.RingSizeService.Delete:input_type -> ring_size_api.DeleteRingSize
	4, // 5: ring_size_api.RingSizeService.GetAll:output_type -> ring_size_api.RingSize
	6, // 6: ring_size_api.RingSizeService.Create:output_type -> response_model.Response
	6, // 7: ring_size_api.RingSizeService.Update:output_type -> response_model.Response
	6, // 8: ring_size_api.RingSizeService.Delete:output_type -> response_model.Response
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pro_ring_size_api_proto_init() }
func file_pro_ring_size_api_proto_init() {
	if File_pro_ring_size_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_ring_size_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRingSize); i {
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
		file_pro_ring_size_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRingSize); i {
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
		file_pro_ring_size_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRingSize); i {
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
		file_pro_ring_size_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRingSize); i {
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
		file_pro_ring_size_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RingSize); i {
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
			RawDescriptor: file_pro_ring_size_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pro_ring_size_api_proto_goTypes,
		DependencyIndexes: file_pro_ring_size_api_proto_depIdxs,
		MessageInfos:      file_pro_ring_size_api_proto_msgTypes,
	}.Build()
	File_pro_ring_size_api_proto = out.File
	file_pro_ring_size_api_proto_rawDesc = nil
	file_pro_ring_size_api_proto_goTypes = nil
	file_pro_ring_size_api_proto_depIdxs = nil
}