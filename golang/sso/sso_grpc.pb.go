// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: sso/sso.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SSOService_CreateSession_FullMethodName   = "/sso.v1.SSOService/CreateSession"
	SSOService_ValidateSession_FullMethodName = "/sso.v1.SSOService/ValidateSession"
	SSOService_Logout_FullMethodName          = "/sso.v1.SSOService/Logout"
	SSOService_RevokeSessions_FullMethodName  = "/sso.v1.SSOService/RevokeSessions"
	SSOService_HealthCheck_FullMethodName     = "/sso.v1.SSOService/HealthCheck"
)

// SSOServiceClient is the client API for SSOService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// SSO service for managing authentication sessions
type SSOServiceClient interface {
	// Creates a new session for authenticated user
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error)
	// Validates existing session token
	ValidateSession(ctx context.Context, in *ValidateSessionRequest, opts ...grpc.CallOption) (*ValidateSessionResponse, error)
	// Logs out user by invalidating their session
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	// Revokes all sessions for a specific user and application
	RevokeSessions(ctx context.Context, in *RevokeSessionsRequest, opts ...grpc.CallOption) (*RevokeSessionsResponse, error)
	// Checks service health status
	HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type sSOServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSSOServiceClient(cc grpc.ClientConnInterface) SSOServiceClient {
	return &sSOServiceClient{cc}
}

func (c *sSOServiceClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSessionResponse)
	err := c.cc.Invoke(ctx, SSOService_CreateSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSOServiceClient) ValidateSession(ctx context.Context, in *ValidateSessionRequest, opts ...grpc.CallOption) (*ValidateSessionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidateSessionResponse)
	err := c.cc.Invoke(ctx, SSOService_ValidateSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSOServiceClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, SSOService_Logout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSOServiceClient) RevokeSessions(ctx context.Context, in *RevokeSessionsRequest, opts ...grpc.CallOption) (*RevokeSessionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RevokeSessionsResponse)
	err := c.cc.Invoke(ctx, SSOService_RevokeSessions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSOServiceClient) HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, SSOService_HealthCheck_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SSOServiceServer is the server API for SSOService service.
// All implementations must embed UnimplementedSSOServiceServer
// for forward compatibility.
//
// SSO service for managing authentication sessions
type SSOServiceServer interface {
	// Creates a new session for authenticated user
	CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error)
	// Validates existing session token
	ValidateSession(context.Context, *ValidateSessionRequest) (*ValidateSessionResponse, error)
	// Logs out user by invalidating their session
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	// Revokes all sessions for a specific user and application
	RevokeSessions(context.Context, *RevokeSessionsRequest) (*RevokeSessionsResponse, error)
	// Checks service health status
	HealthCheck(context.Context, *emptypb.Empty) (*HealthCheckResponse, error)
	mustEmbedUnimplementedSSOServiceServer()
}

// UnimplementedSSOServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSSOServiceServer struct{}

func (UnimplementedSSOServiceServer) CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedSSOServiceServer) ValidateSession(context.Context, *ValidateSessionRequest) (*ValidateSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateSession not implemented")
}
func (UnimplementedSSOServiceServer) Logout(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedSSOServiceServer) RevokeSessions(context.Context, *RevokeSessionsRequest) (*RevokeSessionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeSessions not implemented")
}
func (UnimplementedSSOServiceServer) HealthCheck(context.Context, *emptypb.Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedSSOServiceServer) mustEmbedUnimplementedSSOServiceServer() {}
func (UnimplementedSSOServiceServer) testEmbeddedByValue()                    {}

// UnsafeSSOServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SSOServiceServer will
// result in compilation errors.
type UnsafeSSOServiceServer interface {
	mustEmbedUnimplementedSSOServiceServer()
}

func RegisterSSOServiceServer(s grpc.ServiceRegistrar, srv SSOServiceServer) {
	// If the following call pancis, it indicates UnimplementedSSOServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SSOService_ServiceDesc, srv)
}

func _SSOService_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSOServiceServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSOService_CreateSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSOServiceServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSOService_ValidateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSOServiceServer).ValidateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSOService_ValidateSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSOServiceServer).ValidateSession(ctx, req.(*ValidateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSOService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSOServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSOService_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSOServiceServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSOService_RevokeSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeSessionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSOServiceServer).RevokeSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSOService_RevokeSessions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSOServiceServer).RevokeSessions(ctx, req.(*RevokeSessionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSOService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSOServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSOService_HealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSOServiceServer).HealthCheck(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// SSOService_ServiceDesc is the grpc.ServiceDesc for SSOService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SSOService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sso.v1.SSOService",
	HandlerType: (*SSOServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSession",
			Handler:    _SSOService_CreateSession_Handler,
		},
		{
			MethodName: "ValidateSession",
			Handler:    _SSOService_ValidateSession_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _SSOService_Logout_Handler,
		},
		{
			MethodName: "RevokeSessions",
			Handler:    _SSOService_RevokeSessions_Handler,
		},
		{
			MethodName: "HealthCheck",
			Handler:    _SSOService_HealthCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sso/sso.proto",
}
