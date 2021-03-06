// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/showcase/v1beta1/identity.proto

package genproto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A user.
type User struct {
	// The resource name of the user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The display_name of the user.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The email address of the user.
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	// The timestamp at which the user was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The latest timestamp at which the user was updated.
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_25043513edbd8d39, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *User) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

// The request message for the google.showcase.v1beta1.Identity\CreateUser
// method.
type CreateUserRequest struct {
	// The user to create.
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_25043513edbd8d39, []int{1}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// The request message for the google.showcase.v1beta1.Identity\GetUser
// method.
type GetUserRequest struct {
	// The resource name of the requested user.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_25043513edbd8d39, []int{2}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message for the google.showcase.v1beta1.Identity\UpdateUser
// method.
type UpdateUserRequest struct {
	// The user to update.
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// The field mask to determine wich fields are to be updated. If empty, the
	// server will assume all fields are to be updated.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_25043513edbd8d39, []int{3}
}

func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(m, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UpdateUserRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// The request message for the google.showcase.v1beta1.Identity\DeleteUser
// method.
type DeleteUserRequest struct {
	// The resource name of the user to delete.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_25043513edbd8d39, []int{4}
}

func (m *DeleteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRequest.Unmarshal(m, b)
}
func (m *DeleteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRequest.Marshal(b, m, deterministic)
}
func (m *DeleteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRequest.Merge(m, src)
}
func (m *DeleteUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRequest.Size(m)
}
func (m *DeleteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRequest proto.InternalMessageInfo

func (m *DeleteUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message for the google.showcase.v1beta1.Identity\ListUsers
// method.
type ListUsersRequest struct {
	// The maximum number of users to return. Server may return fewer users
	// than requested. If unspecified, server will pick an appropriate default.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The value of google.showcase.v1beta1.ListUsersResponse.next_page_token
	// returned from the previous call to
	// `google.showcase.v1beta1.Identity\ListUsers` method.
	PageToken            string   `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_25043513edbd8d39, []int{5}
}

func (m *ListUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersRequest.Unmarshal(m, b)
}
func (m *ListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersRequest.Marshal(b, m, deterministic)
}
func (m *ListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersRequest.Merge(m, src)
}
func (m *ListUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersRequest.Size(m)
}
func (m *ListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersRequest proto.InternalMessageInfo

func (m *ListUsersRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListUsersRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The response message for the google.showcase.v1beta1.Identity\ListUsers
// method.
type ListUsersResponse struct {
	// The list of users.
	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	// A token to retrieve next page of results.
	// Pass this value in ListUsersRequest.page_token field in the subsequent
	// call to `google.showcase.v1beta1.Message\ListUsers` method to retrieve the
	// next page of results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersResponse) Reset()         { *m = ListUsersResponse{} }
func (m *ListUsersResponse) String() string { return proto.CompactTextString(m) }
func (*ListUsersResponse) ProtoMessage()    {}
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_25043513edbd8d39, []int{6}
}

func (m *ListUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersResponse.Unmarshal(m, b)
}
func (m *ListUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersResponse.Marshal(b, m, deterministic)
}
func (m *ListUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersResponse.Merge(m, src)
}
func (m *ListUsersResponse) XXX_Size() int {
	return xxx_messageInfo_ListUsersResponse.Size(m)
}
func (m *ListUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersResponse proto.InternalMessageInfo

func (m *ListUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *ListUsersResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "google.showcase.v1beta1.User")
	proto.RegisterType((*CreateUserRequest)(nil), "google.showcase.v1beta1.CreateUserRequest")
	proto.RegisterType((*GetUserRequest)(nil), "google.showcase.v1beta1.GetUserRequest")
	proto.RegisterType((*UpdateUserRequest)(nil), "google.showcase.v1beta1.UpdateUserRequest")
	proto.RegisterType((*DeleteUserRequest)(nil), "google.showcase.v1beta1.DeleteUserRequest")
	proto.RegisterType((*ListUsersRequest)(nil), "google.showcase.v1beta1.ListUsersRequest")
	proto.RegisterType((*ListUsersResponse)(nil), "google.showcase.v1beta1.ListUsersResponse")
}

func init() {
	proto.RegisterFile("google/showcase/v1beta1/identity.proto", fileDescriptor_25043513edbd8d39)
}

var fileDescriptor_25043513edbd8d39 = []byte{
	// 751 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x5d, 0x4f, 0x13, 0x4d,
	0x14, 0xce, 0xf6, 0xe3, 0x7d, 0x61, 0xfa, 0xbe, 0x40, 0xe7, 0x4d, 0xa0, 0x2c, 0xe5, 0xb5, 0x6e,
	0x4c, 0xc1, 0x46, 0x77, 0x43, 0x21, 0x8a, 0x18, 0x13, 0x8b, 0x8a, 0xd1, 0x08, 0x21, 0x15, 0x6e,
	0xbc, 0x69, 0xa6, 0xdb, 0x43, 0x3b, 0x61, 0xbf, 0xdc, 0x99, 0x22, 0x1f, 0x21, 0x31, 0xfa, 0x0b,
	0x8c, 0xf1, 0x27, 0xf8, 0x67, 0xbc, 0xf5, 0x8e, 0x2b, 0x2e, 0xbc, 0xf2, 0x27, 0x78, 0x65, 0x66,
	0x66, 0xbb, 0x2d, 0xad, 0x2d, 0x44, 0xaf, 0x76, 0x7b, 0xce, 0x33, 0xe7, 0x79, 0xce, 0x99, 0xe7,
	0x6c, 0x51, 0xb1, 0xe9, 0xfb, 0x4d, 0x07, 0x2c, 0xd6, 0xf2, 0xdf, 0xd8, 0x84, 0x81, 0x75, 0xb0,
	0x54, 0x07, 0x4e, 0x96, 0x2c, 0xda, 0x00, 0x8f, 0x53, 0x7e, 0x64, 0x06, 0xa1, 0xcf, 0x7d, 0x3c,
	0xa3, 0x70, 0x66, 0x07, 0x67, 0x46, 0x38, 0x3d, 0x1f, 0x15, 0x20, 0x01, 0xb5, 0x88, 0xe7, 0xf9,
	0x9c, 0x70, 0xea, 0x7b, 0x4c, 0x1d, 0xd3, 0x67, 0x7a, 0xb2, 0xb6, 0x43, 0xc1, 0xe3, 0x51, 0xe2,
	0x5a, 0x4f, 0x62, 0x8f, 0x82, 0xd3, 0xa8, 0xd5, 0xa1, 0x45, 0x0e, 0xa8, 0x1f, 0x46, 0x80, 0xd9,
	0x1e, 0x40, 0x08, 0xcc, 0x6f, 0x87, 0x36, 0x44, 0xa9, 0xb9, 0x28, 0x25, 0x7f, 0xd5, 0xdb, 0x7b,
	0x16, 0xb8, 0x41, 0x47, 0xa8, 0x5e, 0xe8, 0x4f, 0xaa, 0xea, 0x2e, 0x61, 0xfb, 0x7d, 0xd4, 0x31,
	0x82, 0x53, 0x17, 0x18, 0x27, 0x6e, 0xa0, 0x00, 0xc6, 0xa7, 0x04, 0x4a, 0xed, 0x32, 0x08, 0x31,
	0x46, 0x29, 0x8f, 0xb8, 0x90, 0xd3, 0x0a, 0xda, 0xe2, 0x78, 0x55, 0xbe, 0xe3, 0x22, 0xfa, 0xa7,
	0x41, 0x59, 0xe0, 0x90, 0xa3, 0x9a, 0xcc, 0x25, 0x44, 0x6e, 0x3d, 0x79, 0x5e, 0x49, 0x54, 0x33,
	0x51, 0x62, 0x4b, 0xe0, 0x66, 0x51, 0x1a, 0x5c, 0x42, 0x9d, 0x5c, 0xb2, 0x0b, 0x50, 0x11, 0xfc,
	0x10, 0x65, 0xec, 0x10, 0x08, 0x87, 0x9a, 0x60, 0xce, 0xa5, 0x0a, 0xda, 0x62, 0xa6, 0xac, 0x9b,
	0xd1, 0x84, 0x3b, 0xb2, 0xcc, 0x9d, 0x8e, 0x2c, 0x71, 0x38, 0x59, 0x45, 0xea, 0x8c, 0x88, 0x8a,
	0x0a, 0xed, 0xa0, 0x11, 0x57, 0x48, 0x5f, 0xb1, 0x82, 0x3a, 0x23, 0xa2, 0x6b, 0xe5, 0xef, 0x15,
	0x0b, 0xe5, 0xe3, 0xdb, 0x54, 0x27, 0x49, 0x40, 0x99, 0x69, 0xfb, 0xae, 0x25, 0xbb, 0x9f, 0x6c,
	0x33, 0x08, 0x99, 0x75, 0x22, 0x1e, 0x35, 0xda, 0x38, 0x35, 0x36, 0x50, 0xf6, 0x91, 0xd4, 0x20,
	0xd2, 0x55, 0x78, 0xdd, 0x06, 0xc6, 0xf1, 0x12, 0x4a, 0x09, 0x80, 0x9c, 0x51, 0xa6, 0x3c, 0x6f,
	0x0e, 0xf1, 0x89, 0x29, 0xcf, 0x48, 0xa8, 0xf1, 0x1c, 0x4d, 0x3c, 0x05, 0xde, 0x5b, 0x64, 0xb5,
	0x77, 0xd0, 0xeb, 0x37, 0xce, 0x2b, 0x89, 0x1f, 0x95, 0xff, 0x47, 0xcb, 0x53, 0xd7, 0x61, 0xbc,
	0xd7, 0x50, 0x76, 0x57, 0xb6, 0xf5, 0x67, 0xa2, 0xf0, 0xfd, 0x78, 0xa4, 0xc2, 0x2a, 0xf2, 0x5a,
	0x7f, 0x35, 0xd2, 0x0d, 0xe1, 0xa6, 0x4d, 0xc2, 0xf6, 0x3b, 0xd3, 0x14, 0xef, 0xc6, 0x26, 0xca,
	0x3e, 0x06, 0x07, 0x2e, 0x8a, 0xf8, 0xfd, 0xa6, 0xb6, 0xd0, 0xd4, 0x0b, 0xca, 0xe4, 0x84, 0x58,
	0xa7, 0xda, 0x1c, 0x1a, 0x0f, 0x48, 0x13, 0x6a, 0x8c, 0x1e, 0xab, 0x92, 0xe9, 0xea, 0x98, 0x08,
	0xbc, 0xa4, 0xc7, 0x80, 0xe7, 0x11, 0x92, 0x49, 0xee, 0xef, 0x83, 0xa7, 0x2c, 0x59, 0x95, 0xf0,
	0x1d, 0x11, 0x30, 0x02, 0x94, 0xed, 0xa9, 0xc7, 0x02, 0xdf, 0x63, 0x80, 0x97, 0x51, 0x5a, 0x5e,
	0x70, 0x4e, 0x2b, 0x24, 0x2f, 0x1f, 0x92, 0xc2, 0xe2, 0x22, 0x9a, 0xf4, 0xe0, 0x90, 0xd7, 0x06,
	0xd8, 0xfe, 0x15, 0xe1, 0xed, 0x0e, 0x63, 0xf9, 0x73, 0x1a, 0x8d, 0x3d, 0x8b, 0xbe, 0x20, 0xf8,
	0x83, 0x86, 0x50, 0xd7, 0x38, 0xb8, 0x34, 0x94, 0x69, 0xc0, 0x5d, 0xfa, 0x68, 0x55, 0xc6, 0xea,
	0x59, 0x25, 0x2f, 0x84, 0x99, 0xbd, 0x2b, 0x79, 0x4b, 0x46, 0xe4, 0xaa, 0xbd, 0xfb, 0xfa, 0xed,
	0x63, 0xe2, 0x3f, 0x63, 0x22, 0xfe, 0xaa, 0xc9, 0x2e, 0xd6, 0xb4, 0x12, 0x3e, 0x42, 0x7f, 0x47,
	0x1e, 0xc4, 0x0b, 0x43, 0x39, 0x2e, 0xba, 0xf4, 0x32, 0x31, 0xc5, 0xb3, 0x8a, 0xbc, 0x3e, 0x49,
	0x3a, 0x8b, 0x67, 0x62, 0xd2, 0x13, 0x11, 0x7d, 0xa0, 0xb6, 0xa9, 0x74, 0x8a, 0xdf, 0x6a, 0x08,
	0x75, 0x2d, 0x3b, 0x62, 0x1c, 0x03, 0xbe, 0xbe, 0x4c, 0xc1, 0x82, 0xa4, 0xbe, 0x5e, 0xce, 0x77,
	0xa9, 0xe5, 0x2c, 0x2e, 0xf0, 0x8b, 0xee, 0x0f, 0x11, 0xea, 0xfa, 0x75, 0x84, 0x82, 0x01, 0x53,
	0xeb, 0xd3, 0x03, 0x1b, 0xf1, 0x44, 0x7c, 0x7c, 0xfb, 0x9a, 0x2f, 0x0d, 0x6d, 0xfe, 0x18, 0x8d,
	0xc7, 0x56, 0xc4, 0x37, 0x87, 0x12, 0xf7, 0xdb, 0x5f, 0x2f, 0x5d, 0x05, 0xaa, 0x9c, 0x6d, 0x4c,
	0x4b, 0x11, 0x53, 0xb8, 0xef, 0xda, 0xf5, 0xec, 0x97, 0xca, 0x84, 0xe3, 0xdb, 0xc4, 0x69, 0xf9,
	0x8c, 0xaf, 0xdd, 0x5d, 0xb9, 0x73, 0x6f, 0x7d, 0x17, 0xcd, 0xd9, 0xbe, 0x3b, 0xac, 0xf6, 0xb6,
	0xf6, 0x6a, 0xa5, 0x49, 0x79, 0xab, 0x5d, 0x97, 0x1b, 0xda, 0x5d, 0x58, 0xab, 0x49, 0x02, 0x6a,
	0xdf, 0x8e, 0xff, 0x34, 0x19, 0x84, 0x07, 0x10, 0x5a, 0x4d, 0xf0, 0xd4, 0x58, 0xfe, 0x92, 0x8f,
	0xe5, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xa0, 0x6e, 0x41, 0x5e, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IdentityClient is the client API for Identity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdentityClient interface {
	// Creates a user.
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	// Retrieves the User with the given uri.
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	// Updates a user.
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	// Deletes a user, their profile, and all of their authored messages.
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Lists all users.
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
}

type identityClient struct {
	cc *grpc.ClientConn
}

func NewIdentityClient(cc *grpc.ClientConn) IdentityClient {
	return &identityClient{cc}
}

func (c *identityClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/google.showcase.v1beta1.Identity/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/google.showcase.v1beta1.Identity/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/google.showcase.v1beta1.Identity/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.showcase.v1beta1.Identity/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/google.showcase.v1beta1.Identity/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServer is the server API for Identity service.
type IdentityServer interface {
	// Creates a user.
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// Retrieves the User with the given uri.
	GetUser(context.Context, *GetUserRequest) (*User, error)
	// Updates a user.
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	// Deletes a user, their profile, and all of their authored messages.
	DeleteUser(context.Context, *DeleteUserRequest) (*empty.Empty, error)
	// Lists all users.
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
}

// UnimplementedIdentityServer can be embedded to have forward compatible implementations.
type UnimplementedIdentityServer struct {
}

func (*UnimplementedIdentityServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedIdentityServer) GetUser(ctx context.Context, req *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedIdentityServer) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedIdentityServer) DeleteUser(ctx context.Context, req *DeleteUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (*UnimplementedIdentityServer) ListUsers(ctx context.Context, req *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}

func RegisterIdentityServer(s *grpc.Server, srv IdentityServer) {
	s.RegisterService(&_Identity_serviceDesc, srv)
}

func _Identity_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.showcase.v1beta1.Identity/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.showcase.v1beta1.Identity/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.showcase.v1beta1.Identity/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.showcase.v1beta1.Identity/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.showcase.v1beta1.Identity/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Identity_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.showcase.v1beta1.Identity",
	HandlerType: (*IdentityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Identity_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Identity_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Identity_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Identity_DeleteUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _Identity_ListUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/showcase/v1beta1/identity.proto",
}
