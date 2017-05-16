// Code generated by protoc-gen-go.
// source: pkg/protobuf/user.proto
// DO NOT EDIT!

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	pkg/protobuf/user.proto

It has these top-level messages:
	LoginRequest
	LoginResponse
	SetPasswordRequest
	DeleteRequest
	Empty
*/
package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LoginRequest struct {
	Email    string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type SetPasswordRequest struct {
	Password string `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
}

func (m *SetPasswordRequest) Reset()                    { *m = SetPasswordRequest{} }
func (m *SetPasswordRequest) String() string            { return proto.CompactTextString(m) }
func (*SetPasswordRequest) ProtoMessage()               {}
func (*SetPasswordRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SetPasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type DeleteRequest struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DeleteRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*LoginRequest)(nil), "user.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "user.LoginResponse")
	proto.RegisterType((*SetPasswordRequest)(nil), "user.SetPasswordRequest")
	proto.RegisterType((*DeleteRequest)(nil), "user.DeleteRequest")
	proto.RegisterType((*Empty)(nil), "user.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for User service

type UserClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	SetPassword(ctx context.Context, in *SetPasswordRequest, opts ...grpc.CallOption) (*Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Empty, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/user.User/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SetPassword(ctx context.Context, in *SetPasswordRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/user.User/SetPassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/user.User/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	SetPassword(context.Context, *SetPasswordRequest) (*Empty, error)
	Delete(context.Context, *DeleteRequest) (*Empty, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/SetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SetPassword(ctx, req.(*SetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "SetPassword",
			Handler:    _User_SetPassword_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _User_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/protobuf/user.proto",
}

func init() { proto.RegisterFile("pkg/protobuf/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0xc8, 0x4e, 0xd7,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x03, 0xf3,
	0x84, 0x58, 0x40, 0x6c, 0x25, 0x07, 0x2e, 0x1e, 0x9f, 0xfc, 0xf4, 0xcc, 0xbc, 0xa0, 0xd4, 0xc2,
	0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x11, 0x2e, 0xd6, 0xd4, 0xdc, 0xc4, 0xcc, 0x1c, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x48, 0x8a, 0x8b, 0xa3, 0x20, 0xb1, 0xb8, 0xb8, 0x3c, 0xbf,
	0x28, 0x45, 0x82, 0x09, 0x2c, 0x01, 0xe7, 0x2b, 0xa9, 0x72, 0xf1, 0x42, 0x4d, 0x28, 0x2e, 0xc8,
	0xcf, 0x2b, 0x4e, 0x05, 0x19, 0x51, 0x92, 0x9f, 0x9d, 0x9a, 0x07, 0x33, 0x02, 0xcc, 0x51, 0x32,
	0xe0, 0x12, 0x0a, 0x4e, 0x2d, 0x09, 0x80, 0xea, 0x82, 0x59, 0x87, 0x6c, 0x30, 0x23, 0xa6, 0xc1,
	0x2e, 0xa9, 0x39, 0xa9, 0x25, 0xa9, 0x78, 0xdd, 0xa6, 0xc4, 0xce, 0xc5, 0xea, 0x9a, 0x5b, 0x50,
	0x52, 0x69, 0x34, 0x8b, 0x91, 0x8b, 0x25, 0xb4, 0x38, 0xb5, 0x48, 0xc8, 0x80, 0x8b, 0x15, 0xec,
	0x22, 0x21, 0x21, 0x3d, 0xb0, 0x7f, 0x91, 0x3d, 0x28, 0x25, 0x8c, 0x22, 0x06, 0x75, 0xb2, 0x09,
	0x17, 0x37, 0x92, 0xe3, 0x84, 0x24, 0x20, 0x6a, 0x30, 0xdd, 0x2b, 0xc5, 0x0d, 0x91, 0x01, 0x5b,
	0x28, 0xa4, 0xc5, 0xc5, 0x06, 0x71, 0xa0, 0x10, 0xd4, 0x50, 0x14, 0xe7, 0xa2, 0xa8, 0x4d, 0x62,
	0x03, 0x07, 0xba, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x44, 0x7c, 0x6c, 0x8f, 0x01, 0x00,
	0x00,
}
