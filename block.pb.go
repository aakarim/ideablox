// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/block.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	protobuf/block.proto

It has these top-level messages:
	Block
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

type Block struct {
	Timestamp     int64  `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Data          []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	PrevBlockHash []byte `protobuf:"bytes,3,opt,name=prevBlockHash,proto3" json:"prevBlockHash,omitempty"`
	Hash          []byte `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
	Nonce         int64  `protobuf:"varint,5,opt,name=nonce" json:"nonce,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Block) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Block) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Block) GetPrevBlockHash() []byte {
	if m != nil {
		return m.PrevBlockHash
	}
	return nil
}

func (m *Block) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Block) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func init() {
	proto.RegisterType((*Block)(nil), "main.Block")
}

func init() { proto.RegisterFile("protobuf/block.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0xca, 0xc9, 0x4f, 0xce, 0xd6, 0x03, 0x73, 0x85, 0x58, 0x72,
	0x13, 0x33, 0xf3, 0x94, 0x3a, 0x19, 0xb9, 0x58, 0x9d, 0x40, 0xa2, 0x42, 0x32, 0x5c, 0x9c, 0x25,
	0x99, 0xb9, 0xa9, 0xc5, 0x25, 0x89, 0xb9, 0x05, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x08,
	0x01, 0x21, 0x21, 0x2e, 0x96, 0x94, 0xc4, 0x92, 0x44, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x9e, 0x20,
	0x30, 0x5b, 0x48, 0x85, 0x8b, 0xb7, 0xa0, 0x28, 0xb5, 0x0c, 0xac, 0xdd, 0x23, 0xb1, 0x38, 0x43,
	0x82, 0x19, 0x2c, 0x89, 0x2a, 0x08, 0xd2, 0x99, 0x01, 0x92, 0x64, 0x81, 0xe8, 0x04, 0xb1, 0x85,
	0x44, 0xb8, 0x58, 0xf3, 0xf2, 0xf3, 0x92, 0x53, 0x25, 0x58, 0xc1, 0xf6, 0x40, 0x38, 0x49, 0x6c,
	0x60, 0x87, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x31, 0x9a, 0x78, 0xb0, 0x00, 0x00,
	0x00,
}
