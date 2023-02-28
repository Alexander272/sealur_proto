// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: pro/snp_api.proto

package snp_api

import (
	snp_data_model "github.com/Alexander272/sealur_proto/api/pro/models/snp_data_model"
	snp_model "github.com/Alexander272/sealur_proto/api/pro/models/snp_model"
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

type GetSnp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SnpTypeId        string `protobuf:"bytes,1,opt,name=snpTypeId,proto3" json:"snpTypeId,omitempty"`
	SnpStandardId    string `protobuf:"bytes,2,opt,name=snpStandardId,proto3" json:"snpStandardId,omitempty"`
	StandardId       string `protobuf:"bytes,3,opt,name=standardId,proto3" json:"standardId,omitempty"`
	FlangeStandardId string `protobuf:"bytes,4,opt,name=flangeStandardId,proto3" json:"flangeStandardId,omitempty"`
}

func (x *GetSnp) Reset() {
	*x = GetSnp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_snp_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSnp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSnp) ProtoMessage() {}

func (x *GetSnp) ProtoReflect() protoreflect.Message {
	mi := &file_pro_snp_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSnp.ProtoReflect.Descriptor instead.
func (*GetSnp) Descriptor() ([]byte, []int) {
	return file_pro_snp_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetSnp) GetSnpTypeId() string {
	if x != nil {
		return x.SnpTypeId
	}
	return ""
}

func (x *GetSnp) GetSnpStandardId() string {
	if x != nil {
		return x.SnpStandardId
	}
	return ""
}

func (x *GetSnp) GetStandardId() string {
	if x != nil {
		return x.StandardId
	}
	return ""
}

func (x *GetSnp) GetFlangeStandardId() string {
	if x != nil {
		return x.FlangeStandardId
	}
	return ""
}

type GetSnpData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StandardId string `protobuf:"bytes,1,opt,name=standardId,proto3" json:"standardId,omitempty"`
}

func (x *GetSnpData) Reset() {
	*x = GetSnpData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_snp_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSnpData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSnpData) ProtoMessage() {}

func (x *GetSnpData) ProtoReflect() protoreflect.Message {
	mi := &file_pro_snp_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSnpData.ProtoReflect.Descriptor instead.
func (*GetSnpData) Descriptor() ([]byte, []int) {
	return file_pro_snp_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetSnpData) GetStandardId() string {
	if x != nil {
		return x.StandardId
	}
	return ""
}

type Snp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// snp_data
	Snp *snp_data_model.SnpData `protobuf:"bytes,1,opt,name=snp,proto3" json:"snp,omitempty"` // size
}

func (x *Snp) Reset() {
	*x = Snp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_snp_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Snp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snp) ProtoMessage() {}

func (x *Snp) ProtoReflect() protoreflect.Message {
	mi := &file_pro_snp_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snp.ProtoReflect.Descriptor instead.
func (*Snp) Descriptor() ([]byte, []int) {
	return file_pro_snp_api_proto_rawDescGZIP(), []int{2}
}

func (x *Snp) GetSnp() *snp_data_model.SnpData {
	if x != nil {
		return x.Snp
	}
	return nil
}

type SnpData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// flange_type & snp_type
	// snp_material
	// * если стандарт не определен, то еще
	// snp_filler
	// mounting
	SnpData *snp_model.SnpData `protobuf:"bytes,1,opt,name=snpData,proto3" json:"snpData,omitempty"`
}

func (x *SnpData) Reset() {
	*x = SnpData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pro_snp_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnpData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnpData) ProtoMessage() {}

func (x *SnpData) ProtoReflect() protoreflect.Message {
	mi := &file_pro_snp_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnpData.ProtoReflect.Descriptor instead.
func (*SnpData) Descriptor() ([]byte, []int) {
	return file_pro_snp_api_proto_rawDescGZIP(), []int{3}
}

func (x *SnpData) GetSnpData() *snp_model.SnpData {
	if x != nil {
		return x.SnpData
	}
	return nil
}

var File_pro_snp_api_proto protoreflect.FileDescriptor

var file_pro_snp_api_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x2f, 0x73, 0x6e, 0x70, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x6e, 0x70, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x70, 0x72,
	0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x73, 0x6e, 0x70, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70,
	0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x73, 0x6e, 0x70, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x06, 0x47, 0x65,
	0x74, 0x53, 0x6e, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6e, 0x70, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6e, 0x70, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x6e, 0x70, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72,
	0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x6e, 0x70, 0x53, 0x74,
	0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x6e,
	0x64, 0x61, 0x72, 0x64, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x66, 0x6c, 0x61, 0x6e,
	0x67, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61,
	0x72, 0x64, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x6e, 0x70, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64,
	0x49, 0x64, 0x22, 0x30, 0x0a, 0x03, 0x53, 0x6e, 0x70, 0x12, 0x29, 0x0a, 0x03, 0x73, 0x6e, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x6e, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x03, 0x73, 0x6e, 0x70, 0x22, 0x37, 0x0a, 0x07, 0x53, 0x6e, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x2c, 0x0a, 0x07, 0x73, 0x6e, 0x70, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x6e, 0x70,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x73, 0x6e, 0x70, 0x44, 0x61, 0x74, 0x61, 0x32, 0x68, 0x0a,
	0x0e, 0x53, 0x6e, 0x70, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x24, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0f, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x6e, 0x70, 0x1a, 0x0c, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x53, 0x6e, 0x70, 0x12, 0x30, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x13, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6e,
	0x70, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x10, 0x2e, 0x73, 0x6e, 0x70, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x6e, 0x70, 0x44, 0x61, 0x74, 0x61, 0x42, 0x11, 0x5a, 0x0f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x2f, 0x73, 0x6e, 0x70, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pro_snp_api_proto_rawDescOnce sync.Once
	file_pro_snp_api_proto_rawDescData = file_pro_snp_api_proto_rawDesc
)

func file_pro_snp_api_proto_rawDescGZIP() []byte {
	file_pro_snp_api_proto_rawDescOnce.Do(func() {
		file_pro_snp_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pro_snp_api_proto_rawDescData)
	})
	return file_pro_snp_api_proto_rawDescData
}

var file_pro_snp_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pro_snp_api_proto_goTypes = []interface{}{
	(*GetSnp)(nil),                 // 0: snp_api.GetSnp
	(*GetSnpData)(nil),             // 1: snp_api.GetSnpData
	(*Snp)(nil),                    // 2: snp_api.Snp
	(*SnpData)(nil),                // 3: snp_api.SnpData
	(*snp_data_model.SnpData)(nil), // 4: snp_data_model.SnpData
	(*snp_model.SnpData)(nil),      // 5: snp_model.SnpData
}
var file_pro_snp_api_proto_depIdxs = []int32{
	4, // 0: snp_api.Snp.snp:type_name -> snp_data_model.SnpData
	5, // 1: snp_api.SnpData.snpData:type_name -> snp_model.SnpData
	0, // 2: snp_api.SnpDataService.Get:input_type -> snp_api.GetSnp
	1, // 3: snp_api.SnpDataService.GetData:input_type -> snp_api.GetSnpData
	2, // 4: snp_api.SnpDataService.Get:output_type -> snp_api.Snp
	3, // 5: snp_api.SnpDataService.GetData:output_type -> snp_api.SnpData
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pro_snp_api_proto_init() }
func file_pro_snp_api_proto_init() {
	if File_pro_snp_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pro_snp_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSnp); i {
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
		file_pro_snp_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSnpData); i {
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
		file_pro_snp_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Snp); i {
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
		file_pro_snp_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnpData); i {
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
			RawDescriptor: file_pro_snp_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pro_snp_api_proto_goTypes,
		DependencyIndexes: file_pro_snp_api_proto_depIdxs,
		MessageInfos:      file_pro_snp_api_proto_msgTypes,
	}.Build()
	File_pro_snp_api_proto = out.File
	file_pro_snp_api_proto_rawDesc = nil
	file_pro_snp_api_proto_goTypes = nil
	file_pro_snp_api_proto_depIdxs = nil
}
