// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/service/match/service.proto

package match

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	types "github.com/kubegames/kubegames-hall/app/service/match/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() { proto.RegisterFile("app/service/match/service.proto", fileDescriptor_791ab8538e00c846) }

var fileDescriptor_791ab8538e00c846 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x2c, 0x28, 0xd0,
	0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0xcf, 0x4d, 0x2c, 0x49, 0xce, 0x80, 0xf1, 0xf4,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x78, 0xa1, 0xdc, 0x78, 0xb0, 0xa4, 0x94, 0x32, 0xa6, 0xfa,
	0x92, 0xca, 0x82, 0xd4, 0x62, 0x08, 0x09, 0xd1, 0x23, 0xa5, 0x02, 0x52, 0x94, 0x93, 0x99, 0xa4,
	0x9f, 0x9e, 0x9f, 0x9f, 0x9e, 0x93, 0xaa, 0x9f, 0x58, 0x90, 0xa9, 0x9f, 0x98, 0x97, 0x97, 0x5f,
	0x92, 0x58, 0x92, 0x99, 0x9f, 0x07, 0x53, 0xa5, 0x80, 0x50, 0x95, 0x9e, 0xaf, 0x0f, 0x16, 0x4b,
	0x2a, 0x4d, 0x03, 0xf3, 0x20, 0x2a, 0x8c, 0x0e, 0x31, 0x72, 0xf1, 0xf8, 0x82, 0xec, 0x08, 0x86,
	0x58, 0x28, 0x14, 0xc5, 0xc5, 0x0a, 0xe6, 0x0b, 0x49, 0xea, 0x81, 0xed, 0x8e, 0x87, 0xd8, 0x0a,
	0x16, 0x0b, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x91, 0x92, 0xc2, 0x26, 0x55, 0x5c, 0x90, 0x9f,
	0x57, 0x9c, 0xaa, 0x24, 0xd5, 0x74, 0xf9, 0xc9, 0x64, 0x26, 0x11, 0x25, 0x7e, 0xa8, 0xd3, 0xcb,
	0x0c, 0x21, 0x0c, 0x2b, 0x46, 0x2d, 0xa1, 0x30, 0x2e, 0x96, 0x80, 0xcc, 0xbc, 0x74, 0x21, 0x09,
	0x14, 0xfd, 0x20, 0x21, 0x98, 0xc9, 0x92, 0x58, 0x64, 0xa0, 0x06, 0x4b, 0x82, 0x0d, 0x16, 0x56,
	0xe2, 0x43, 0x18, 0x5c, 0x90, 0x99, 0x97, 0x6e, 0xc5, 0xa8, 0xe5, 0x14, 0x7e, 0xe2, 0xa1, 0x1c,
	0xc3, 0x83, 0x87, 0x72, 0x8c, 0x1f, 0x1e, 0xca, 0x31, 0xae, 0x78, 0x24, 0xc7, 0x78, 0xe2, 0x91,
	0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xce, 0x78, 0x2c, 0xc7, 0x10, 0x65,
	0x9a, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x5d, 0x9a, 0x94, 0x9a,
	0x9e, 0x98, 0x9b, 0x5a, 0x8c, 0x60, 0xe9, 0x66, 0x24, 0xe6, 0xe4, 0xe8, 0x63, 0x84, 0x7b, 0x12,
	0x1b, 0x38, 0x90, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x97, 0x23, 0xd0, 0xc3, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MatchServiceClient is the client API for MatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MatchServiceClient interface {
	//match player
	Match(ctx context.Context, in *types.MatchRequest, opts ...grpc.CallOption) (*types.MatchResponse, error)
	//Ping
	Ping(ctx context.Context, in *types.PingRequest, opts ...grpc.CallOption) (*types.PingResponse, error)
}

type matchServiceClient struct {
	cc *grpc.ClientConn
}

func NewMatchServiceClient(cc *grpc.ClientConn) MatchServiceClient {
	return &matchServiceClient{cc}
}

func (c *matchServiceClient) Match(ctx context.Context, in *types.MatchRequest, opts ...grpc.CallOption) (*types.MatchResponse, error) {
	out := new(types.MatchResponse)
	err := c.cc.Invoke(ctx, "/service_match.MatchService/Match", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchServiceClient) Ping(ctx context.Context, in *types.PingRequest, opts ...grpc.CallOption) (*types.PingResponse, error) {
	out := new(types.PingResponse)
	err := c.cc.Invoke(ctx, "/service_match.MatchService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatchServiceServer is the server API for MatchService service.
type MatchServiceServer interface {
	//match player
	Match(context.Context, *types.MatchRequest) (*types.MatchResponse, error)
	//Ping
	Ping(context.Context, *types.PingRequest) (*types.PingResponse, error)
}

// UnimplementedMatchServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMatchServiceServer struct {
}

func (*UnimplementedMatchServiceServer) Match(ctx context.Context, req *types.MatchRequest) (*types.MatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Match not implemented")
}
func (*UnimplementedMatchServiceServer) Ping(ctx context.Context, req *types.PingRequest) (*types.PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterMatchServiceServer(s *grpc.Server, srv MatchServiceServer) {
	s.RegisterService(&_MatchService_serviceDesc, srv)
}

func _MatchService_Match_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.MatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchServiceServer).Match(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_match.MatchService/Match",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchServiceServer).Match(ctx, req.(*types.MatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_match.MatchService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchServiceServer).Ping(ctx, req.(*types.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MatchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service_match.MatchService",
	HandlerType: (*MatchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Match",
			Handler:    _MatchService_Match_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _MatchService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/service/match/service.proto",
}
