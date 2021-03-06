// Code generated by protoc-gen-go.
// source: v2ray.com/core/common/protocol/server_spec.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	v2ray.com/core/common/protocol/server_spec.proto
	v2ray.com/core/common/protocol/user.proto

It has these top-level messages:
	ServerEndpoint
	User
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_common_net "v2ray.com/core/common/net"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ServerEndpoint struct {
	Address *v2ray_core_common_net.IPOrDomain `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Port    uint32                            `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	User    []*User                           `protobuf:"bytes,3,rep,name=user" json:"user,omitempty"`
}

func (m *ServerEndpoint) Reset()                    { *m = ServerEndpoint{} }
func (m *ServerEndpoint) String() string            { return proto.CompactTextString(m) }
func (*ServerEndpoint) ProtoMessage()               {}
func (*ServerEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServerEndpoint) GetAddress() *v2ray_core_common_net.IPOrDomain {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ServerEndpoint) GetUser() []*User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*ServerEndpoint)(nil), "v2ray.core.common.protocol.ServerEndpoint")
}

func init() { proto.RegisterFile("v2ray.com/core/common/protocol/server_spec.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x8f, 0x3d, 0x4b, 0x04, 0x31,
	0x10, 0x86, 0x59, 0xef, 0x50, 0xc9, 0xa1, 0x42, 0xaa, 0x65, 0x0b, 0x89, 0x36, 0xae, 0xcd, 0x44,
	0x56, 0xbb, 0xeb, 0x0e, 0x2d, 0xac, 0x3c, 0xf6, 0xb0, 0xb1, 0x91, 0x35, 0x3b, 0xc5, 0x81, 0xc9,
	0x84, 0x49, 0x3c, 0xf0, 0x97, 0xf8, 0x77, 0x65, 0x13, 0x53, 0xf9, 0xd5, 0x0d, 0x6f, 0x9e, 0xf7,
	0x23, 0xe2, 0x6a, 0xd7, 0xf1, 0xf0, 0x0e, 0x86, 0xac, 0x36, 0xc4, 0xa8, 0x0d, 0x59, 0x4b, 0x4e,
	0x7b, 0xa6, 0x48, 0x86, 0x5e, 0x75, 0x40, 0xde, 0x21, 0x3f, 0x07, 0x8f, 0x06, 0x92, 0x28, 0x9b,
	0xe2, 0x60, 0x84, 0x4c, 0x43, 0xa1, 0x9b, 0x8b, 0x9f, 0xd3, 0x1c, 0x46, 0x3d, 0x8c, 0x23, 0x63,
	0x08, 0x99, 0x6d, 0x2e, 0xff, 0xa9, 0x7d, 0x0b, 0xc8, 0x19, 0x3d, 0xff, 0xa8, 0xc4, 0xf1, 0x26,
	0xad, 0xb8, 0x73, 0xa3, 0xa7, 0xad, 0x8b, 0x72, 0x29, 0x0e, 0xbe, 0xe2, 0xea, 0x4a, 0x55, 0xed,
	0xa2, 0x3b, 0x83, 0xef, 0xa3, 0x1c, 0x46, 0xb8, 0x5f, 0x3f, 0xf0, 0x2d, 0xd9, 0x61, 0xeb, 0xfa,
	0xe2, 0x90, 0x52, 0xcc, 0x3d, 0x71, 0xac, 0xf7, 0x54, 0xd5, 0x1e, 0xf5, 0xe9, 0x96, 0x37, 0x62,
	0x3e, 0x35, 0xd6, 0x33, 0x35, 0x6b, 0x17, 0x9d, 0x82, 0xdf, 0xbf, 0x08, 0x8f, 0x01, 0xb9, 0x4f,
	0xf4, 0x6a, 0x29, 0x4e, 0x0d, 0xd9, 0x3f, 0xe0, 0xd5, 0x49, 0x1e, 0xbe, 0xf1, 0x68, 0xd6, 0x93,
	0xf6, 0x74, 0x58, 0x9e, 0x5e, 0xf6, 0xd3, 0x75, 0xfd, 0x19, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x28,
	0xb3, 0x3a, 0x81, 0x01, 0x00, 0x00,
}
