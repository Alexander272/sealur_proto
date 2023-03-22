// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: pro/position_api.proto

package position_api

import (
	position_model "github.com/Alexander272/sealur_proto/api/pro/models/position_model"
	response_model "github.com/Alexander272/sealur_proto/api/pro/models/response_model"
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

type CreatePosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position *position_model.FullPosition `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *CreatePosition) Reset() {
	*x = CreatePosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_position_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePosition) ProtoMessage() {}

func (x *CreatePosition) ProtoReflect() protoreflect.Message {
	mi := &file_pro_position_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePosition.ProtoReflect.Descriptor instead.
func (*CreatePosition) Descriptor() ([]byte, []int) {
	return file_pro_position_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePosition) GetPosition() *position_model.FullPosition {
	if x != nil {
		return x.Position
	}
	return nil
}

type CopyPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount      string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	OrderId     string `protobuf:"bytes,3,opt,name=orderId,proto3" json:"orderId,omitempty"`
	FromOrderId string `protobuf:"bytes,4,opt,name=fromOrderId,proto3" json:"fromOrderId,omitempty"`
}

func (x *CopyPosition) Reset() {
	*x = CopyPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_position_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CopyPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyPosition) ProtoMessage() {}

func (x *CopyPosition) ProtoReflect() protoreflect.Message {
	mi := &file_pro_position_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyPosition.ProtoReflect.Descriptor instead.
func (*CopyPosition) Descriptor() ([]byte, []int) {
	return file_pro_position_api_proto_rawDescGZIP(), []int{1}
}

func (x *CopyPosition) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CopyPosition) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *CopyPosition) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *CopyPosition) GetFromOrderId() string {
	if x != nil {
		return x.FromOrderId
	}
	return ""
}

type UpdatePosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position *position_model.FullPosition `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *UpdatePosition) Reset() {
	*x = UpdatePosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_position_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePosition) ProtoMessage() {}

func (x *UpdatePosition) ProtoReflect() protoreflect.Message {
	mi := &file_pro_position_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePosition.ProtoReflect.Descriptor instead.
func (*UpdatePosition) Descriptor() ([]byte, []int) {
	return file_pro_position_api_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePosition) GetPosition() *position_model.FullPosition {
	if x != nil {
		return x.Position
	}
	return nil
}

type DeletePosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePosition) Reset() {
	*x = DeletePosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_position_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePosition) ProtoMessage() {}

func (x *DeletePosition) ProtoReflect() protoreflect.Message {
	mi := &file_pro_position_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePosition.ProtoReflect.Descriptor instead.
func (*DeletePosition) Descriptor() ([]byte, []int) {
	return file_pro_position_api_proto_rawDescGZIP(), []int{3}
}

func (x *DeletePosition) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CopyDrawing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Drawing string `protobuf:"bytes,1,opt,name=drawing,proto3" json:"drawing,omitempty"`
}

func (x *CopyDrawing) Reset() {
	*x = CopyDrawing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_position_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CopyDrawing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyDrawing) ProtoMessage() {}

func (x *CopyDrawing) ProtoReflect() protoreflect.Message {
	mi := &file_pro_position_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyDrawing.ProtoReflect.Descriptor instead.
func (*CopyDrawing) Descriptor() ([]byte, []int) {
	return file_pro_position_api_proto_rawDescGZIP(), []int{4}
}

func (x *CopyDrawing) GetDrawing() string {
	if x != nil {
		return x.Drawing
	}
	return ""
}

var File_pro_position_api_proto protoreflect.FileDescriptor

var file_pro_position_api_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x46, 0x75,
	0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x72, 0x0a, 0x0c, 0x43, 0x6f, 0x70, 0x79, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72, 0x6f,
	0x6d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x46, 0x75,
	0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x0b, 0x43, 0x6f, 0x70, 0x79, 0x44, 0x72,
	0x61, 0x77, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x72, 0x61, 0x77, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x72, 0x61, 0x77, 0x69, 0x6e, 0x67, 0x32,
	0x98, 0x02, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x1c, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x43, 0x6f, 0x70,
	0x79, 0x12, 0x1a, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x6f, 0x70, 0x79, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x19, 0x2e,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x70,
	0x79, 0x44, 0x72, 0x61, 0x77, 0x69, 0x6e, 0x67, 0x12, 0x40, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x1c, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pro_position_api_proto_rawDescOnce sync.Once
	file_pro_position_api_proto_rawDescData = file_pro_position_api_proto_rawDesc
)

func file_pro_position_api_proto_rawDescGZIP() []byte {
	file_pro_position_api_proto_rawDescOnce.Do(func() {
		file_pro_position_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_position_api_proto_rawDescData)
	})
	return file_pro_position_api_proto_rawDescData
}

var file_pro_position_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pro_position_api_proto_goTypes = []interface{}{
	(*CreatePosition)(nil),              // 0: position_api.CreatePosition
	(*CopyPosition)(nil),                // 1: position_api.CopyPosition
	(*UpdatePosition)(nil),              // 2: position_api.UpdatePosition
	(*DeletePosition)(nil),              // 3: position_api.DeletePosition
	(*CopyDrawing)(nil),                 // 4: position_api.CopyDrawing
	(*position_model.FullPosition)(nil), // 5: position_model.FullPosition
	(*response_model.IdResponse)(nil),   // 6: response_model.IdResponse
	(*response_model.Response)(nil),     // 7: response_model.Response
}
var file_pro_position_api_proto_depIdxs = []int32{
	5, // 0: position_api.CreatePosition.position:type_name -> position_model.FullPosition
	5, // 1: position_api.UpdatePosition.position:type_name -> position_model.FullPosition
	0, // 2: position_api.PositionService.Create:input_type -> position_api.CreatePosition
	2, // 3: position_api.PositionService.Update:input_type -> position_api.UpdatePosition
	1, // 4: position_api.PositionService.Copy:input_type -> position_api.CopyPosition
	3, // 5: position_api.PositionService.Delete:input_type -> position_api.DeletePosition
	6, // 6: position_api.PositionService.Create:output_type -> response_model.IdResponse
	7, // 7: position_api.PositionService.Update:output_type -> response_model.Response
	4, // 8: position_api.PositionService.Copy:output_type -> position_api.CopyDrawing
	7, // 9: position_api.PositionService.Delete:output_type -> response_model.Response
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pro_position_api_proto_init() }
func file_pro_position_api_proto_init() {
	if File_pro_position_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_position_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePosition); i {
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
		file_pro_position_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CopyPosition); i {
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
		file_pro_position_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePosition); i {
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
		file_pro_position_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePosition); i {
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
		file_pro_position_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CopyDrawing); i {
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
			RawDescriptor: file_pro_position_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pro_position_api_proto_goTypes,
		DependencyIndexes: file_pro_position_api_proto_depIdxs,
		MessageInfos:      file_pro_position_api_proto_msgTypes,
	}.Build()
	File_pro_position_api_proto = out.File
	file_pro_position_api_proto_rawDesc = nil
	file_pro_position_api_proto_goTypes = nil
	file_pro_position_api_proto_depIdxs = nil
}
