// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/kafka_hystrix_client_error.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/kafka_hystrix_client_error.proto

It has these top-level messages:
	Errors
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Errors struct {
	Code   string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Entity string `protobuf:"bytes,2,opt,name=entity" json:"entity,omitempty"`
}

func (m *Errors) Reset()                    { *m = Errors{} }
func (m *Errors) String() string            { return proto1.CompactTextString(m) }
func (*Errors) ProtoMessage()               {}
func (*Errors) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Errors) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Errors) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func init() {
	proto1.RegisterType((*Errors)(nil), "proto.Errors")
}

func init() { proto1.RegisterFile("proto/kafka_hystrix_client_error.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x4e, 0x4c, 0xcb, 0x4e, 0x8c, 0xcf, 0xa8, 0x2c, 0x2e, 0x29, 0xca, 0xac, 0x88,
	0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x89, 0x4f, 0x2d, 0x2a, 0xca, 0x2f, 0xd2, 0x03, 0x2b, 0x10,
	0x62, 0x05, 0x53, 0x4a, 0x26, 0x5c, 0x6c, 0xae, 0x20, 0xd1, 0x62, 0x21, 0x21, 0x2e, 0x96, 0xe4,
	0xfc, 0x94, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x8c, 0x8b, 0x2d,
	0x35, 0xaf, 0x24, 0xb3, 0xa4, 0x52, 0x82, 0x09, 0x2c, 0x0a, 0xe5, 0x25, 0xb1, 0x81, 0x35, 0x1b,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x22, 0x4f, 0xa2, 0xc4, 0x6d, 0x00, 0x00, 0x00,
}