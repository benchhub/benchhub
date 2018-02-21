// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
*/
package grpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import agent "github.com/benchhub/benchhub/pkg/agent/agentpb"

import context "golang.org/x/net/context"
import grpc1 "google.golang.org/grpc"

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
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for BenchHubAgent service

type BenchHubAgentClient interface {
	Ping(ctx context.Context, in *agent.Ping, opts ...grpc1.CallOption) (*agent.Pong, error)
}

type benchHubAgentClient struct {
	cc *grpc1.ClientConn
}

func NewBenchHubAgentClient(cc *grpc1.ClientConn) BenchHubAgentClient {
	return &benchHubAgentClient{cc}
}

func (c *benchHubAgentClient) Ping(ctx context.Context, in *agent.Ping, opts ...grpc1.CallOption) (*agent.Pong, error) {
	out := new(agent.Pong)
	err := grpc1.Invoke(ctx, "/benchubcentralrpc.BenchHubAgent/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BenchHubAgent service

type BenchHubAgentServer interface {
	Ping(context.Context, *agent.Ping) (*agent.Pong, error)
}

func RegisterBenchHubAgentServer(s *grpc1.Server, srv BenchHubAgentServer) {
	s.RegisterService(&_BenchHubAgent_serviceDesc, srv)
}

func _BenchHubAgent_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(agent.Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BenchHubAgentServer).Ping(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/benchubcentralrpc.BenchHubAgent/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BenchHubAgentServer).Ping(ctx, req.(*agent.Ping))
	}
	return interceptor(ctx, in, info, handler)
}

var _BenchHubAgent_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "benchubcentralrpc.BenchHubAgent",
	HandlerType: (*BenchHubAgentServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _BenchHubAgent_Ping_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptorRpc) }

var fileDescriptorRpc = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2a, 0x48, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4c, 0x4a, 0xcd, 0x4b, 0xce, 0x28, 0x4d, 0x4a, 0x4e,
	0xcd, 0x2b, 0x29, 0x4a, 0xcc, 0x29, 0x2a, 0x48, 0x96, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d,
	0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xab, 0x4c, 0x2a, 0x4d, 0x03,
	0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0x62, 0x82, 0x94, 0x15, 0x92, 0x72, 0xb0, 0x61, 0x19, 0xa5, 0x49,
	0x08, 0x46, 0x41, 0x76, 0xba, 0x7e, 0x62, 0x7a, 0x6a, 0x5e, 0x09, 0x84, 0x2c, 0x48, 0xd2, 0x2f,
	0xa9, 0x2c, 0x48, 0x2d, 0x86, 0xe8, 0x35, 0x32, 0xe6, 0xe2, 0x75, 0x02, 0xa9, 0xf4, 0x28, 0x4d,
	0x72, 0x04, 0x49, 0x0b, 0x29, 0x71, 0xb1, 0x04, 0x64, 0xe6, 0xa5, 0x0b, 0x71, 0xeb, 0x81, 0x95,
	0xeb, 0x81, 0x38, 0x52, 0x70, 0x4e, 0x7e, 0x5e, 0xba, 0x12, 0x83, 0x93, 0xd8, 0x89, 0x87, 0x72,
	0x0c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0x63, 0x14, 0x4b,
	0x7a, 0x51, 0x41, 0x72, 0x12, 0x1b, 0xd8, 0x4c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x52,
	0xa8, 0xc4, 0x2f, 0xde, 0x00, 0x00, 0x00,
}
