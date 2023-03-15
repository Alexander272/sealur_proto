// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: user/user_api.proto

package user_api

import (
	response_model "github.com/Alexander272/sealur_proto/api/pro/models/response_model"
	user_model "github.com/Alexander272/sealur_proto/api/user/models/user_model"
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

type GetAllUser_SearchField int32

const (
	GetAllUser_organization GetAllUser_SearchField = 0
	GetAllUser_city         GetAllUser_SearchField = 1
	GetAllUser_login        GetAllUser_SearchField = 2
	GetAllUser_name         GetAllUser_SearchField = 3
	GetAllUser_email        GetAllUser_SearchField = 4
)

// Enum value maps for GetAllUser_SearchField.
var (
	GetAllUser_SearchField_name = map[int32]string{
		0: "organization",
		1: "city",
		2: "login",
		3: "name",
		4: "email",
	}
	GetAllUser_SearchField_value = map[string]int32{
		"organization": 0,
		"city":         1,
		"login":        2,
		"name":         3,
		"email":        4,
	}
)

func (x GetAllUser_SearchField) Enum() *GetAllUser_SearchField {
	p := new(GetAllUser_SearchField)
	*p = x
	return p
}

func (x GetAllUser_SearchField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetAllUser_SearchField) Descriptor() protoreflect.EnumDescriptor {
	return file_user_user_api_proto_enumTypes[0].Descriptor()
}

func (GetAllUser_SearchField) Type() protoreflect.EnumType {
	return &file_user_user_api_proto_enumTypes[0]
}

func (x GetAllUser_SearchField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetAllUser_SearchField.Descriptor instead.
func (GetAllUser_SearchField) EnumDescriptor() ([]byte, []int) {
	return file_user_user_api_proto_rawDescGZIP(), []int{1, 0}
}

type GetUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Id       string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *GetUser) Reset() {
	*x = GetUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUser) ProtoMessage() {}

func (x *GetUser) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUser.ProtoReflect.Descriptor instead.
func (*GetUser) Descriptor() ([]byte, []int) {
	return file_user_user_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetUser) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *GetUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetAllUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32              `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32              `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Search *GetAllUser_Search `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetAllUser) Reset() {
	*x = GetAllUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUser) ProtoMessage() {}

func (x *GetAllUser) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUser.ProtoReflect.Descriptor instead.
func (*GetAllUser) Descriptor() ([]byte, []int) {
	return file_user_user_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllUser) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllUser) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAllUser) GetSearch() *GetAllUser_Search {
	if x != nil {
		return x.Search
	}
	return nil
}

type GetNewUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetNewUser) Reset() {
	*x = GetNewUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNewUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewUser) ProtoMessage() {}

func (x *GetNewUser) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewUser.ProtoReflect.Descriptor instead.
func (*GetNewUser) Descriptor() ([]byte, []int) {
	return file_user_user_api_proto_rawDescGZIP(), []int{2}
}

type CreateUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Company  string `protobuf:"bytes,1,opt,name=company,proto3" json:"company,omitempty"`
	Address  string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Inn      string `protobuf:"bytes,3,opt,name=inn,proto3" json:"inn,omitempty"`
	Kpp      string `protobuf:"bytes,4,opt,name=kpp,proto3" json:"kpp,omitempty"`
	Region   string `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
	City     string `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
	Name     string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Position string `protobuf:"bytes,8,opt,name=position,proto3" json:"position,omitempty"`
	Email    string `protobuf:"bytes,9,opt,name=email,proto3" json:"email,omitempty"`
	Phone    string `protobuf:"bytes,10,opt,name=phone,proto3" json:"phone,omitempty"`
	Password string `protobuf:"bytes,11,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateUser) Reset() {
	*x = CreateUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUser) ProtoMessage() {}

func (x *CreateUser) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUser.ProtoReflect.Descriptor instead.
func (*CreateUser) Descriptor() ([]byte, []int) {
	return file_user_user_api_proto_rawDescGZIP(), []int{3}
}

func (x *CreateUser) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *CreateUser) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateUser) GetInn() string {
	if x != nil {
		return x.Inn
	}
	return ""
}

func (x *CreateUser) GetKpp() string {
	if x != nil {
		return x.Kpp
	}
	return ""
}

func (x *CreateUser) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *CreateUser) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CreateUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUser) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *CreateUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUser) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UpdateUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Company  string `protobuf:"bytes,2,opt,name=company,proto3" json:"company,omitempty"`
	Address  string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Inn      string `protobuf:"bytes,4,opt,name=inn,proto3" json:"inn,omitempty"`
	Kpp      string `protobuf:"bytes,5,opt,name=kpp,proto3" json:"kpp,omitempty"`
	Region   string `protobuf:"bytes,6,opt,name=region,proto3" json:"region,omitempty"`
	City     string `protobuf:"bytes,7,opt,name=city,proto3" json:"city,omitempty"`
	Name     string `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Position string `protobuf:"bytes,9,opt,name=position,proto3" json:"position,omitempty"`
	Email    string `protobuf:"bytes,10,opt,name=email,proto3" json:"email,omitempty"`
	Phone    string `protobuf:"bytes,11,opt,name=phone,proto3" json:"phone,omitempty"`
	Password string `protobuf:"bytes,12,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UpdateUser) Reset() {
	*x = UpdateUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUser) ProtoMessage() {}

func (x *UpdateUser) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUser.ProtoReflect.Descriptor instead.
func (*UpdateUser) Descriptor() ([]byte, []int) {
	return file_user_user_api_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateUser) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *UpdateUser) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UpdateUser) GetInn() string {
	if x != nil {
		return x.Inn
	}
	return ""
}

func (x *UpdateUser) GetKpp() string {
	if x != nil {
		return x.Kpp
	}
	return ""
}

func (x *UpdateUser) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *UpdateUser) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UpdateUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateUser) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *UpdateUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateUser) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type DeleteUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUser) Reset() {
	*x = DeleteUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUser) ProtoMessage() {}

func (x *DeleteUser) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUser.ProtoReflect.Descriptor instead.
func (*DeleteUser) Descriptor() ([]byte, []int) {
	return file_user_user_api_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ConfirmUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ConfirmUser) Reset() {
	*x = ConfirmUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmUser) ProtoMessage() {}

func (x *ConfirmUser) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmUser.ProtoReflect.Descriptor instead.
func (*ConfirmUser) Descriptor() ([]byte, []int) {
	return file_user_user_api_proto_rawDescGZIP(), []int{6}
}

func (x *ConfirmUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *user_model.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_user_api_proto_rawDescGZIP(), []int{7}
}

func (x *User) GetUser() *user_model.User {
	if x != nil {
		return x.User
	}
	return nil
}

type Users struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*user_model.User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Count int32              `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Users) Reset() {
	*x = Users{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Users) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Users) ProtoMessage() {}

func (x *Users) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Users.ProtoReflect.Descriptor instead.
func (*Users) Descriptor() ([]byte, []int) {
	return file_user_user_api_proto_rawDescGZIP(), []int{8}
}

func (x *Users) GetUsers() []*user_model.User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *Users) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetAllUser_Search struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field []GetAllUser_SearchField `protobuf:"varint,1,rep,packed,name=field,proto3,enum=user_api.GetAllUser_SearchField" json:"field,omitempty"`
	Value string                   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetAllUser_Search) Reset() {
	*x = GetAllUser_Search{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUser_Search) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUser_Search) ProtoMessage() {}

func (x *GetAllUser_Search) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUser_Search.ProtoReflect.Descriptor instead.
func (*GetAllUser_Search) Descriptor() ([]byte, []int) {
	return file_user_user_api_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetAllUser_Search) GetField() []GetAllUser_SearchField {
	if x != nil {
		return x.Field
	}
	return nil
}

func (x *GetAllUser_Search) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_user_user_api_proto protoreflect.FileDescriptor

var file_user_user_api_proto_rawDesc = []byte{
	0x0a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x1a,
	0x1f, 0x70, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x92, 0x02, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x1a, 0x56, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x36, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x49, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x10, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x10, 0x01,
	0x12, 0x09, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x10, 0x04,
	0x22, 0x0c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x22, 0x88,
	0x02, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6e, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x69, 0x6e, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x70, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x70, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x98, 0x02, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x6e, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6e, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x70, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x70, 0x70,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x1c, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x1d, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x2c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x45, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x26, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x83, 0x03, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x11, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x2f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72,
	0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x2f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x12, 0x14, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65,
	0x72, 0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x38, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x18,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x38, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6c, 0x65, 0x78, 0x61,
	0x6e, 0x64, 0x65, 0x72, 0x32, 0x37, 0x32, 0x2f, 0x73, 0x65, 0x61, 0x6c, 0x75, 0x72, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_user_api_proto_rawDescOnce sync.Once
	file_user_user_api_proto_rawDescData = file_user_user_api_proto_rawDesc
)

func file_user_user_api_proto_rawDescGZIP() []byte {
	file_user_user_api_proto_rawDescOnce.Do(func() {
		file_user_user_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_user_api_proto_rawDescData)
	})
	return file_user_user_api_proto_rawDescData
}

var file_user_user_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_user_api_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_user_user_api_proto_goTypes = []interface{}{
	(GetAllUser_SearchField)(0),     // 0: user_api.GetAllUser.SearchField
	(*GetUser)(nil),                 // 1: user_api.GetUser
	(*GetAllUser)(nil),              // 2: user_api.GetAllUser
	(*GetNewUser)(nil),              // 3: user_api.GetNewUser
	(*CreateUser)(nil),              // 4: user_api.CreateUser
	(*UpdateUser)(nil),              // 5: user_api.UpdateUser
	(*DeleteUser)(nil),              // 6: user_api.DeleteUser
	(*ConfirmUser)(nil),             // 7: user_api.ConfirmUser
	(*User)(nil),                    // 8: user_api.User
	(*Users)(nil),                   // 9: user_api.Users
	(*GetAllUser_Search)(nil),       // 10: user_api.GetAllUser.Search
	(*user_model.User)(nil),         // 11: user_model.User
	(*response_model.Response)(nil), // 12: response_model.Response
}
var file_user_user_api_proto_depIdxs = []int32{
	10, // 0: user_api.GetAllUser.search:type_name -> user_api.GetAllUser.Search
	11, // 1: user_api.User.user:type_name -> user_model.User
	11, // 2: user_api.Users.users:type_name -> user_model.User
	0,  // 3: user_api.GetAllUser.Search.field:type_name -> user_api.GetAllUser.SearchField
	1,  // 4: user_api.UserService.Get:input_type -> user_api.GetUser
	2,  // 5: user_api.UserService.GetAll:input_type -> user_api.GetAllUser
	3,  // 6: user_api.UserService.GetNew:input_type -> user_api.GetNewUser
	4,  // 7: user_api.UserService.Create:input_type -> user_api.CreateUser
	7,  // 8: user_api.UserService.Confirm:input_type -> user_api.ConfirmUser
	5,  // 9: user_api.UserService.Update:input_type -> user_api.UpdateUser
	6,  // 10: user_api.UserService.Delete:input_type -> user_api.DeleteUser
	8,  // 11: user_api.UserService.Get:output_type -> user_api.User
	9,  // 12: user_api.UserService.GetAll:output_type -> user_api.Users
	9,  // 13: user_api.UserService.GetNew:output_type -> user_api.Users
	12, // 14: user_api.UserService.Create:output_type -> response_model.Response
	12, // 15: user_api.UserService.Confirm:output_type -> response_model.Response
	12, // 16: user_api.UserService.Update:output_type -> response_model.Response
	12, // 17: user_api.UserService.Delete:output_type -> response_model.Response
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_user_user_api_proto_init() }
func file_user_user_api_proto_init() {
	if File_user_user_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_user_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUser); i {
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
		file_user_user_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUser); i {
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
		file_user_user_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNewUser); i {
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
		file_user_user_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUser); i {
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
		file_user_user_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUser); i {
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
		file_user_user_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUser); i {
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
		file_user_user_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmUser); i {
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
		file_user_user_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_user_user_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Users); i {
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
		file_user_user_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUser_Search); i {
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
			RawDescriptor: file_user_user_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_user_api_proto_goTypes,
		DependencyIndexes: file_user_user_api_proto_depIdxs,
		EnumInfos:         file_user_user_api_proto_enumTypes,
		MessageInfos:      file_user_user_api_proto_msgTypes,
	}.Build()
	File_user_user_api_proto = out.File
	file_user_user_api_proto_rawDesc = nil
	file_user_user_api_proto_goTypes = nil
	file_user_user_api_proto_depIdxs = nil
}