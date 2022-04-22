// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/service/platform/service.proto

package platform

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	types "github.com/kubegames/kubegames-hall/app/service/platform/types"
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

func init() {
	proto.RegisterFile("app/service/platform/service.proto", fileDescriptor_7456b9fc0b88b8fb)
}

var fileDescriptor_7456b9fc0b88b8fb = []byte{
	// 752 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xdf, 0x4e, 0x13, 0x4b,
	0x18, 0xc0, 0xbb, 0x25, 0xe7, 0xe4, 0x30, 0xe4, 0xe4, 0x1c, 0xc7, 0x84, 0x68, 0x8d, 0xcb, 0xb2,
	0x60, 0x95, 0x46, 0x3a, 0x14, 0x13, 0x43, 0x88, 0x31, 0x29, 0x21, 0x92, 0x26, 0x92, 0x90, 0xa2,
	0x37, 0x5c, 0x88, 0xb3, 0xed, 0x74, 0xd9, 0xb0, 0xdd, 0x19, 0x66, 0x67, 0x2b, 0x95, 0x10, 0xa3,
	0x57, 0x5c, 0x63, 0x8c, 0x3e, 0x02, 0xf1, 0x49, 0xb8, 0x34, 0xf1, 0x05, 0xa0, 0xfa, 0x00, 0xbe,
	0x81, 0x66, 0x66, 0x77, 0xbb, 0xdd, 0xfe, 0x83, 0x9b, 0xed, 0xee, 0xf7, 0xfd, 0xbe, 0xf9, 0x7e,
	0x33, 0x5f, 0xdb, 0x05, 0x26, 0x66, 0x0c, 0xf9, 0x84, 0xb7, 0x9c, 0x1a, 0x41, 0xcc, 0xc5, 0xa2,
	0x41, 0x79, 0x33, 0x0e, 0x14, 0x19, 0xa7, 0x82, 0xc2, 0xff, 0xe3, 0xf8, 0x6e, 0x14, 0xcf, 0xe5,
	0x87, 0x56, 0x89, 0x36, 0x23, 0x7e, 0x78, 0x0d, 0x2b, 0x73, 0xf3, 0x92, 0x73, 0x1d, 0x0b, 0xd9,
	0x94, 0xda, 0x2e, 0x41, 0x98, 0x39, 0x08, 0x7b, 0x1e, 0x15, 0x58, 0x38, 0xd4, 0x8b, 0x29, 0x23,
	0xa1, 0x6c, 0x8a, 0x54, 0xcc, 0x0a, 0x1a, 0xea, 0x29, 0x22, 0x1e, 0xc7, 0x84, 0x7a, 0xac, 0xed,
	0xda, 0xc4, 0xdb, 0xa5, 0x8c, 0x78, 0x98, 0x39, 0xad, 0x65, 0x44, 0x99, 0x5a, 0x6d, 0x70, 0xe5,
	0xe5, 0xdf, 0x00, 0xfc, 0xb7, 0x15, 0xe9, 0x6d, 0x87, 0xba, 0xf0, 0x18, 0x80, 0x32, 0x63, 0x6e,
	0xbb, 0x4a, 0x2d, 0x2a, 0xe0, 0x6c, 0xb1, 0xbb, 0xb9, 0x50, 0x3c, 0xc9, 0x55, 0xc9, 0x41, 0x40,
	0x7c, 0x91, 0x33, 0xc7, 0x21, 0x3e, 0xa3, 0x9e, 0x4f, 0xcc, 0x07, 0xa7, 0xe5, 0xac, 0x95, 0xf9,
	0xf0, 0xfd, 0xe7, 0xc7, 0xec, 0x5d, 0xf3, 0x56, 0x72, 0x24, 0xad, 0x12, 0xc2, 0x92, 0x46, 0x5c,
	0xe2, 0xab, 0x5a, 0x01, 0x7e, 0xd2, 0x94, 0x52, 0x9b, 0xf0, 0xe7, 0x04, 0xb7, 0xc8, 0x06, 0x6e,
	0x12, 0x98, 0xef, 0xef, 0xd0, 0x07, 0xc4, 0x26, 0xf7, 0xaf, 0xe4, 0x22, 0x9d, 0x52, 0xa2, 0x93,
	0x2f, 0xcc, 0xa7, 0x74, 0x98, 0x2a, 0x41, 0xae, 0xac, 0x41, 0x47, 0xe1, 0x53, 0x65, 0xfd, 0x18,
	0x1e, 0x82, 0xc9, 0x35, 0x4e, 0x71, 0xbd, 0x86, 0x7d, 0x01, 0x8d, 0xfe, 0x46, 0xdd, 0x54, 0xac,
	0x32, 0x3b, 0x86, 0x88, 0x24, 0xf2, 0x89, 0xc4, 0x1d, 0x73, 0x3a, 0x25, 0x61, 0xc5, 0xb0, 0x3c,
	0x91, 0x77, 0x60, 0xf2, 0x99, 0xe3, 0xd5, 0xab, 0x94, 0x36, 0xfd, 0xc1, 0xce, 0xdd, 0xd4, 0xc8,
	0xce, 0x3d, 0x44, 0xd4, 0x79, 0x31, 0xe9, 0x6c, 0x42, 0x23, 0xd5, 0xb9, 0xe1, 0x78, 0x75, 0xc4,
	0x25, 0x8d, 0x8e, 0x6c, 0xdc, 0x24, 0x72, 0xeb, 0xef, 0x35, 0x30, 0x55, 0x0d, 0x3c, 0xcf, 0xf1,
	0x6c, 0xb9, 0x0e, 0x1c, 0x18, 0x78, 0x4f, 0x32, 0xb6, 0x98, 0x1b, 0xcb, 0x44, 0x1e, 0x0b, 0x89,
	0x87, 0x6e, 0xde, 0x4e, 0x79, 0xf0, 0x10, 0x57, 0x2a, 0xf2, 0x10, 0xa4, 0x43, 0xb9, 0x76, 0x10,
	0x38, 0x9c, 0x0c, 0x77, 0xe8, 0x49, 0x8e, 0x74, 0x48, 0x31, 0x57, 0x3b, 0xe0, 0x10, 0x4f, 0x39,
	0x54, 0x89, 0x4b, 0xb0, 0x3f, 0xc2, 0xa1, 0x27, 0x39, 0xfa, 0x1c, 0x7a, 0x99, 0x21, 0x0e, 0xb9,
	0xbe, 0x73, 0x08, 0xf1, 0xae, 0xc3, 0x67, 0x0d, 0xc0, 0x97, 0xcc, 0xa5, 0xb8, 0x1e, 0x7e, 0xb7,
	0xab, 0xa4, 0x46, 0x79, 0x1d, 0x2e, 0xf4, 0xb7, 0x19, 0x64, 0x62, 0xa3, 0xc2, 0x75, 0xd0, 0x48,
	0xac, 0x90, 0x88, 0xcd, 0x98, 0xb9, 0x94, 0x58, 0xa0, 0xaa, 0x10, 0x57, 0x05, 0xd2, 0xec, 0x4c,
	0x03, 0xd3, 0xbd, 0x4b, 0x6d, 0x13, 0x21, 0x5c, 0x52, 0xf1, 0x1a, 0x14, 0x2e, 0x8e, 0x6b, 0x99,
	0x70, 0xb1, 0x61, 0xf1, 0xba, 0x78, 0x64, 0x89, 0x12, 0xcb, 0x79, 0x73, 0x66, 0x98, 0xa5, 0xaf,
	0x8a, 0x90, 0xe3, 0x35, 0xa8, 0x54, 0x3d, 0xd1, 0xc0, 0xd4, 0x06, 0x11, 0x72, 0x06, 0x5b, 0x94,
	0xba, 0x83, 0x83, 0xec, 0x49, 0x8e, 0x1c, 0x64, 0x8a, 0x89, 0x4c, 0x96, 0x12, 0x93, 0x7b, 0x70,
	0x2e, 0x65, 0x62, 0x13, 0xa1, 0x86, 0x88, 0x18, 0xa5, 0x2e, 0x3a, 0x92, 0xb7, 0x95, 0xf5, 0xe3,
	0xb5, 0x93, 0xec, 0x69, 0xf9, 0xab, 0x06, 0x4b, 0x60, 0x6e, 0x3f, 0xb0, 0x88, 0xfc, 0xc5, 0xf9,
	0x86, 0xbc, 0x1a, 0x71, 0xad, 0x81, 0x99, 0x63, 0xf8, 0x6f, 0xb0, 0x6d, 0x13, 0x5e, 0xf8, 0x0b,
	0x4c, 0x6c, 0x56, 0x5e, 0x2c, 0x4f, 0x94, 0x8a, 0x4b, 0x3b, 0xaf, 0xc1, 0x2b, 0xf0, 0x6f, 0x39,
	0x10, 0x7b, 0x94, 0x3b, 0x6f, 0xd5, 0xbf, 0x3a, 0xdc, 0xfc, 0x27, 0x0b, 0x57, 0x64, 0x88, 0x78,
	0xc2, 0xa9, 0xa9, 0x98, 0x21, 0xe8, 0x3e, 0xf1, 0x1e, 0x1a, 0x8c, 0x93, 0x86, 0x73, 0x48, 0xea,
	0x86, 0xd5, 0x36, 0xd6, 0x08, 0xe6, 0x84, 0xaf, 0x46, 0x9f, 0xc6, 0x13, 0x85, 0x3c, 0xcd, 0xa5,
	0x17, 0x33, 0xb2, 0xd6, 0x4d, 0x70, 0xa3, 0xbf, 0x43, 0xe6, 0xfc, 0x52, 0xcf, 0x5c, 0x5c, 0xea,
	0xda, 0xaf, 0x4b, 0x5d, 0x3b, 0xeb, 0xe8, 0xda, 0x79, 0x47, 0xd7, 0xbe, 0x75, 0x74, 0xed, 0xa2,
	0xa3, 0x6b, 0x5f, 0x7e, 0xe8, 0x99, 0x9d, 0x15, 0xdb, 0x11, 0x7b, 0x81, 0x55, 0xac, 0xd1, 0x26,
	0xea, 0x6e, 0x29, 0xb9, 0x5b, 0xdc, 0xc3, 0xae, 0x8b, 0x86, 0xbd, 0x1e, 0xad, 0xbf, 0xd5, 0x3b,
	0xe9, 0xd1, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x73, 0xa4, 0x5d, 0x20, 0x73, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PlatformServiceClient is the client API for PlatformService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlatformServiceClient interface {
	//申请机器人
	ApplyRobot(ctx context.Context, in *types.ApplyRobotRequest, opts ...grpc.CallOption) (*types.ApplyRobotResponse, error)
	//玩家离开游戏
	PlayerLeaveGame(ctx context.Context, in *types.PlayerLeaveGameRequest, opts ...grpc.CallOption) (*types.PlayerLeaveGameResponse, error)
	// 广播
	Broadcast(ctx context.Context, in *types.BroadcastRequest, opts ...grpc.CallOption) (*types.BroadcastResponse, error)
	//获取房间列表
	FindRooms(ctx context.Context, in *types.FindRoomsRequest, opts ...grpc.CallOption) (*types.FindRoomsResponse, error)
	//运行房间
	RunningRoom(ctx context.Context, in *types.RunningRoomRequest, opts ...grpc.CallOption) (*types.RunningRoomResponse, error)
	//抢占房间
	AcquireRoom(ctx context.Context, in *types.AcquireRoomRequest, opts ...grpc.CallOption) (*types.AcquireRoomResponse, error)
	//释放房间
	ReleaseRoom(ctx context.Context, in *types.ReleaseRoomRequest, opts ...grpc.CallOption) (*types.ReleaseRoomResponse, error)
	//战绩上传
	UploadPlayerRecord(ctx context.Context, in *types.UploadPlayerRecordRequest, opts ...grpc.CallOption) (*types.UploadPlayerRecordResponse, error)
	//上传用户结算信息
	UploadPlayerSettleInfo(ctx context.Context, in *types.UploadPlayerSettleInfoRequest, opts ...grpc.CallOption) (*types.UploadPlayerSettleInfoResponse, error)
	//获取房间血池值
	GetRoomPool(ctx context.Context, in *types.GetRoomPoolRequest, opts ...grpc.CallOption) (*types.GetRoomPoolResponse, error)
}

type platformServiceClient struct {
	cc *grpc.ClientConn
}

func NewPlatformServiceClient(cc *grpc.ClientConn) PlatformServiceClient {
	return &platformServiceClient{cc}
}

func (c *platformServiceClient) ApplyRobot(ctx context.Context, in *types.ApplyRobotRequest, opts ...grpc.CallOption) (*types.ApplyRobotResponse, error) {
	out := new(types.ApplyRobotResponse)
	err := c.cc.Invoke(ctx, "/platform_service.PlatformService/ApplyRobot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformServiceClient) PlayerLeaveGame(ctx context.Context, in *types.PlayerLeaveGameRequest, opts ...grpc.CallOption) (*types.PlayerLeaveGameResponse, error) {
	out := new(types.PlayerLeaveGameResponse)
	err := c.cc.Invoke(ctx, "/platform_service.PlatformService/PlayerLeaveGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformServiceClient) Broadcast(ctx context.Context, in *types.BroadcastRequest, opts ...grpc.CallOption) (*types.BroadcastResponse, error) {
	out := new(types.BroadcastResponse)
	err := c.cc.Invoke(ctx, "/platform_service.PlatformService/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformServiceClient) FindRooms(ctx context.Context, in *types.FindRoomsRequest, opts ...grpc.CallOption) (*types.FindRoomsResponse, error) {
	out := new(types.FindRoomsResponse)
	err := c.cc.Invoke(ctx, "/platform_service.PlatformService/FindRooms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformServiceClient) RunningRoom(ctx context.Context, in *types.RunningRoomRequest, opts ...grpc.CallOption) (*types.RunningRoomResponse, error) {
	out := new(types.RunningRoomResponse)
	err := c.cc.Invoke(ctx, "/platform_service.PlatformService/RunningRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformServiceClient) AcquireRoom(ctx context.Context, in *types.AcquireRoomRequest, opts ...grpc.CallOption) (*types.AcquireRoomResponse, error) {
	out := new(types.AcquireRoomResponse)
	err := c.cc.Invoke(ctx, "/platform_service.PlatformService/AcquireRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformServiceClient) ReleaseRoom(ctx context.Context, in *types.ReleaseRoomRequest, opts ...grpc.CallOption) (*types.ReleaseRoomResponse, error) {
	out := new(types.ReleaseRoomResponse)
	err := c.cc.Invoke(ctx, "/platform_service.PlatformService/ReleaseRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformServiceClient) UploadPlayerRecord(ctx context.Context, in *types.UploadPlayerRecordRequest, opts ...grpc.CallOption) (*types.UploadPlayerRecordResponse, error) {
	out := new(types.UploadPlayerRecordResponse)
	err := c.cc.Invoke(ctx, "/platform_service.PlatformService/UploadPlayerRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformServiceClient) UploadPlayerSettleInfo(ctx context.Context, in *types.UploadPlayerSettleInfoRequest, opts ...grpc.CallOption) (*types.UploadPlayerSettleInfoResponse, error) {
	out := new(types.UploadPlayerSettleInfoResponse)
	err := c.cc.Invoke(ctx, "/platform_service.PlatformService/UploadPlayerSettleInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformServiceClient) GetRoomPool(ctx context.Context, in *types.GetRoomPoolRequest, opts ...grpc.CallOption) (*types.GetRoomPoolResponse, error) {
	out := new(types.GetRoomPoolResponse)
	err := c.cc.Invoke(ctx, "/platform_service.PlatformService/GetRoomPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlatformServiceServer is the server API for PlatformService service.
type PlatformServiceServer interface {
	//申请机器人
	ApplyRobot(context.Context, *types.ApplyRobotRequest) (*types.ApplyRobotResponse, error)
	//玩家离开游戏
	PlayerLeaveGame(context.Context, *types.PlayerLeaveGameRequest) (*types.PlayerLeaveGameResponse, error)
	// 广播
	Broadcast(context.Context, *types.BroadcastRequest) (*types.BroadcastResponse, error)
	//获取房间列表
	FindRooms(context.Context, *types.FindRoomsRequest) (*types.FindRoomsResponse, error)
	//运行房间
	RunningRoom(context.Context, *types.RunningRoomRequest) (*types.RunningRoomResponse, error)
	//抢占房间
	AcquireRoom(context.Context, *types.AcquireRoomRequest) (*types.AcquireRoomResponse, error)
	//释放房间
	ReleaseRoom(context.Context, *types.ReleaseRoomRequest) (*types.ReleaseRoomResponse, error)
	//战绩上传
	UploadPlayerRecord(context.Context, *types.UploadPlayerRecordRequest) (*types.UploadPlayerRecordResponse, error)
	//上传用户结算信息
	UploadPlayerSettleInfo(context.Context, *types.UploadPlayerSettleInfoRequest) (*types.UploadPlayerSettleInfoResponse, error)
	//获取房间血池值
	GetRoomPool(context.Context, *types.GetRoomPoolRequest) (*types.GetRoomPoolResponse, error)
}

// UnimplementedPlatformServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPlatformServiceServer struct {
}

func (*UnimplementedPlatformServiceServer) ApplyRobot(ctx context.Context, req *types.ApplyRobotRequest) (*types.ApplyRobotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyRobot not implemented")
}
func (*UnimplementedPlatformServiceServer) PlayerLeaveGame(ctx context.Context, req *types.PlayerLeaveGameRequest) (*types.PlayerLeaveGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerLeaveGame not implemented")
}
func (*UnimplementedPlatformServiceServer) Broadcast(ctx context.Context, req *types.BroadcastRequest) (*types.BroadcastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (*UnimplementedPlatformServiceServer) FindRooms(ctx context.Context, req *types.FindRoomsRequest) (*types.FindRoomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindRooms not implemented")
}
func (*UnimplementedPlatformServiceServer) RunningRoom(ctx context.Context, req *types.RunningRoomRequest) (*types.RunningRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunningRoom not implemented")
}
func (*UnimplementedPlatformServiceServer) AcquireRoom(ctx context.Context, req *types.AcquireRoomRequest) (*types.AcquireRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcquireRoom not implemented")
}
func (*UnimplementedPlatformServiceServer) ReleaseRoom(ctx context.Context, req *types.ReleaseRoomRequest) (*types.ReleaseRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseRoom not implemented")
}
func (*UnimplementedPlatformServiceServer) UploadPlayerRecord(ctx context.Context, req *types.UploadPlayerRecordRequest) (*types.UploadPlayerRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadPlayerRecord not implemented")
}
func (*UnimplementedPlatformServiceServer) UploadPlayerSettleInfo(ctx context.Context, req *types.UploadPlayerSettleInfoRequest) (*types.UploadPlayerSettleInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadPlayerSettleInfo not implemented")
}
func (*UnimplementedPlatformServiceServer) GetRoomPool(ctx context.Context, req *types.GetRoomPoolRequest) (*types.GetRoomPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomPool not implemented")
}

func RegisterPlatformServiceServer(s *grpc.Server, srv PlatformServiceServer) {
	s.RegisterService(&_PlatformService_serviceDesc, srv)
}

func _PlatformService_ApplyRobot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.ApplyRobotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).ApplyRobot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform_service.PlatformService/ApplyRobot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).ApplyRobot(ctx, req.(*types.ApplyRobotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformService_PlayerLeaveGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.PlayerLeaveGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).PlayerLeaveGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform_service.PlatformService/PlayerLeaveGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).PlayerLeaveGame(ctx, req.(*types.PlayerLeaveGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformService_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.BroadcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform_service.PlatformService/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).Broadcast(ctx, req.(*types.BroadcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformService_FindRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.FindRoomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).FindRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform_service.PlatformService/FindRooms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).FindRooms(ctx, req.(*types.FindRoomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformService_RunningRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.RunningRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).RunningRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform_service.PlatformService/RunningRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).RunningRoom(ctx, req.(*types.RunningRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformService_AcquireRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.AcquireRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).AcquireRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform_service.PlatformService/AcquireRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).AcquireRoom(ctx, req.(*types.AcquireRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformService_ReleaseRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.ReleaseRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).ReleaseRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform_service.PlatformService/ReleaseRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).ReleaseRoom(ctx, req.(*types.ReleaseRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformService_UploadPlayerRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.UploadPlayerRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).UploadPlayerRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform_service.PlatformService/UploadPlayerRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).UploadPlayerRecord(ctx, req.(*types.UploadPlayerRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformService_UploadPlayerSettleInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.UploadPlayerSettleInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).UploadPlayerSettleInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform_service.PlatformService/UploadPlayerSettleInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).UploadPlayerSettleInfo(ctx, req.(*types.UploadPlayerSettleInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformService_GetRoomPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.GetRoomPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).GetRoomPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform_service.PlatformService/GetRoomPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).GetRoomPool(ctx, req.(*types.GetRoomPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlatformService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "platform_service.PlatformService",
	HandlerType: (*PlatformServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApplyRobot",
			Handler:    _PlatformService_ApplyRobot_Handler,
		},
		{
			MethodName: "PlayerLeaveGame",
			Handler:    _PlatformService_PlayerLeaveGame_Handler,
		},
		{
			MethodName: "Broadcast",
			Handler:    _PlatformService_Broadcast_Handler,
		},
		{
			MethodName: "FindRooms",
			Handler:    _PlatformService_FindRooms_Handler,
		},
		{
			MethodName: "RunningRoom",
			Handler:    _PlatformService_RunningRoom_Handler,
		},
		{
			MethodName: "AcquireRoom",
			Handler:    _PlatformService_AcquireRoom_Handler,
		},
		{
			MethodName: "ReleaseRoom",
			Handler:    _PlatformService_ReleaseRoom_Handler,
		},
		{
			MethodName: "UploadPlayerRecord",
			Handler:    _PlatformService_UploadPlayerRecord_Handler,
		},
		{
			MethodName: "UploadPlayerSettleInfo",
			Handler:    _PlatformService_UploadPlayerSettleInfo_Handler,
		},
		{
			MethodName: "GetRoomPool",
			Handler:    _PlatformService_GetRoomPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/service/platform/service.proto",
}