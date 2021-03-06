// Code generated by protoc-gen-go. DO NOT EDIT.
// source: person.proto

package test

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}

var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) Enum() *Person_PhoneType {
	p := new(Person_PhoneType)
	*p = x
	return p
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}

func (x *Person_PhoneType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Person_PhoneType_value, data, "Person_PhoneType")
	if err != nil {
		return err
	}
	*x = Person_PhoneType(value)
	return nil
}

func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0, 0}
}

type Person struct {
	Name                 *string               `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Id                   *int32                `protobuf:"varint,2,req,name=id" json:"id,omitempty"`
	Email                *string               `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type Person_PhoneNumber struct {
	Number               *string           `protobuf:"bytes,1,req,name=number" json:"number,omitempty"`
	Type                 *Person_PhoneType `protobuf:"varint,2,opt,name=type,enum=test.Person_PhoneType,def=1" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0, 0}
}

func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(m, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

const Default_Person_PhoneNumber_Type Person_PhoneType = Person_HOME

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_Person_PhoneNumber_Type
}

type AddressBook struct {
	People               []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}
func (*AddressBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{1}
}

func (m *AddressBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBook.Unmarshal(m, b)
}
func (m *AddressBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBook.Marshal(b, m, deterministic)
}
func (m *AddressBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBook.Merge(m, src)
}
func (m *AddressBook) XXX_Size() int {
	return xxx_messageInfo_AddressBook.Size(m)
}
func (m *AddressBook) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBook.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBook proto.InternalMessageInfo

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterEnum("test.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
	proto.RegisterType((*Person)(nil), "test.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "test.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "test.AddressBook")
}

func init() { proto.RegisterFile("person.proto", fileDescriptor_4c9e10cf24b1156d) }

var fileDescriptor_4c9e10cf24b1156d = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x50, 0xcf, 0x4b, 0xf3, 0x40,
	0x10, 0xfd, 0xb2, 0xdd, 0x2e, 0x5f, 0x26, 0xa5, 0x84, 0x41, 0xca, 0xe2, 0x69, 0x09, 0x1e, 0x16,
	0x84, 0x50, 0xea, 0xcd, 0x9b, 0x85, 0x82, 0xa2, 0x35, 0x65, 0x11, 0x7a, 0xae, 0x64, 0xc0, 0x60,
	0x93, 0x5d, 0xb2, 0xf1, 0xd0, 0x7f, 0xdd, 0x93, 0xec, 0x36, 0x48, 0xc1, 0xdb, 0xfb, 0xc5, 0x7b,
	0xcc, 0xc0, 0xcc, 0x51, 0xef, 0x6d, 0x57, 0xba, 0xde, 0x0e, 0x16, 0xf9, 0x40, 0x7e, 0x28, 0xbe,
	0x13, 0x10, 0xbb, 0x28, 0x23, 0x02, 0xef, 0x0e, 0x2d, 0xc9, 0x44, 0x31, 0x9d, 0x9a, 0x88, 0x71,
	0x0e, 0xac, 0xa9, 0x25, 0x53, 0x4c, 0x4f, 0x0d, 0x6b, 0x6a, 0xbc, 0x82, 0x29, 0xb5, 0x87, 0xe6,
	0x28, 0x27, 0x2a, 0xd1, 0xa9, 0x39, 0x13, 0x5c, 0x82, 0x70, 0x1f, 0xb6, 0x23, 0x2f, 0xb9, 0x9a,
	0xe8, 0x6c, 0x25, 0xcb, 0xd0, 0x5d, 0x9e, 0x7b, 0xcb, 0x5d, 0xb0, 0x5e, 0xbf, 0xda, 0x77, 0xea,
	0xcd, 0x98, 0xbb, 0xde, 0x43, 0x76, 0x21, 0xe3, 0x02, 0x44, 0x17, 0xd1, 0x38, 0x3e, 0x32, 0x5c,
	0x02, 0x1f, 0x4e, 0x8e, 0x24, 0x53, 0x89, 0x9e, 0xaf, 0x16, 0x7f, 0x6b, 0xdf, 0x4e, 0x8e, 0xee,
	0xf9, 0x63, 0xb5, 0xdd, 0x98, 0x98, 0x2c, 0x6e, 0x21, 0xfd, 0x35, 0x10, 0x40, 0x6c, 0xab, 0xf5,
	0xd3, 0xcb, 0x26, 0xff, 0x87, 0xff, 0x21, 0xc6, 0xf2, 0x24, 0xa0, 0x7d, 0x65, 0x9e, 0x73, 0x56,
	0xdc, 0x41, 0xf6, 0x50, 0xd7, 0x3d, 0x79, 0xbf, 0xb6, 0xf6, 0x13, 0x6f, 0x40, 0x38, 0xb2, 0xee,
	0x18, 0x5e, 0x10, 0xce, 0x98, 0x5d, 0xee, 0x99, 0xd1, 0xfb, 0x09, 0x00, 0x00, 0xff, 0xff, 0x95,
	0xe9, 0x74, 0xb3, 0x46, 0x01, 0x00, 0x00,
}
