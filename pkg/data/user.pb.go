// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package data

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

type User struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName          string   `protobuf:"bytes,2,opt,name=displayName,proto3" json:"displayName,omitempty"`
	ContactAddress       *Address `protobuf:"bytes,3,opt,name=contactAddress,proto3" json:"contactAddress,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	OktaID               string   `protobuf:"bytes,5,opt,name=oktaID,proto3" json:"oktaID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
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

func (m *User) GetContactAddress() *Address {
	if m != nil {
		return m.ContactAddress
	}
	return nil
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetOktaID() string {
	if m != nil {
		return m.OktaID
	}
	return ""
}

type CreateUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	OrgName              string   `protobuf:"bytes,2,opt,name=orgName,proto3" json:"orgName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
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

func (m *CreateUserRequest) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

type CreateUserResponse struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status               *status.Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateUserResponse) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type UpdateUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	OrgName              string   `protobuf:"bytes,2,opt,name=orgName,proto3" json:"orgName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
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

func (m *UpdateUserRequest) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

type UpdateUserResponse struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status               *status.Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateUserResponse) Reset()         { *m = UpdateUserResponse{} }
func (m *UpdateUserResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateUserResponse) ProtoMessage()    {}
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *UpdateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserResponse.Unmarshal(m, b)
}
func (m *UpdateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserResponse.Marshal(b, m, deterministic)
}
func (m *UpdateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserResponse.Merge(m, src)
}
func (m *UpdateUserResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateUserResponse.Size(m)
}
func (m *UpdateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserResponse proto.InternalMessageInfo

func (m *UpdateUserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateUserResponse) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type DeleteUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OrgName              string   `protobuf:"bytes,2,opt,name=orgName,proto3" json:"orgName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
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

func (m *DeleteUserRequest) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

type DeleteUserResponse struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status               *status.Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeleteUserResponse) Reset()         { *m = DeleteUserResponse{} }
func (m *DeleteUserResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteUserResponse) ProtoMessage()    {}
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *DeleteUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserResponse.Unmarshal(m, b)
}
func (m *DeleteUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserResponse.Marshal(b, m, deterministic)
}
func (m *DeleteUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserResponse.Merge(m, src)
}
func (m *DeleteUserResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteUserResponse.Size(m)
}
func (m *DeleteUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserResponse proto.InternalMessageInfo

func (m *DeleteUserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeleteUserResponse) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type GetUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OrgName              string   `protobuf:"bytes,2,opt,name=orgName,proto3" json:"orgName,omitempty"`
	QueryParam1          string   `protobuf:"bytes,3,opt,name=queryParam1,proto3" json:"queryParam1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
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

func (m *GetUserRequest) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

func (m *GetUserRequest) GetQueryParam1() string {
	if m != nil {
		return m.QueryParam1
	}
	return ""
}

type GetUserResponse struct {
	User                 *User          `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	Status               *status.Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GetUserResponse) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "data.User")
	proto.RegisterType((*CreateUserRequest)(nil), "data.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "data.CreateUserResponse")
	proto.RegisterType((*UpdateUserRequest)(nil), "data.UpdateUserRequest")
	proto.RegisterType((*UpdateUserResponse)(nil), "data.UpdateUserResponse")
	proto.RegisterType((*DeleteUserRequest)(nil), "data.DeleteUserRequest")
	proto.RegisterType((*DeleteUserResponse)(nil), "data.DeleteUserResponse")
	proto.RegisterType((*GetUserRequest)(nil), "data.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "data.GetUserResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 635 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4f, 0x4f, 0x14, 0x31,
	0x14, 0xcf, 0xb0, 0xcb, 0x12, 0xde, 0x46, 0x94, 0x06, 0x65, 0xdd, 0x18, 0x32, 0x19, 0x63, 0xc4,
	0xd5, 0x9d, 0x81, 0x05, 0x2f, 0x98, 0x18, 0x41, 0x12, 0xe3, 0x41, 0xa3, 0x83, 0x1c, 0x4d, 0x2c,
	0x33, 0x2f, 0xb3, 0x03, 0xbb, 0x6d, 0x69, 0x3b, 0x22, 0x12, 0x2e, 0x7e, 0x04, 0xbd, 0x79, 0xf3,
	0x53, 0x78, 0xf0, 0xe8, 0x47, 0xf0, 0x2b, 0xf8, 0x41, 0xcc, 0xb4, 0xb3, 0x32, 0x30, 0x88, 0x62,
	0xb8, 0xf5, 0xbd, 0xd7, 0xfe, 0xfe, 0xbc, 0xbe, 0x16, 0x20, 0x53, 0x28, 0x7d, 0x21, 0xb9, 0xe6,
	0xa4, 0x1e, 0x53, 0x4d, 0xdb, 0xf7, 0x4c, 0x10, 0x75, 0x13, 0x64, 0x5d, 0xb5, 0x47, 0x93, 0x04,
	0x65, 0xc0, 0x85, 0x4e, 0x39, 0x53, 0x01, 0x65, 0x8c, 0x6b, 0x6a, 0xd6, 0xf6, 0x4c, 0x7b, 0x36,
	0xe1, 0x3c, 0x19, 0x60, 0x20, 0x45, 0x14, 0x28, 0x4d, 0x75, 0x36, 0x2a, 0xdc, 0x28, 0x0a, 0x54,
	0xa4, 0xa7, 0x1c, 0x9b, 0xa6, 0x71, 0x2c, 0x51, 0xa9, 0x2d, 0xce, 0x77, 0x6c, 0xca, 0xfb, 0xe2,
	0x40, 0x7d, 0x53, 0xa1, 0x24, 0x04, 0xea, 0x8c, 0x0e, 0xb1, 0xe5, 0xb8, 0xce, 0xfc, 0x64, 0x68,
	0xd6, 0xc4, 0x85, 0x66, 0x9c, 0x2a, 0x31, 0xa0, 0xfb, 0xcf, 0xf3, 0xd2, 0x98, 0x29, 0x95, 0x53,
	0xe4, 0x3e, 0x4c, 0x45, 0x9c, 0x69, 0x1a, 0xe9, 0x55, 0x0b, 0xdd, 0xaa, 0xb9, 0xce, 0x7c, 0xb3,
	0x77, 0xc9, 0xcf, 0x5d, 0xf9, 0x45, 0x32, 0x3c, 0xb1, 0x89, 0xcc, 0xc0, 0x38, 0x0e, 0x69, 0x3a,
	0x68, 0xd5, 0x0d, 0xa4, 0x0d, 0xc8, 0x35, 0x68, 0xf0, 0x1d, 0x4d, 0x9f, 0xae, 0xb7, 0xc6, 0x4d,
	0xba, 0x88, 0xbc, 0x67, 0x30, 0xfd, 0x58, 0x22, 0xd5, 0x98, 0x0b, 0x0d, 0x71, 0x37, 0x43, 0xa5,
	0xc9, 0x9c, 0xd5, 0x6d, 0xf4, 0x36, 0x7b, 0x60, 0xf9, 0xcc, 0x06, 0xeb, 0xa7, 0x05, 0x13, 0x5c,
	0x26, 0x25, 0xdd, 0xa3, 0xd0, 0x7b, 0x05, 0xa4, 0x0c, 0xa7, 0x04, 0x67, 0x0a, 0x4f, 0xf5, 0xdf,
	0x81, 0x86, 0xed, 0xae, 0x81, 0x68, 0xf6, 0x88, 0x6f, 0xdb, 0xeb, 0x4b, 0x11, 0xf9, 0x1b, 0xa6,
	0x12, 0x16, 0x3b, 0x72, 0x91, 0x9b, 0x22, 0xbe, 0x48, 0x91, 0x65, 0xb8, 0x0b, 0x12, 0xb9, 0x0a,
	0xd3, 0xeb, 0x38, 0xc0, 0xe3, 0x22, 0x4f, 0x03, 0x3d, 0x53, 0x58, 0x19, 0xe2, 0x82, 0x84, 0xbd,
	0x81, 0xa9, 0x27, 0xa8, 0xff, 0x5b, 0x55, 0x3e, 0xa9, 0xbb, 0x19, 0xca, 0xfd, 0x17, 0x54, 0xd2,
	0xe1, 0xa2, 0x19, 0xc2, 0xc9, 0xb0, 0x9c, 0xf2, 0x5e, 0xc3, 0xe5, 0xdf, 0x0c, 0x85, 0xe8, 0xbf,
	0xdd, 0xce, 0x39, 0x0c, 0xf4, 0xbe, 0xd7, 0xa0, 0x99, 0x1f, 0xda, 0x40, 0xf9, 0x36, 0x8d, 0x90,
	0x20, 0x4c, 0x14, 0x74, 0x64, 0xc6, 0x02, 0x1f, 0xf7, 0xd7, 0xbe, 0x7a, 0x22, 0x6b, 0x35, 0x79,
	0xdd, 0x0f, 0x3f, 0x7e, 0x7e, 0x1a, 0xbb, 0x4d, 0x6e, 0x05, 0x5c, 0x26, 0x94, 0xa5, 0xef, 0xed,
	0x03, 0x0e, 0x0e, 0x0a, 0xa7, 0x87, 0x41, 0xfe, 0x75, 0xa8, 0xe0, 0x20, 0x6f, 0xc8, 0x21, 0xd9,
	0x06, 0x38, 0x9a, 0x65, 0x32, 0x6b, 0x31, 0x2b, 0x8f, 0xa5, 0xdd, 0xaa, 0x16, 0x0a, 0xbe, 0x3b,
	0x86, 0xef, 0xa6, 0x37, 0x77, 0x36, 0xdf, 0x8a, 0xd3, 0x21, 0x0c, 0xe0, 0xe8, 0xe6, 0x47, 0x5c,
	0x95, 0x71, 0x1a, 0x71, 0x55, 0x87, 0x64, 0xe4, 0xad, 0xf3, 0x8f, 0xde, 0x76, 0xa1, 0xb9, 0xa1,
	0x25, 0xd2, 0x61, 0x0e, 0xa2, 0xce, 0xd7, 0xc6, 0x25, 0x43, 0xd5, 0x25, 0x77, 0xff, 0x48, 0xa5,
	0x0c, 0x74, 0x99, 0x70, 0xc1, 0x59, 0xfb, 0xe6, 0x7c, 0x5c, 0xfd, 0xea, 0x90, 0x97, 0x00, 0x6b,
	0x94, 0xc5, 0xa9, 0x9b, 0x63, 0x7a, 0x0f, 0x61, 0xdc, 0x44, 0xa4, 0xdd, 0xd7, 0x5a, 0xa8, 0x95,
	0x20, 0xd8, 0xca, 0xc3, 0x9d, 0x54, 0xf5, 0xb9, 0x44, 0xe5, 0x47, 0x7c, 0x18, 0xb4, 0xaf, 0xab,
	0x4c, 0x08, 0x2e, 0xf5, 0xa3, 0x4a, 0xad, 0x57, 0x5b, 0xf4, 0x17, 0x3a, 0x8e, 0xd3, 0xbb, 0x42,
	0x85, 0x18, 0xa4, 0x91, 0x51, 0x12, 0x6c, 0x2b, 0xce, 0x56, 0x2a, 0x99, 0xf0, 0x01, 0xd4, 0x96,
	0x17, 0x96, 0xc9, 0x32, 0x74, 0x42, 0xd4, 0x99, 0x64, 0x18, 0xbb, 0x7b, 0x7d, 0x64, 0xae, 0xee,
	0xa3, 0x2b, 0x51, 0xf1, 0x4c, 0x46, 0xe8, 0xc6, 0x1c, 0x95, 0xcb, 0xb8, 0x76, 0xf1, 0x5d, 0xaa,
	0xb4, 0x4f, 0x1a, 0x50, 0xff, 0x3c, 0xe6, 0x4c, 0x6c, 0x35, 0xcc, 0x97, 0xbe, 0xf4, 0x2b, 0x00,
	0x00, 0xff, 0xff, 0xf0, 0x7b, 0x4b, 0xfe, 0x5e, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	StreamUsers(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (UserService_StreamUsersClient, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/data.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/data.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/data.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) StreamUsers(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (UserService_StreamUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UserService_serviceDesc.Streams[0], "/data.UserService/StreamUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceStreamUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_StreamUsersClient interface {
	Recv() (*GetUserResponse, error)
	grpc.ClientStream
}

type userServiceStreamUsersClient struct {
	grpc.ClientStream
}

func (x *userServiceStreamUsersClient) Recv() (*GetUserResponse, error) {
	m := new(GetUserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	StreamUsers(*GetUserRequest, UserService_StreamUsersServer) error
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServiceServer) DeleteUser(ctx context.Context, req *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (*UnimplementedUserServiceServer) StreamUsers(req *GetUserRequest, srv UserService_StreamUsersServer) error {
	return status1.Errorf(codes.Unimplemented, "method StreamUsers not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_StreamUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetUserRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).StreamUsers(m, &userServiceStreamUsersServer{stream})
}

type UserService_StreamUsersServer interface {
	Send(*GetUserResponse) error
	grpc.ServerStream
}

type userServiceStreamUsersServer struct {
	grpc.ServerStream
}

func (x *userServiceStreamUsersServer) Send(m *GetUserResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "data.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamUsers",
			Handler:       _UserService_StreamUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "user.proto",
}
