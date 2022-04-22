// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/service/game/service.proto

package game

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	types "github.com/kubegames/kubegames-hall/app/service/game/types"
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

func init() { proto.RegisterFile("app/service/game/service.proto", fileDescriptor_ac20ec3a588897f2) }

var fileDescriptor_ac20ec3a588897f2 = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcd, 0x6e, 0x13, 0x31,
	0x14, 0x85, 0x3b, 0x29, 0xa0, 0xc6, 0x94, 0x88, 0xba, 0x52, 0x4b, 0x43, 0x19, 0x46, 0x11, 0x12,
	0x52, 0x04, 0x31, 0x0d, 0x3f, 0x42, 0x15, 0x42, 0x4a, 0x41, 0xa0, 0x2e, 0xba, 0xa1, 0xac, 0xba,
	0x20, 0xf5, 0xcc, 0xdc, 0x4c, 0xac, 0x26, 0xb6, 0x19, 0x3b, 0x81, 0x82, 0xd8, 0xb0, 0x42, 0x2c,
	0x61, 0xc3, 0x23, 0xb0, 0xe2, 0x39, 0xba, 0x44, 0xe2, 0x05, 0x9a, 0xc0, 0x03, 0xf0, 0x08, 0xc8,
	0x9e, 0x9f, 0x24, 0xcd, 0xb0, 0xb1, 0x73, 0xcf, 0x39, 0xbe, 0x9f, 0x6f, 0x34, 0x46, 0x2e, 0x95,
	0x92, 0x28, 0x88, 0x87, 0x2c, 0x00, 0x12, 0xd1, 0x3e, 0x64, 0x45, 0x43, 0xc6, 0x42, 0x0b, 0xbc,
	0x6c, 0xb4, 0x76, 0xaa, 0x55, 0x6b, 0x73, 0x69, 0x7d, 0x2c, 0x41, 0x25, 0x6b, 0x72, 0xa2, 0x7a,
	0xc3, 0x64, 0x7a, 0xcc, 0x27, 0x91, 0x10, 0x51, 0x0f, 0x08, 0x95, 0x8c, 0x50, 0xce, 0x85, 0xa6,
	0x9a, 0x09, 0x9e, 0xa5, 0xbc, 0x49, 0x2a, 0x12, 0xc4, 0x6a, 0xfe, 0xa0, 0x63, 0xab, 0x34, 0xf1,
	0x20, 0x4b, 0xd8, 0x32, 0x68, 0x47, 0xc0, 0xdb, 0x42, 0x02, 0xa7, 0x92, 0x0d, 0x9b, 0x44, 0x48,
	0xdb, 0x6d, 0xbe, 0x73, 0xf3, 0xf3, 0x39, 0x74, 0xf1, 0x39, 0xed, 0xc3, 0x7e, 0x72, 0x4d, 0xec,
	0xa3, 0x4a, 0x2b, 0x0c, 0x8d, 0xf2, 0x02, 0x5e, 0x0f, 0x40, 0x69, 0x5c, 0x6d, 0xd8, 0xa1, 0x92,
	0x4b, 0xcf, 0x7a, 0xd5, 0xab, 0x85, 0x9e, 0x92, 0x82, 0x2b, 0xa8, 0xad, 0x7f, 0xfc, 0xf5, 0xe7,
	0x6b, 0x69, 0xa5, 0xb6, 0x9c, 0xcc, 0x3e, 0xdc, 0x22, 0x34, 0x0c, 0xb7, 0x9d, 0x3a, 0x0e, 0x50,
	0x79, 0x5f, 0xd3, 0x58, 0x9b, 0x34, 0xde, 0x9c, 0x6e, 0x91, 0xcb, 0x19, 0xe0, 0xda, 0x7f, 0xdc,
	0x14, 0xb1, 0x61, 0x11, 0xab, 0xb5, 0x4a, 0x8e, 0x50, 0x26, 0x63, 0x20, 0x0c, 0x95, 0x9f, 0xf4,
	0x84, 0x82, 0x79, 0x48, 0x2e, 0x17, 0x42, 0xa6, 0xdc, 0x14, 0x72, 0xdd, 0x42, 0x36, 0xea, 0xeb,
	0x39, 0x24, 0x30, 0x19, 0xf2, 0xde, 0x94, 0xbb, 0x4f, 0x3f, 0x60, 0x40, 0x4b, 0xcf, 0x18, 0xb7,
	0xc3, 0xe3, 0x99, 0x7f, 0x24, 0x53, 0x33, 0xd0, 0x66, 0xb1, 0x99, 0x72, 0x5c, 0xcb, 0xb9, 0x82,
	0xd7, 0x72, 0x4e, 0x87, 0xf1, 0x70, 0x82, 0x39, 0x44, 0xe5, 0xec, 0x8c, 0xc2, 0x85, 0xad, 0x54,
	0xe1, 0x44, 0x53, 0x6e, 0x4a, 0x5a, 0xb3, 0xa4, 0xcb, 0xb8, 0x32, 0x43, 0x52, 0x3b, 0x9f, 0x4a,
	0x5f, 0x5a, 0x3f, 0x1c, 0x7c, 0x1f, 0xdd, 0x3c, 0x1a, 0xf8, 0x60, 0x3c, 0xe5, 0x99, 0x35, 0x59,
	0xa8, 0x64, 0x9e, 0x7a, 0x43, 0xa3, 0x08, 0xe2, 0x6c, 0xaf, 0x9f, 0x47, 0x8b, 0x7b, 0xbb, 0x2f,
	0x9b, 0x8b, 0x5b, 0x8d, 0x3b, 0x07, 0x87, 0xe8, 0x15, 0xba, 0xd4, 0x1a, 0xe8, 0xae, 0x88, 0xd9,
	0x3b, 0xfb, 0x91, 0xe1, 0xbd, 0xa5, 0x12, 0x7e, 0x68, 0x24, 0xe0, 0x9a, 0x05, 0x56, 0xf3, 0xb4,
	0x38, 0x02, 0x7e, 0xcb, 0x93, 0x31, 0x74, 0xd8, 0x5b, 0x08, 0x3d, 0xff, 0xd8, 0xdb, 0x01, 0x1a,
	0x43, 0xbc, 0x9d, 0xee, 0xde, 0x23, 0x1b, 0x79, 0x5c, 0x9d, 0x6d, 0xe6, 0x95, 0xfc, 0x55, 0xb4,
	0x72, 0x96, 0xb0, 0x70, 0x32, 0x72, 0x17, 0x4e, 0x47, 0xae, 0xf3, 0x77, 0xe4, 0x3a, 0xdf, 0xc7,
	0xae, 0x73, 0x32, 0x76, 0x9d, 0x9f, 0x63, 0xd7, 0x39, 0x1d, 0xbb, 0xce, 0xb7, 0xdf, 0xee, 0xc2,
	0xc1, 0xbd, 0x88, 0xe9, 0xee, 0xc0, 0x6f, 0x04, 0xa2, 0x4f, 0xf2, 0xb1, 0x26, 0xbf, 0x6e, 0x77,
	0x69, 0xaf, 0x47, 0xce, 0xbe, 0x54, 0xff, 0x82, 0x7d, 0x1e, 0x77, 0xff, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x1f, 0x90, 0xcd, 0x7e, 0xf2, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameServiceClient interface {
	//添加游戏
	AddGameRequest(ctx context.Context, in *types.AddGameRequest, opts ...grpc.CallOption) (*types.AddGameResponse, error)
	//启动游戏服务
	StartGame(ctx context.Context, in *types.StartGameRequest, opts ...grpc.CallOption) (*types.StartGameResponse, error)
	//关闭游戏服务
	CloseGame(ctx context.Context, in *types.CloseGameRequest, opts ...grpc.CallOption) (*types.CloseGameResponse, error)
	//查找游戏服务
	FindGame(ctx context.Context, in *types.FindGameRequest, opts ...grpc.CallOption) (*types.FindGameResponse, error)
	//查找游戏服务
	FindGames(ctx context.Context, in *types.FindGamesRequest, opts ...grpc.CallOption) (*types.FindGamesResponse, error)
}

type gameServiceClient struct {
	cc *grpc.ClientConn
}

func NewGameServiceClient(cc *grpc.ClientConn) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) AddGameRequest(ctx context.Context, in *types.AddGameRequest, opts ...grpc.CallOption) (*types.AddGameResponse, error) {
	out := new(types.AddGameResponse)
	err := c.cc.Invoke(ctx, "/game_service.GameService/AddGameRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) StartGame(ctx context.Context, in *types.StartGameRequest, opts ...grpc.CallOption) (*types.StartGameResponse, error) {
	out := new(types.StartGameResponse)
	err := c.cc.Invoke(ctx, "/game_service.GameService/StartGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) CloseGame(ctx context.Context, in *types.CloseGameRequest, opts ...grpc.CallOption) (*types.CloseGameResponse, error) {
	out := new(types.CloseGameResponse)
	err := c.cc.Invoke(ctx, "/game_service.GameService/CloseGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) FindGame(ctx context.Context, in *types.FindGameRequest, opts ...grpc.CallOption) (*types.FindGameResponse, error) {
	out := new(types.FindGameResponse)
	err := c.cc.Invoke(ctx, "/game_service.GameService/FindGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) FindGames(ctx context.Context, in *types.FindGamesRequest, opts ...grpc.CallOption) (*types.FindGamesResponse, error) {
	out := new(types.FindGamesResponse)
	err := c.cc.Invoke(ctx, "/game_service.GameService/FindGames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServiceServer is the server API for GameService service.
type GameServiceServer interface {
	//添加游戏
	AddGameRequest(context.Context, *types.AddGameRequest) (*types.AddGameResponse, error)
	//启动游戏服务
	StartGame(context.Context, *types.StartGameRequest) (*types.StartGameResponse, error)
	//关闭游戏服务
	CloseGame(context.Context, *types.CloseGameRequest) (*types.CloseGameResponse, error)
	//查找游戏服务
	FindGame(context.Context, *types.FindGameRequest) (*types.FindGameResponse, error)
	//查找游戏服务
	FindGames(context.Context, *types.FindGamesRequest) (*types.FindGamesResponse, error)
}

// UnimplementedGameServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGameServiceServer struct {
}

func (*UnimplementedGameServiceServer) AddGameRequest(ctx context.Context, req *types.AddGameRequest) (*types.AddGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGameRequest not implemented")
}
func (*UnimplementedGameServiceServer) StartGame(ctx context.Context, req *types.StartGameRequest) (*types.StartGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartGame not implemented")
}
func (*UnimplementedGameServiceServer) CloseGame(ctx context.Context, req *types.CloseGameRequest) (*types.CloseGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseGame not implemented")
}
func (*UnimplementedGameServiceServer) FindGame(ctx context.Context, req *types.FindGameRequest) (*types.FindGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindGame not implemented")
}
func (*UnimplementedGameServiceServer) FindGames(ctx context.Context, req *types.FindGamesRequest) (*types.FindGamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindGames not implemented")
}

func RegisterGameServiceServer(s *grpc.Server, srv GameServiceServer) {
	s.RegisterService(&_GameService_serviceDesc, srv)
}

func _GameService_AddGameRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.AddGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).AddGameRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_service.GameService/AddGameRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).AddGameRequest(ctx, req.(*types.AddGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_StartGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.StartGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).StartGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_service.GameService/StartGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).StartGame(ctx, req.(*types.StartGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_CloseGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.CloseGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).CloseGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_service.GameService/CloseGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).CloseGame(ctx, req.(*types.CloseGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_FindGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.FindGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).FindGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_service.GameService/FindGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).FindGame(ctx, req.(*types.FindGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_FindGames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.FindGamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).FindGames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_service.GameService/FindGames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).FindGames(ctx, req.(*types.FindGamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GameService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "game_service.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddGameRequest",
			Handler:    _GameService_AddGameRequest_Handler,
		},
		{
			MethodName: "StartGame",
			Handler:    _GameService_StartGame_Handler,
		},
		{
			MethodName: "CloseGame",
			Handler:    _GameService_CloseGame_Handler,
		},
		{
			MethodName: "FindGame",
			Handler:    _GameService_FindGame_Handler,
		},
		{
			MethodName: "FindGames",
			Handler:    _GameService_FindGames_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/service/game/service.proto",
}
