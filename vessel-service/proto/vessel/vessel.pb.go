// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vessel/vessel.proto

package vessel

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=max_weight,json=maxWeight" json:"max_weight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available" json:"available,omitempty"`
	OwnerId              string   `protobuf:"bytes,6,opt,name=owner_id,json=ownerId" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_vessel_59b3339503f4b360, []int{0}
}
func (m *Vessel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vessel.Unmarshal(m, b)
}
func (m *Vessel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vessel.Marshal(b, m, deterministic)
}
func (dst *Vessel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vessel.Merge(dst, src)
}
func (m *Vessel) XXX_Size() int {
	return xxx_messageInfo_Vessel.Size(m)
}
func (m *Vessel) XXX_DiscardUnknown() {
	xxx_messageInfo_Vessel.DiscardUnknown(m)
}

var xxx_messageInfo_Vessel proto.InternalMessageInfo

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type Specification struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,2,opt,name=max_weight,json=maxWeight" json:"max_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_vessel_59b3339503f4b360, []int{1}
}
func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (dst *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(dst, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	Vessel               *Vessel   `protobuf:"bytes,1,opt,name=vessel" json:"vessel,omitempty"`
	Vessels              []*Vessel `protobuf:"bytes,2,rep,name=vessels" json:"vessels,omitempty"`
	Created              bool      `protobuf:"varint,3,opt,name=created" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_vessel_59b3339503f4b360, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func init() {
	proto.RegisterType((*Vessel)(nil), "vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "vessel.Specification")
	proto.RegisterType((*Response)(nil), "vessel.Response")
}

func init() { proto.RegisterFile("proto/vessel/vessel.proto", fileDescriptor_vessel_59b3339503f4b360) }

var fileDescriptor_vessel_59b3339503f4b360 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4f, 0x4b, 0xfb, 0x40,
	0x10, 0xfd, 0x6d, 0xda, 0xa6, 0xe9, 0xfc, 0x68, 0x91, 0x01, 0x61, 0x5b, 0x14, 0x42, 0x0e, 0x92,
	0x83, 0x54, 0xa8, 0x37, 0x6f, 0x22, 0x08, 0x7a, 0xdc, 0x82, 0x1e, 0xcb, 0x76, 0x77, 0xd4, 0x85,
	0xe6, 0x0f, 0x49, 0x48, 0xdb, 0x6f, 0xe3, 0x47, 0x15, 0x26, 0x49, 0xa5, 0xa5, 0x78, 0xda, 0x79,
	0x6f, 0xde, 0x3e, 0xde, 0xbe, 0x85, 0x69, 0x5e, 0x64, 0x55, 0x76, 0x57, 0x53, 0x59, 0xd2, 0xa6,
	0x3d, 0xe6, 0xcc, 0xa1, 0xdf, 0xa0, 0xe8, 0x5b, 0x80, 0xff, 0xc6, 0x23, 0x4e, 0xc0, 0x73, 0x56,
	0x8a, 0x50, 0xc4, 0x23, 0xe5, 0x39, 0x8b, 0x33, 0x08, 0x8c, 0xce, 0xb5, 0x71, 0xd5, 0x5e, 0x7a,
	0xa1, 0x88, 0x07, 0xea, 0x80, 0xf1, 0x1a, 0x20, 0xd1, 0xbb, 0xd5, 0x96, 0xdc, 0xe7, 0x57, 0x25,
	0x7b, 0xbc, 0x1d, 0x25, 0x7a, 0xf7, 0xce, 0x04, 0x22, 0xf4, 0x53, 0x9d, 0x90, 0xec, 0xb3, 0x19,
	0xcf, 0x78, 0x05, 0x23, 0x5d, 0x6b, 0xb7, 0xd1, 0xeb, 0x0d, 0xc9, 0x41, 0x28, 0xe2, 0x40, 0xfd,
	0x12, 0x38, 0x85, 0x20, 0xdb, 0xa6, 0x54, 0xac, 0x9c, 0x95, 0x3e, 0xdf, 0x1a, 0x32, 0x7e, 0xb1,
	0xd1, 0x2b, 0x8c, 0x97, 0x39, 0x19, 0xf7, 0xe1, 0x8c, 0xae, 0x5c, 0x96, 0x1e, 0x05, 0x13, 0x7f,
	0x06, 0xf3, 0x4e, 0x82, 0x45, 0x35, 0x04, 0x8a, 0xca, 0x3c, 0x4b, 0x4b, 0xc2, 0x1b, 0x68, 0x4b,
	0x60, 0x93, 0xff, 0x8b, 0xc9, 0xbc, 0x6d, 0xa8, 0xe9, 0x43, 0xb5, 0x5b, 0x8c, 0x61, 0xd8, 0x4c,
	0xa5, 0xf4, 0xc2, 0xde, 0x19, 0x61, 0xb7, 0x46, 0x09, 0x43, 0x53, 0x90, 0xae, 0xc8, 0x72, 0x25,
	0x81, 0xea, 0xe0, 0x62, 0x0f, 0xe3, 0x46, 0xbc, 0xa4, 0xa2, 0x76, 0x86, 0xf0, 0x01, 0xc6, 0xcf,
	0x2e, 0xb5, 0x8f, 0x87, 0x02, 0x2e, 0x3b, 0xd3, 0xa3, 0xb7, 0xce, 0x2e, 0x3a, 0xba, 0x8b, 0x1d,
	0xfd, 0xc3, 0x5b, 0xf0, 0x9f, 0xd8, 0x17, 0x4f, 0x92, 0x9c, 0x53, 0xaf, 0x7d, 0xfe, 0xf0, 0xfb,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xdc, 0xb4, 0x34, 0x0d, 0x02, 0x00, 0x00,
}
