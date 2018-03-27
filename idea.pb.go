// Code generated by protoc-gen-go. DO NOT EDIT.
// source: idea.proto

/*
package main is a generated protocol buffer package.

It is generated from these files:
	idea.proto

It has these top-level messages:
	Idea
*/
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

type Idea struct {
	Text   string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=author" json:"author,omitempty"`
}

func (m *Idea) Reset()                    { *m = Idea{} }
func (m *Idea) String() string            { return proto.CompactTextString(m) }
func (*Idea) ProtoMessage()               {}
func (*Idea) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Idea) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Idea) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func init() {
	proto.RegisterType((*Idea)(nil), "ideablox.Idea")
}

func init() { proto.RegisterFile("idea.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 89 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x4c, 0x49, 0x4d,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0xb1, 0x93, 0x72, 0xf2, 0x2b, 0x94, 0x8c,
	0xb8, 0x58, 0x3c, 0x53, 0x52, 0x13, 0x85, 0x84, 0xb8, 0x58, 0x4a, 0x52, 0x2b, 0x4a, 0x24, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x31, 0x2e, 0xb6, 0xc4, 0xd2, 0x92, 0x8c, 0xfc,
	0x22, 0x09, 0x26, 0xb0, 0x28, 0x94, 0x97, 0xc4, 0x06, 0x36, 0xc4, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x2f, 0xff, 0xb9, 0x29, 0x52, 0x00, 0x00, 0x00,
}