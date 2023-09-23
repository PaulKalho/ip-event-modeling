// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/services/impulse_svc/v1/impulse_svc.proto

package impulse_svc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ImpulseService_CreateUser_FullMethodName          = "/proto.services.impulse_svc.v1.ImpulseService/CreateUser"
	ImpulseService_UpdateUser_FullMethodName          = "/proto.services.impulse_svc.v1.ImpulseService/UpdateUser"
	ImpulseService_TrackChallenge_FullMethodName      = "/proto.services.impulse_svc.v1.ImpulseService/TrackChallenge"
	ImpulseService_GetActiveChallenges_FullMethodName = "/proto.services.impulse_svc.v1.ImpulseService/GetActiveChallenges"
	ImpulseService_GetScore_FullMethodName            = "/proto.services.impulse_svc.v1.ImpulseService/GetScore"
	ImpulseService_GetRewards_FullMethodName          = "/proto.services.impulse_svc.v1.ImpulseService/GetRewards"
	ImpulseService_GetAllRewards_FullMethodName       = "/proto.services.impulse_svc.v1.ImpulseService/GetAllRewards"
	ImpulseService_GetAllTeams_FullMethodName         = "/proto.services.impulse_svc.v1.ImpulseService/GetAllTeams"
	ImpulseService_StatsForTeamByUser_FullMethodName  = "/proto.services.impulse_svc.v1.ImpulseService/StatsForTeamByUser"
)

// ImpulseServiceClient is the client API for ImpulseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImpulseServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	TrackChallenge(ctx context.Context, in *TrackChallengeRequest, opts ...grpc.CallOption) (*TrackChallengeResponse, error)
	GetActiveChallenges(ctx context.Context, in *GetActiveChallengesRequest, opts ...grpc.CallOption) (*GetActiveChallengesResponse, error)
	GetScore(ctx context.Context, in *GetScoreRequest, opts ...grpc.CallOption) (*GetScoreResponse, error)
	GetRewards(ctx context.Context, in *GetRewardsRequest, opts ...grpc.CallOption) (*GetRewardsResponse, error)
	GetAllRewards(ctx context.Context, in *GetAllRewardsRequest, opts ...grpc.CallOption) (*GetAllRewardsResponse, error)
	GetAllTeams(ctx context.Context, in *GetAllTeamsRequest, opts ...grpc.CallOption) (*GetAllTeamsResponse, error)
	StatsForTeamByUser(ctx context.Context, in *StatsForTeamByUserRequest, opts ...grpc.CallOption) (*StatsForTeamByUserResponse, error)
}

type impulseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImpulseServiceClient(cc grpc.ClientConnInterface) ImpulseServiceClient {
	return &impulseServiceClient{cc}
}

func (c *impulseServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, ImpulseService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *impulseServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, ImpulseService_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *impulseServiceClient) TrackChallenge(ctx context.Context, in *TrackChallengeRequest, opts ...grpc.CallOption) (*TrackChallengeResponse, error) {
	out := new(TrackChallengeResponse)
	err := c.cc.Invoke(ctx, ImpulseService_TrackChallenge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *impulseServiceClient) GetActiveChallenges(ctx context.Context, in *GetActiveChallengesRequest, opts ...grpc.CallOption) (*GetActiveChallengesResponse, error) {
	out := new(GetActiveChallengesResponse)
	err := c.cc.Invoke(ctx, ImpulseService_GetActiveChallenges_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *impulseServiceClient) GetScore(ctx context.Context, in *GetScoreRequest, opts ...grpc.CallOption) (*GetScoreResponse, error) {
	out := new(GetScoreResponse)
	err := c.cc.Invoke(ctx, ImpulseService_GetScore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *impulseServiceClient) GetRewards(ctx context.Context, in *GetRewardsRequest, opts ...grpc.CallOption) (*GetRewardsResponse, error) {
	out := new(GetRewardsResponse)
	err := c.cc.Invoke(ctx, ImpulseService_GetRewards_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *impulseServiceClient) GetAllRewards(ctx context.Context, in *GetAllRewardsRequest, opts ...grpc.CallOption) (*GetAllRewardsResponse, error) {
	out := new(GetAllRewardsResponse)
	err := c.cc.Invoke(ctx, ImpulseService_GetAllRewards_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *impulseServiceClient) GetAllTeams(ctx context.Context, in *GetAllTeamsRequest, opts ...grpc.CallOption) (*GetAllTeamsResponse, error) {
	out := new(GetAllTeamsResponse)
	err := c.cc.Invoke(ctx, ImpulseService_GetAllTeams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *impulseServiceClient) StatsForTeamByUser(ctx context.Context, in *StatsForTeamByUserRequest, opts ...grpc.CallOption) (*StatsForTeamByUserResponse, error) {
	out := new(StatsForTeamByUserResponse)
	err := c.cc.Invoke(ctx, ImpulseService_StatsForTeamByUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImpulseServiceServer is the server API for ImpulseService service.
// All implementations must embed UnimplementedImpulseServiceServer
// for forward compatibility
type ImpulseServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	TrackChallenge(context.Context, *TrackChallengeRequest) (*TrackChallengeResponse, error)
	GetActiveChallenges(context.Context, *GetActiveChallengesRequest) (*GetActiveChallengesResponse, error)
	GetScore(context.Context, *GetScoreRequest) (*GetScoreResponse, error)
	GetRewards(context.Context, *GetRewardsRequest) (*GetRewardsResponse, error)
	GetAllRewards(context.Context, *GetAllRewardsRequest) (*GetAllRewardsResponse, error)
	GetAllTeams(context.Context, *GetAllTeamsRequest) (*GetAllTeamsResponse, error)
	StatsForTeamByUser(context.Context, *StatsForTeamByUserRequest) (*StatsForTeamByUserResponse, error)
	mustEmbedUnimplementedImpulseServiceServer()
}

// UnimplementedImpulseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImpulseServiceServer struct {
}

func (UnimplementedImpulseServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedImpulseServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedImpulseServiceServer) TrackChallenge(context.Context, *TrackChallengeRequest) (*TrackChallengeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrackChallenge not implemented")
}
func (UnimplementedImpulseServiceServer) GetActiveChallenges(context.Context, *GetActiveChallengesRequest) (*GetActiveChallengesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveChallenges not implemented")
}
func (UnimplementedImpulseServiceServer) GetScore(context.Context, *GetScoreRequest) (*GetScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScore not implemented")
}
func (UnimplementedImpulseServiceServer) GetRewards(context.Context, *GetRewardsRequest) (*GetRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRewards not implemented")
}
func (UnimplementedImpulseServiceServer) GetAllRewards(context.Context, *GetAllRewardsRequest) (*GetAllRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRewards not implemented")
}
func (UnimplementedImpulseServiceServer) GetAllTeams(context.Context, *GetAllTeamsRequest) (*GetAllTeamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTeams not implemented")
}
func (UnimplementedImpulseServiceServer) StatsForTeamByUser(context.Context, *StatsForTeamByUserRequest) (*StatsForTeamByUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatsForTeamByUser not implemented")
}
func (UnimplementedImpulseServiceServer) mustEmbedUnimplementedImpulseServiceServer() {}

// UnsafeImpulseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImpulseServiceServer will
// result in compilation errors.
type UnsafeImpulseServiceServer interface {
	mustEmbedUnimplementedImpulseServiceServer()
}

func RegisterImpulseServiceServer(s grpc.ServiceRegistrar, srv ImpulseServiceServer) {
	s.RegisterService(&ImpulseService_ServiceDesc, srv)
}

func _ImpulseService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImpulseServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImpulseService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImpulseServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImpulseService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImpulseServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImpulseService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImpulseServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImpulseService_TrackChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImpulseServiceServer).TrackChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImpulseService_TrackChallenge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImpulseServiceServer).TrackChallenge(ctx, req.(*TrackChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImpulseService_GetActiveChallenges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActiveChallengesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImpulseServiceServer).GetActiveChallenges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImpulseService_GetActiveChallenges_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImpulseServiceServer).GetActiveChallenges(ctx, req.(*GetActiveChallengesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImpulseService_GetScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImpulseServiceServer).GetScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImpulseService_GetScore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImpulseServiceServer).GetScore(ctx, req.(*GetScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImpulseService_GetRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImpulseServiceServer).GetRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImpulseService_GetRewards_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImpulseServiceServer).GetRewards(ctx, req.(*GetRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImpulseService_GetAllRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImpulseServiceServer).GetAllRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImpulseService_GetAllRewards_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImpulseServiceServer).GetAllRewards(ctx, req.(*GetAllRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImpulseService_GetAllTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllTeamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImpulseServiceServer).GetAllTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImpulseService_GetAllTeams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImpulseServiceServer).GetAllTeams(ctx, req.(*GetAllTeamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImpulseService_StatsForTeamByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsForTeamByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImpulseServiceServer).StatsForTeamByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImpulseService_StatsForTeamByUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImpulseServiceServer).StatsForTeamByUser(ctx, req.(*StatsForTeamByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImpulseService_ServiceDesc is the grpc.ServiceDesc for ImpulseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImpulseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.services.impulse_svc.v1.ImpulseService",
	HandlerType: (*ImpulseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _ImpulseService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _ImpulseService_UpdateUser_Handler,
		},
		{
			MethodName: "TrackChallenge",
			Handler:    _ImpulseService_TrackChallenge_Handler,
		},
		{
			MethodName: "GetActiveChallenges",
			Handler:    _ImpulseService_GetActiveChallenges_Handler,
		},
		{
			MethodName: "GetScore",
			Handler:    _ImpulseService_GetScore_Handler,
		},
		{
			MethodName: "GetRewards",
			Handler:    _ImpulseService_GetRewards_Handler,
		},
		{
			MethodName: "GetAllRewards",
			Handler:    _ImpulseService_GetAllRewards_Handler,
		},
		{
			MethodName: "GetAllTeams",
			Handler:    _ImpulseService_GetAllTeams_Handler,
		},
		{
			MethodName: "StatsForTeamByUser",
			Handler:    _ImpulseService_StatsForTeamByUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/services/impulse_svc/v1/impulse_svc.proto",
}
