// Code generated by protoc-gen-gogo.
// source: api.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto
	ipam.proto

It has these top-level messages:
	IPAMRequest
	IPConfig
	DNS
	IPAMResponse
	IPReleaseResponse
	IPReleaseRequest
*/
package api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NetworkManager service

type NetworkManagerClient interface {
	AllocateIP(ctx context.Context, in *IPAMRequest, opts ...grpc.CallOption) (*IPAMResponse, error)
	ReleaseIP(ctx context.Context, in *IPReleaseRequest, opts ...grpc.CallOption) (*IPReleaseResponse, error)
}

type networkManagerClient struct {
	cc *grpc.ClientConn
}

func NewNetworkManagerClient(cc *grpc.ClientConn) NetworkManagerClient {
	return &networkManagerClient{cc}
}

func (c *networkManagerClient) AllocateIP(ctx context.Context, in *IPAMRequest, opts ...grpc.CallOption) (*IPAMResponse, error) {
	out := new(IPAMResponse)
	err := grpc.Invoke(ctx, "/api.NetworkManager/AllocateIP", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkManagerClient) ReleaseIP(ctx context.Context, in *IPReleaseRequest, opts ...grpc.CallOption) (*IPReleaseResponse, error) {
	out := new(IPReleaseResponse)
	err := grpc.Invoke(ctx, "/api.NetworkManager/ReleaseIP", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NetworkManager service

type NetworkManagerServer interface {
	AllocateIP(context.Context, *IPAMRequest) (*IPAMResponse, error)
	ReleaseIP(context.Context, *IPReleaseRequest) (*IPReleaseResponse, error)
}

func RegisterNetworkManagerServer(s *grpc.Server, srv NetworkManagerServer) {
	s.RegisterService(&_NetworkManager_serviceDesc, srv)
}

func _NetworkManager_AllocateIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPAMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkManagerServer).AllocateIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkManager/AllocateIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkManagerServer).AllocateIP(ctx, req.(*IPAMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkManager_ReleaseIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkManagerServer).ReleaseIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkManager/ReleaseIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkManagerServer).ReleaseIP(ctx, req.(*IPReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.NetworkManager",
	HandlerType: (*NetworkManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllocateIP",
			Handler:    _NetworkManager_AllocateIP_Handler,
		},
		{
			MethodName: "ReleaseIP",
			Handler:    _NetworkManager_ReleaseIP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptorApi) }

var fileDescriptorApi = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0xe2, 0xca, 0x2c, 0x48, 0xcc,
	0x85, 0x08, 0x48, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x99, 0xfa, 0x20, 0x16, 0x44, 0xd4, 0xa8,
	0x99, 0x91, 0x8b, 0xcf, 0x2f, 0xb5, 0xa4, 0x3c, 0xbf, 0x28, 0xdb, 0x37, 0x31, 0x2f, 0x31, 0x3d,
	0xb5, 0x48, 0xc8, 0x98, 0x8b, 0xcb, 0x31, 0x27, 0x27, 0x3f, 0x39, 0xb1, 0x24, 0xd5, 0x33, 0x40,
	0x48, 0x40, 0x0f, 0x64, 0xa6, 0x67, 0x80, 0xa3, 0x6f, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89,
	0x94, 0x20, 0x92, 0x48, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x12, 0x83, 0x90, 0x0d, 0x17, 0x67,
	0x50, 0x6a, 0x4e, 0x6a, 0x62, 0x31, 0x48, 0x8f, 0x28, 0x54, 0x05, 0x54, 0x04, 0xa6, 0x51, 0x0c,
	0x5d, 0x18, 0xa6, 0x3b, 0x89, 0x0d, 0xec, 0x18, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6d,
	0xdf, 0x11, 0x19, 0xc0, 0x00, 0x00, 0x00,
}
