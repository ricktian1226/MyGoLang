// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

package main

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

type Test0 struct {
	A                *int64 `protobuf:"varint,1,req,name=a" json:"a,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Test0) Reset()                    { *m = Test0{} }
func (m *Test0) String() string            { return proto.CompactTextString(m) }
func (*Test0) ProtoMessage()               {}
func (*Test0) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Test0) GetA() int64 {
	if m != nil && m.A != nil {
		return *m.A
	}
	return 0
}

func init() {
	proto.RegisterType((*Test0)(nil), "main.Test0")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 61 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0x12, 0xe2, 0x62,
	0x0d, 0x01, 0x8a, 0x19, 0x08, 0x71, 0x72, 0x31, 0x26, 0x4a, 0x30, 0x2a, 0x30, 0x69, 0x30, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x9e, 0x9f, 0x52, 0x26, 0x00, 0x00, 0x00,
}
