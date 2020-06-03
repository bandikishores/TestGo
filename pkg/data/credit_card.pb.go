// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: credit_card.proto

package data

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Timestamp struct {
	Timestamp            *types.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-" yaml:"-" xml:"-" gorm:"-"`
	XXX_unrecognized     []byte           `json:"-" yaml:"-" xml:"-" gorm:"-"`
	XXX_sizecache        int32            `json:"-" yaml:"-" xml:"-" gorm:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc3303e5a2b31807, []int{0}
}
func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timestamp.Unmarshal(m, b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
}
func (m *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(m, src)
}
func (m *Timestamp) XXX_Size() int {
	return xxx_messageInfo_Timestamp.Size(m)
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *Timestamp) GetTimestamp() *types.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type BandiUser struct {
	// @inject_tag: gorm:"primary_key"
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty" gorm:"primary_key"`
	// @inject_tag: gorm:"ForeignKey:UserName"
	CreditCards          []*BandiCreditCard `protobuf:"bytes,2,rep,name=CreditCards,proto3" json:"CreditCards,omitempty" gorm:"ForeignKey:UserName"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-" yaml:"-" xml:"-" gorm:"-"`
	XXX_unrecognized     []byte             `json:"-" yaml:"-" xml:"-" gorm:"-"`
	XXX_sizecache        int32              `json:"-" yaml:"-" xml:"-" gorm:"-"`
}

func (m *BandiUser) Reset()         { *m = BandiUser{} }
func (m *BandiUser) String() string { return proto.CompactTextString(m) }
func (*BandiUser) ProtoMessage()    {}
func (*BandiUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc3303e5a2b31807, []int{1}
}
func (m *BandiUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BandiUser.Unmarshal(m, b)
}
func (m *BandiUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BandiUser.Marshal(b, m, deterministic)
}
func (m *BandiUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BandiUser.Merge(m, src)
}
func (m *BandiUser) XXX_Size() int {
	return xxx_messageInfo_BandiUser.Size(m)
}
func (m *BandiUser) XXX_DiscardUnknown() {
	xxx_messageInfo_BandiUser.DiscardUnknown(m)
}

var xxx_messageInfo_BandiUser proto.InternalMessageInfo

func (m *BandiUser) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *BandiUser) GetCreditCards() []*BandiCreditCard {
	if m != nil {
		return m.CreditCards
	}
	return nil
}

// Contains Composite Primary Key and Foreign Key
// User Details "belongs to" an User
type BandiUserDetails struct {
	// @inject_tag: gorm:"primary_key"
	Firstname string `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty" gorm:"primary_key"`
	// @inject_tag: gorm:"primary_key"
	Lastname string `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty" gorm:"primary_key"`
	// @inject_tag: gorm:"ForeignKey:ForeignKeyUserName;association_foreignkey:Username"
	User                 *BandiUser `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty" gorm:"ForeignKey:ForeignKeyUserName;association_foreignkey:Username"`
	ForeignKeyUserName   string     `protobuf:"bytes,4,opt,name=foreignKeyUserName,proto3" json:"foreignKeyUserName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-" yaml:"-" xml:"-" gorm:"-"`
	XXX_unrecognized     []byte     `json:"-" yaml:"-" xml:"-" gorm:"-"`
	XXX_sizecache        int32      `json:"-" yaml:"-" xml:"-" gorm:"-"`
}

func (m *BandiUserDetails) Reset()         { *m = BandiUserDetails{} }
func (m *BandiUserDetails) String() string { return proto.CompactTextString(m) }
func (*BandiUserDetails) ProtoMessage()    {}
func (*BandiUserDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc3303e5a2b31807, []int{2}
}
func (m *BandiUserDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BandiUserDetails.Unmarshal(m, b)
}
func (m *BandiUserDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BandiUserDetails.Marshal(b, m, deterministic)
}
func (m *BandiUserDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BandiUserDetails.Merge(m, src)
}
func (m *BandiUserDetails) XXX_Size() int {
	return xxx_messageInfo_BandiUserDetails.Size(m)
}
func (m *BandiUserDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_BandiUserDetails.DiscardUnknown(m)
}

var xxx_messageInfo_BandiUserDetails proto.InternalMessageInfo

func (m *BandiUserDetails) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *BandiUserDetails) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *BandiUserDetails) GetUser() *BandiUser {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *BandiUserDetails) GetForeignKeyUserName() string {
	if m != nil {
		return m.ForeignKeyUserName
	}
	return ""
}

type BandiCreditCard struct {
	// @inject_tag: gorm:"unique;not null;size:16;unique_index"
	Number string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty" gorm:"unique;not null;size:16;unique_index"`
	// @inject_tag: gorm:"not null" bandi:"only4digits"
	Cvv int32 `protobuf:"varint,2,opt,name=cvv,proto3" json:"cvv,omitempty" gorm:"not null" bandi:"only4digits"`
	// @inject_tag: gorm:"type:timestamp;index:expiry"
	Expiry *Timestamp `protobuf:"bytes,3,opt,name=expiry,proto3" json:"expiry,omitempty" gorm:"type:timestamp;index:expiry"`
	// @inject_tag: gorm:"-"
	IgnoreME string `protobuf:"bytes,4,opt,name=ignoreME,proto3" json:"ignoreMEJson" xml:"ignoreMEXML" gorm:"-"`
	// @inject_tag: sql:"type:string REFERENCES bandi_users(username)"
	UserName string `protobuf:"bytes,5,opt,name=userName,proto3" json:"userName,omitempty" sql:"type:string REFERENCES bandi_users(username)"`
	// @inject_tag: gorm:"primary_key;AUTO_INCREMENT"
	ID                   string   `protobuf:"bytes,6,opt,name=ID,proto3" json:"ID,omitempty" gorm:"primary_key;AUTO_INCREMENT"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-" xml:"-" gorm:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-" xml:"-" gorm:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-" xml:"-" gorm:"-"`
}

func (m *BandiCreditCard) Reset()         { *m = BandiCreditCard{} }
func (m *BandiCreditCard) String() string { return proto.CompactTextString(m) }
func (*BandiCreditCard) ProtoMessage()    {}
func (*BandiCreditCard) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc3303e5a2b31807, []int{3}
}
func (m *BandiCreditCard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BandiCreditCard.Unmarshal(m, b)
}
func (m *BandiCreditCard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BandiCreditCard.Marshal(b, m, deterministic)
}
func (m *BandiCreditCard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BandiCreditCard.Merge(m, src)
}
func (m *BandiCreditCard) XXX_Size() int {
	return xxx_messageInfo_BandiCreditCard.Size(m)
}
func (m *BandiCreditCard) XXX_DiscardUnknown() {
	xxx_messageInfo_BandiCreditCard.DiscardUnknown(m)
}

var xxx_messageInfo_BandiCreditCard proto.InternalMessageInfo

func (m *BandiCreditCard) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *BandiCreditCard) GetCvv() int32 {
	if m != nil {
		return m.Cvv
	}
	return 0
}

func (m *BandiCreditCard) GetExpiry() *Timestamp {
	if m != nil {
		return m.Expiry
	}
	return nil
}

func (m *BandiCreditCard) GetIgnoreME() string {
	if m != nil {
		return m.IgnoreME
	}
	return ""
}

func (m *BandiCreditCard) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *BandiCreditCard) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type GetBandiUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-" xml:"-" gorm:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-" xml:"-" gorm:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-" xml:"-" gorm:"-"`
}

func (m *GetBandiUserRequest) Reset()         { *m = GetBandiUserRequest{} }
func (m *GetBandiUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetBandiUserRequest) ProtoMessage()    {}
func (*GetBandiUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc3303e5a2b31807, []int{4}
}
func (m *GetBandiUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBandiUserRequest.Unmarshal(m, b)
}
func (m *GetBandiUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBandiUserRequest.Marshal(b, m, deterministic)
}
func (m *GetBandiUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBandiUserRequest.Merge(m, src)
}
func (m *GetBandiUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetBandiUserRequest.Size(m)
}
func (m *GetBandiUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBandiUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBandiUserRequest proto.InternalMessageInfo

func (m *GetBandiUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetBandiUserResponse struct {
	User                 *BandiUser `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-" yaml:"-" xml:"-" gorm:"-"`
	XXX_unrecognized     []byte     `json:"-" yaml:"-" xml:"-" gorm:"-"`
	XXX_sizecache        int32      `json:"-" yaml:"-" xml:"-" gorm:"-"`
}

func (m *GetBandiUserResponse) Reset()         { *m = GetBandiUserResponse{} }
func (m *GetBandiUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetBandiUserResponse) ProtoMessage()    {}
func (*GetBandiUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc3303e5a2b31807, []int{5}
}
func (m *GetBandiUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBandiUserResponse.Unmarshal(m, b)
}
func (m *GetBandiUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBandiUserResponse.Marshal(b, m, deterministic)
}
func (m *GetBandiUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBandiUserResponse.Merge(m, src)
}
func (m *GetBandiUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetBandiUserResponse.Size(m)
}
func (m *GetBandiUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBandiUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBandiUserResponse proto.InternalMessageInfo

func (m *GetBandiUserResponse) GetUser() *BandiUser {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateBandiUserRequest struct {
	User                 *BandiUser `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-" yaml:"-" xml:"-" gorm:"-"`
	XXX_unrecognized     []byte     `json:"-" yaml:"-" xml:"-" gorm:"-"`
	XXX_sizecache        int32      `json:"-" yaml:"-" xml:"-" gorm:"-"`
}

func (m *CreateBandiUserRequest) Reset()         { *m = CreateBandiUserRequest{} }
func (m *CreateBandiUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBandiUserRequest) ProtoMessage()    {}
func (*CreateBandiUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc3303e5a2b31807, []int{6}
}
func (m *CreateBandiUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBandiUserRequest.Unmarshal(m, b)
}
func (m *CreateBandiUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBandiUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateBandiUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBandiUserRequest.Merge(m, src)
}
func (m *CreateBandiUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBandiUserRequest.Size(m)
}
func (m *CreateBandiUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBandiUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBandiUserRequest proto.InternalMessageInfo

func (m *CreateBandiUserRequest) GetUser() *BandiUser {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateBandiUserResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-" xml:"-" gorm:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-" xml:"-" gorm:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-" xml:"-" gorm:"-"`
}

func (m *CreateBandiUserResponse) Reset()         { *m = CreateBandiUserResponse{} }
func (m *CreateBandiUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateBandiUserResponse) ProtoMessage()    {}
func (*CreateBandiUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc3303e5a2b31807, []int{7}
}
func (m *CreateBandiUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBandiUserResponse.Unmarshal(m, b)
}
func (m *CreateBandiUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBandiUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateBandiUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBandiUserResponse.Merge(m, src)
}
func (m *CreateBandiUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateBandiUserResponse.Size(m)
}
func (m *CreateBandiUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBandiUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBandiUserResponse proto.InternalMessageInfo

func (m *CreateBandiUserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Timestamp)(nil), "data.Timestamp")
	golang_proto.RegisterType((*Timestamp)(nil), "data.Timestamp")
	proto.RegisterType((*BandiUser)(nil), "data.BandiUser")
	golang_proto.RegisterType((*BandiUser)(nil), "data.BandiUser")
	proto.RegisterType((*BandiUserDetails)(nil), "data.BandiUserDetails")
	golang_proto.RegisterType((*BandiUserDetails)(nil), "data.BandiUserDetails")
	proto.RegisterType((*BandiCreditCard)(nil), "data.BandiCreditCard")
	golang_proto.RegisterType((*BandiCreditCard)(nil), "data.BandiCreditCard")
	proto.RegisterType((*GetBandiUserRequest)(nil), "data.GetBandiUserRequest")
	golang_proto.RegisterType((*GetBandiUserRequest)(nil), "data.GetBandiUserRequest")
	proto.RegisterType((*GetBandiUserResponse)(nil), "data.GetBandiUserResponse")
	golang_proto.RegisterType((*GetBandiUserResponse)(nil), "data.GetBandiUserResponse")
	proto.RegisterType((*CreateBandiUserRequest)(nil), "data.CreateBandiUserRequest")
	golang_proto.RegisterType((*CreateBandiUserRequest)(nil), "data.CreateBandiUserRequest")
	proto.RegisterType((*CreateBandiUserResponse)(nil), "data.CreateBandiUserResponse")
	golang_proto.RegisterType((*CreateBandiUserResponse)(nil), "data.CreateBandiUserResponse")
}

func init() { proto.RegisterFile("credit_card.proto", fileDescriptor_dc3303e5a2b31807) }
func init() { golang_proto.RegisterFile("credit_card.proto", fileDescriptor_dc3303e5a2b31807) }

var fileDescriptor_dc3303e5a2b31807 = []byte{
	// 691 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4f, 0x4f, 0x13, 0x41,
	0x14, 0x4f, 0xff, 0xd0, 0xd0, 0x81, 0x88, 0x8c, 0x50, 0x6a, 0x53, 0x64, 0xb2, 0x44, 0x85, 0x06,
	0xba, 0xc9, 0x7a, 0xd0, 0x40, 0x3c, 0x74, 0x29, 0x9a, 0xaa, 0x10, 0xb3, 0xd6, 0xc4, 0x78, 0xd1,
	0x69, 0x77, 0xba, 0x8c, 0xb6, 0x33, 0x65, 0x66, 0x5a, 0x20, 0xc6, 0x8b, 0x1f, 0xa1, 0x7e, 0x02,
	0xaf, 0x7e, 0x0a, 0xbf, 0x83, 0x77, 0x0e, 0xea, 0xc9, 0xa3, 0x9f, 0xc0, 0xcc, 0xec, 0x76, 0xdb,
	0x42, 0x31, 0xde, 0xe6, 0xbd, 0xdf, 0x7b, 0xbf, 0xdf, 0xef, 0xbd, 0x7d, 0x59, 0xb0, 0xd8, 0x14,
	0xc4, 0xa7, 0xea, 0x4d, 0x13, 0x0b, 0xbf, 0xdc, 0x15, 0x5c, 0x71, 0x98, 0xf6, 0xb1, 0xc2, 0x85,
	0x2d, 0x13, 0x34, 0xb7, 0x03, 0xc2, 0xb6, 0xe5, 0x09, 0x0e, 0x02, 0x22, 0x6c, 0xde, 0x55, 0x94,
	0x33, 0x69, 0x63, 0xc6, 0xb8, 0xc2, 0xe6, 0x1d, 0xf6, 0x14, 0x8a, 0x01, 0xe7, 0x41, 0x9b, 0xd8,
	0xb8, 0x4b, 0xa7, 0xa0, 0x4b, 0x01, 0x0f, 0xb8, 0x79, 0xda, 0xfa, 0x15, 0x65, 0xd7, 0xa2, 0x1e,
	0x13, 0x35, 0x7a, 0x2d, 0x5b, 0xd1, 0x0e, 0x91, 0x0a, 0x77, 0xba, 0x61, 0x81, 0xb5, 0x0f, 0xb2,
	0xf5, 0x61, 0x0a, 0x3e, 0x00, 0xd9, 0x18, 0xcf, 0x27, 0x50, 0x62, 0x63, 0xce, 0x29, 0x94, 0x43,
	0x86, 0xf2, 0x90, 0xa1, 0x1c, 0x97, 0x7b, 0xa3, 0x62, 0xeb, 0x2d, 0xc8, 0xba, 0x98, 0xf9, 0xf4,
	0xa5, 0x24, 0x02, 0x16, 0xc0, 0x6c, 0x4f, 0x12, 0xc1, 0x70, 0x87, 0x18, 0x96, 0xac, 0x17, 0xc7,
	0xf0, 0x3e, 0x98, 0xdb, 0x33, 0xdb, 0xd8, 0xc3, 0xc2, 0x97, 0xf9, 0x24, 0x4a, 0x6d, 0xcc, 0x39,
	0xcb, 0x65, 0xbd, 0x8e, 0xb2, 0x61, 0x18, 0xa1, 0xde, 0x78, 0xa5, 0xf5, 0x25, 0x01, 0xae, 0xc7,
	0x12, 0x55, 0xa2, 0x30, 0x6d, 0x4b, 0x58, 0x04, 0xd9, 0x16, 0x15, 0x52, 0x8d, 0x49, 0x8d, 0x12,
	0xda, 0x47, 0x1b, 0x47, 0x60, 0x32, 0xf4, 0x31, 0x8c, 0xe1, 0x3a, 0x48, 0x6b, 0x4f, 0xf9, 0x94,
	0x99, 0x72, 0x61, 0xcc, 0x80, 0xe6, 0xf7, 0x0c, 0x08, 0xcb, 0x00, 0xb6, 0xb8, 0x20, 0x34, 0x60,
	0x4f, 0xc9, 0x99, 0xce, 0x1f, 0x6a, 0xaa, 0xb4, 0xa1, 0x9a, 0x82, 0x58, 0x5f, 0x53, 0x60, 0xe1,
	0xc2, 0x10, 0x70, 0x17, 0x64, 0x58, 0xaf, 0xd3, 0x20, 0x22, 0xf4, 0xe7, 0xae, 0x0f, 0x2a, 0xa8,
	0x94, 0x39, 0x34, 0x29, 0x27, 0x57, 0x3f, 0x22, 0x28, 0x2c, 0x46, 0xba, 0x1a, 0x85, 0x79, 0x2f,
	0x6a, 0x81, 0x0e, 0x48, 0x35, 0xfb, 0x7d, 0x63, 0x7e, 0xc6, 0x45, 0x83, 0xca, 0x6a, 0x49, 0xc7,
	0x4e, 0xee, 0x39, 0x96, 0xf2, 0x84, 0x0b, 0x1f, 0xb5, 0xb8, 0x40, 0xea, 0x88, 0x4a, 0x64, 0x5c,
	0x6b, 0x10, 0x1e, 0x82, 0x0c, 0x39, 0xed, 0x52, 0x71, 0x36, 0x39, 0x5b, 0xfc, 0xd9, 0xdc, 0xdb,
	0x83, 0x8a, 0x55, 0x9a, 0x1d, 0x52, 0x5c, 0x49, 0x16, 0xb1, 0xc0, 0x3e, 0x98, 0xa5, 0x01, 0xe3,
	0x82, 0x1c, 0xec, 0x87, 0xa3, 0xbb, 0xaf, 0x07, 0x15, 0xa7, 0x14, 0x27, 0x9d, 0x3b, 0x75, 0x22,
	0x15, 0xaa, 0x31, 0xa5, 0x3f, 0x71, 0x1b, 0x3d, 0xa2, 0xa4, 0xed, 0x6f, 0xa1, 0x2a, 0x67, 0x77,
	0x15, 0x92, 0x8a, 0x0b, 0x82, 0x28, 0x43, 0x55, 0xf7, 0xf7, 0xf9, 0xda, 0xfc, 0xb0, 0xe7, 0x89,
	0xe4, 0xec, 0xcf, 0xf9, 0xda, 0xe2, 0x69, 0xa7, 0xbd, 0x63, 0x0d, 0x93, 0xaf, 0x0e, 0x9e, 0x59,
	0x5e, 0x4c, 0x3b, 0xbc, 0x22, 0xb3, 0xf2, 0x99, 0xd1, 0x15, 0xe9, 0x18, 0xba, 0x20, 0x59, 0xab,
	0xe6, 0x33, 0xc6, 0x8d, 0x33, 0xa8, 0xd8, 0xa5, 0x64, 0xad, 0xea, 0x6c, 0x56, 0x50, 0x8f, 0xd1,
	0xe3, 0x1e, 0x41, 0xb5, 0xaa, 0xed, 0xf1, 0x13, 0x54, 0xe7, 0xef, 0x09, 0x8b, 0xa6, 0x9a, 0x58,
	0xb3, 0x97, 0xac, 0x55, 0xad, 0x4d, 0x70, 0xe3, 0x31, 0x51, 0xa3, 0x4f, 0x4e, 0x8e, 0x7b, 0x44,
	0x2a, 0x08, 0x41, 0x7a, 0xec, 0x9a, 0xcc, 0xdb, 0xda, 0x05, 0x4b, 0x93, 0xa5, 0xb2, 0xcb, 0x99,
	0x1c, 0x1d, 0x51, 0xe2, 0x1f, 0x47, 0x64, 0x3d, 0x04, 0xb9, 0x3d, 0x41, 0xb0, 0x22, 0x97, 0xa4,
	0xfe, 0xab, 0x7d, 0x1b, 0xac, 0x5c, 0x6a, 0x8f, 0xe4, 0xa7, 0x58, 0x75, 0x7e, 0x24, 0xc0, 0xbc,
	0xa9, 0x7c, 0x41, 0x44, 0x9f, 0x36, 0x09, 0xf4, 0xc1, 0xfc, 0xb8, 0x77, 0x78, 0x33, 0x94, 0x99,
	0x32, 0x7a, 0xa1, 0x30, 0x0d, 0x0a, 0xb5, 0xac, 0xd5, 0x4f, 0xdf, 0x7f, 0x7d, 0x4e, 0xae, 0xc0,
	0x65, 0x9b, 0x37, 0xde, 0x91, 0xa6, 0x92, 0xb6, 0x76, 0x27, 0xed, 0x0f, 0x5a, 0xf5, 0x23, 0x6c,
	0x83, 0x85, 0x0b, 0x2e, 0x61, 0x31, 0x64, 0x9b, 0x3e, 0x7b, 0x61, 0xf5, 0x0a, 0x34, 0x92, 0x2b,
	0x1a, 0xb9, 0x9c, 0x75, 0x6d, 0x52, 0x6e, 0xc7, 0xec, 0xc4, 0x4d, 0x7f, 0xfb, 0x79, 0x2b, 0xd1,
	0xc8, 0x98, 0x5f, 0xd2, 0xbd, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xa4, 0x49, 0x6b, 0x5f,
	0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BandiServiceClient is the client API for BandiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BandiServiceClient interface {
	GetBandiUser(ctx context.Context, in *GetBandiUserRequest, opts ...grpc.CallOption) (*GetBandiUserResponse, error)
	CreateBandiUser(ctx context.Context, in *CreateBandiUserRequest, opts ...grpc.CallOption) (*CreateBandiUserResponse, error)
}

type bandiServiceClient struct {
	cc *grpc.ClientConn
}

func NewBandiServiceClient(cc *grpc.ClientConn) BandiServiceClient {
	return &bandiServiceClient{cc}
}

func (c *bandiServiceClient) GetBandiUser(ctx context.Context, in *GetBandiUserRequest, opts ...grpc.CallOption) (*GetBandiUserResponse, error) {
	out := new(GetBandiUserResponse)
	err := c.cc.Invoke(ctx, "/data.BandiService/GetBandiUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bandiServiceClient) CreateBandiUser(ctx context.Context, in *CreateBandiUserRequest, opts ...grpc.CallOption) (*CreateBandiUserResponse, error) {
	out := new(CreateBandiUserResponse)
	err := c.cc.Invoke(ctx, "/data.BandiService/CreateBandiUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BandiServiceServer is the server API for BandiService service.
type BandiServiceServer interface {
	GetBandiUser(context.Context, *GetBandiUserRequest) (*GetBandiUserResponse, error)
	CreateBandiUser(context.Context, *CreateBandiUserRequest) (*CreateBandiUserResponse, error)
}

// UnimplementedBandiServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBandiServiceServer struct {
}

func (*UnimplementedBandiServiceServer) GetBandiUser(ctx context.Context, req *GetBandiUserRequest) (*GetBandiUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBandiUser not implemented")
}
func (*UnimplementedBandiServiceServer) CreateBandiUser(ctx context.Context, req *CreateBandiUserRequest) (*CreateBandiUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBandiUser not implemented")
}

func RegisterBandiServiceServer(s *grpc.Server, srv BandiServiceServer) {
	s.RegisterService(&_BandiService_serviceDesc, srv)
}

func _BandiService_GetBandiUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBandiUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BandiServiceServer).GetBandiUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.BandiService/GetBandiUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BandiServiceServer).GetBandiUser(ctx, req.(*GetBandiUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BandiService_CreateBandiUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBandiUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BandiServiceServer).CreateBandiUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.BandiService/CreateBandiUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BandiServiceServer).CreateBandiUser(ctx, req.(*CreateBandiUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BandiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "data.BandiService",
	HandlerType: (*BandiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBandiUser",
			Handler:    _BandiService_GetBandiUser_Handler,
		},
		{
			MethodName: "CreateBandiUser",
			Handler:    _BandiService_CreateBandiUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "credit_card.proto",
}
