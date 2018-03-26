// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/protobuf/deploy/deploy.proto

/*
Package deploy is a generated protocol buffer package.

It is generated from these files:
	pkg/protobuf/deploy/deploy.proto

It has these top-level messages:
	DeployRequest
	DeployResponse
	ListRequest
	ListResponse
	RollbackRequest
	Empty
*/
package deploy

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

type DeployRequest struct {
	// Types that are valid to be assigned to Value:
	//	*DeployRequest_Info_
	//	*DeployRequest_File_
	Value isDeployRequest_Value `protobuf_oneof:"value"`
}

func (m *DeployRequest) Reset()                    { *m = DeployRequest{} }
func (m *DeployRequest) String() string            { return proto.CompactTextString(m) }
func (*DeployRequest) ProtoMessage()               {}
func (*DeployRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isDeployRequest_Value interface {
	isDeployRequest_Value()
}

type DeployRequest_Info_ struct {
	Info *DeployRequest_Info `protobuf:"bytes,1,opt,name=info,oneof"`
}
type DeployRequest_File_ struct {
	File *DeployRequest_File `protobuf:"bytes,2,opt,name=file,oneof"`
}

func (*DeployRequest_Info_) isDeployRequest_Value() {}
func (*DeployRequest_File_) isDeployRequest_Value() {}

func (m *DeployRequest) GetValue() isDeployRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *DeployRequest) GetInfo() *DeployRequest_Info {
	if x, ok := m.GetValue().(*DeployRequest_Info_); ok {
		return x.Info
	}
	return nil
}

func (m *DeployRequest) GetFile() *DeployRequest_File {
	if x, ok := m.GetValue().(*DeployRequest_File_); ok {
		return x.File
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DeployRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DeployRequest_OneofMarshaler, _DeployRequest_OneofUnmarshaler, _DeployRequest_OneofSizer, []interface{}{
		(*DeployRequest_Info_)(nil),
		(*DeployRequest_File_)(nil),
	}
}

func _DeployRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DeployRequest)
	// value
	switch x := m.Value.(type) {
	case *DeployRequest_Info_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Info); err != nil {
			return err
		}
	case *DeployRequest_File_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.File); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DeployRequest.Value has unexpected type %T", x)
	}
	return nil
}

func _DeployRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DeployRequest)
	switch tag {
	case 1: // value.info
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DeployRequest_Info)
		err := b.DecodeMessage(msg)
		m.Value = &DeployRequest_Info_{msg}
		return true, err
	case 2: // value.file
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DeployRequest_File)
		err := b.DecodeMessage(msg)
		m.Value = &DeployRequest_File_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DeployRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DeployRequest)
	// value
	switch x := m.Value.(type) {
	case *DeployRequest_Info_:
		s := proto.Size(x.Info)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DeployRequest_File_:
		s := proto.Size(x.File)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type DeployRequest_Info struct {
	App         string `protobuf:"bytes,1,opt,name=app" json:"app,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *DeployRequest_Info) Reset()                    { *m = DeployRequest_Info{} }
func (m *DeployRequest_Info) String() string            { return proto.CompactTextString(m) }
func (*DeployRequest_Info) ProtoMessage()               {}
func (*DeployRequest_Info) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *DeployRequest_Info) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *DeployRequest_Info) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type DeployRequest_File struct {
	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (m *DeployRequest_File) Reset()                    { *m = DeployRequest_File{} }
func (m *DeployRequest_File) String() string            { return proto.CompactTextString(m) }
func (*DeployRequest_File) ProtoMessage()               {}
func (*DeployRequest_File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *DeployRequest_File) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

type DeployResponse struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *DeployResponse) Reset()                    { *m = DeployResponse{} }
func (m *DeployResponse) String() string            { return proto.CompactTextString(m) }
func (*DeployResponse) ProtoMessage()               {}
func (*DeployResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeployResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type ListRequest struct {
	AppName string `protobuf:"bytes,1,opt,name=app_name,json=appName" json:"app_name,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListRequest) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

type ListResponse struct {
	Deploys []*ListResponse_Deploy `protobuf:"bytes,1,rep,name=deploys" json:"deploys,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListResponse) GetDeploys() []*ListResponse_Deploy {
	if m != nil {
		return m.Deploys
	}
	return nil
}

type ListResponse_Deploy struct {
	Revision    string `protobuf:"bytes,1,opt,name=revision" json:"revision,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Current     bool   `protobuf:"varint,4,opt,name=current" json:"current,omitempty"`
	CreatedAt   string `protobuf:"bytes,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *ListResponse_Deploy) Reset()                    { *m = ListResponse_Deploy{} }
func (m *ListResponse_Deploy) String() string            { return proto.CompactTextString(m) }
func (*ListResponse_Deploy) ProtoMessage()               {}
func (*ListResponse_Deploy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *ListResponse_Deploy) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *ListResponse_Deploy) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ListResponse_Deploy) GetCurrent() bool {
	if m != nil {
		return m.Current
	}
	return false
}

func (m *ListResponse_Deploy) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type RollbackRequest struct {
	AppName  string `protobuf:"bytes,1,opt,name=app_name,json=appName" json:"app_name,omitempty"`
	Revision string `protobuf:"bytes,2,opt,name=revision" json:"revision,omitempty"`
}

func (m *RollbackRequest) Reset()                    { *m = RollbackRequest{} }
func (m *RollbackRequest) String() string            { return proto.CompactTextString(m) }
func (*RollbackRequest) ProtoMessage()               {}
func (*RollbackRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RollbackRequest) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *RollbackRequest) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*DeployRequest)(nil), "deploy.DeployRequest")
	proto.RegisterType((*DeployRequest_Info)(nil), "deploy.DeployRequest.Info")
	proto.RegisterType((*DeployRequest_File)(nil), "deploy.DeployRequest.File")
	proto.RegisterType((*DeployResponse)(nil), "deploy.DeployResponse")
	proto.RegisterType((*ListRequest)(nil), "deploy.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "deploy.ListResponse")
	proto.RegisterType((*ListResponse_Deploy)(nil), "deploy.ListResponse.Deploy")
	proto.RegisterType((*RollbackRequest)(nil), "deploy.RollbackRequest")
	proto.RegisterType((*Empty)(nil), "deploy.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Deploy service

type DeployClient interface {
	Make(ctx context.Context, opts ...grpc.CallOption) (Deploy_MakeClient, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Rollback(ctx context.Context, in *RollbackRequest, opts ...grpc.CallOption) (*Empty, error)
}

type deployClient struct {
	cc *grpc.ClientConn
}

func NewDeployClient(cc *grpc.ClientConn) DeployClient {
	return &deployClient{cc}
}

func (c *deployClient) Make(ctx context.Context, opts ...grpc.CallOption) (Deploy_MakeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Deploy_serviceDesc.Streams[0], c.cc, "/deploy.Deploy/Make", opts...)
	if err != nil {
		return nil, err
	}
	x := &deployMakeClient{stream}
	return x, nil
}

type Deploy_MakeClient interface {
	Send(*DeployRequest) error
	Recv() (*DeployResponse, error)
	grpc.ClientStream
}

type deployMakeClient struct {
	grpc.ClientStream
}

func (x *deployMakeClient) Send(m *DeployRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *deployMakeClient) Recv() (*DeployResponse, error) {
	m := new(DeployResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *deployClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/deploy.Deploy/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deployClient) Rollback(ctx context.Context, in *RollbackRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/deploy.Deploy/Rollback", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Deploy service

type DeployServer interface {
	Make(Deploy_MakeServer) error
	List(context.Context, *ListRequest) (*ListResponse, error)
	Rollback(context.Context, *RollbackRequest) (*Empty, error)
}

func RegisterDeployServer(s *grpc.Server, srv DeployServer) {
	s.RegisterService(&_Deploy_serviceDesc, srv)
}

func _Deploy_Make_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DeployServer).Make(&deployMakeServer{stream})
}

type Deploy_MakeServer interface {
	Send(*DeployResponse) error
	Recv() (*DeployRequest, error)
	grpc.ServerStream
}

type deployMakeServer struct {
	grpc.ServerStream
}

func (x *deployMakeServer) Send(m *DeployResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *deployMakeServer) Recv() (*DeployRequest, error) {
	m := new(DeployRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Deploy_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deploy.Deploy/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deploy_Rollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployServer).Rollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deploy.Deploy/Rollback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployServer).Rollback(ctx, req.(*RollbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Deploy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "deploy.Deploy",
	HandlerType: (*DeployServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Deploy_List_Handler,
		},
		{
			MethodName: "Rollback",
			Handler:    _Deploy_Rollback_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Make",
			Handler:       _Deploy_Make_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/protobuf/deploy/deploy.proto",
}

func init() { proto.RegisterFile("pkg/protobuf/deploy/deploy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xd1, 0x6e, 0x94, 0x40,
	0x14, 0x75, 0xba, 0xb3, 0x0b, 0x7b, 0x69, 0xb5, 0x19, 0xab, 0x22, 0x6a, 0x42, 0x88, 0x0f, 0x3c,
	0x6d, 0x2b, 0xc6, 0x07, 0x7d, 0xd3, 0xa8, 0xa9, 0x46, 0x7d, 0xe0, 0x07, 0x9a, 0x59, 0xf6, 0xa2,
	0x04, 0x96, 0x19, 0x61, 0x68, 0xec, 0x07, 0xf8, 0x33, 0xfe, 0x8b, 0xaf, 0x7e, 0x8f, 0x99, 0x19,
	0x46, 0x4b, 0xd3, 0x74, 0x9f, 0xb8, 0xf7, 0x70, 0xce, 0xbd, 0xf7, 0x9c, 0x00, 0xc4, 0xb2, 0xfe,
	0x7a, 0x2c, 0x3b, 0xa1, 0xc4, 0x7a, 0x28, 0x8f, 0x37, 0x28, 0x1b, 0x71, 0x31, 0x3e, 0x56, 0x06,
	0x66, 0x0b, 0xdb, 0x25, 0x7f, 0x08, 0x1c, 0xbc, 0x35, 0x65, 0x8e, 0xdf, 0x07, 0xec, 0x15, 0x3b,
	0x01, 0x5a, 0xb5, 0xa5, 0x08, 0x49, 0x4c, 0xd2, 0x20, 0x8b, 0x56, 0xa3, 0x6c, 0x42, 0x5a, 0x7d,
	0x68, 0x4b, 0x71, 0x7a, 0x2b, 0x37, 0x4c, 0xad, 0x28, 0xab, 0x06, 0xc3, 0xbd, 0x9b, 0x14, 0xef,
	0xab, 0x06, 0xb5, 0x42, 0x33, 0xa3, 0x57, 0x40, 0xf5, 0x04, 0x76, 0x08, 0x33, 0x2e, 0xa5, 0x59,
	0xb5, 0xcc, 0x75, 0xc9, 0x62, 0x08, 0x36, 0xd8, 0x17, 0x5d, 0x25, 0x55, 0x25, 0x5a, 0x33, 0x72,
	0x99, 0x5f, 0x86, 0xa2, 0xc7, 0x40, 0xf5, 0x2c, 0x76, 0x04, 0xf3, 0xe2, 0xdb, 0xd0, 0xd6, 0x46,
	0xbd, 0x9f, 0xdb, 0xe6, 0x8d, 0x07, 0xf3, 0x73, 0xde, 0x0c, 0x98, 0x3c, 0x85, 0xdb, 0xee, 0x80,
	0x5e, 0x8a, 0xb6, 0x47, 0xc6, 0x80, 0x2a, 0xfc, 0xa1, 0xc6, 0x6d, 0xa6, 0x4e, 0x52, 0x08, 0x3e,
	0x55, 0xbd, 0x72, 0xde, 0x1f, 0x82, 0xcf, 0xa5, 0x3c, 0x6b, 0xf9, 0x16, 0x47, 0x9a, 0xc7, 0xa5,
	0xfc, 0xc2, 0xb7, 0x98, 0xfc, 0x26, 0xb0, 0x6f, 0xa9, 0xe3, 0xb8, 0x17, 0xe0, 0x59, 0xa3, 0x7d,
	0x48, 0xe2, 0x59, 0x1a, 0x64, 0x8f, 0x9c, 0xf1, 0xcb, 0x34, 0x97, 0x82, 0xe3, 0x46, 0x3f, 0x09,
	0x2c, 0x2c, 0xc6, 0x22, 0xf0, 0x3b, 0x3c, 0xaf, 0x7a, 0x6d, 0xd4, 0x6e, 0xfb, 0xd7, 0xef, 0xce,
	0x81, 0x85, 0xe0, 0x15, 0x43, 0xd7, 0x61, 0xab, 0x42, 0x1a, 0x93, 0xd4, 0xcf, 0x5d, 0xcb, 0x9e,
	0x00, 0x14, 0x1d, 0x72, 0x85, 0x9b, 0x33, 0xae, 0xc2, 0xb9, 0x91, 0x2e, 0x47, 0xe4, 0xb5, 0xfa,
	0x48, 0xfd, 0xd9, 0x21, 0x4d, 0x4e, 0xe1, 0x4e, 0x2e, 0x9a, 0x66, 0xcd, 0x8b, 0x7a, 0xb7, 0xfb,
	0xc9, 0xa9, 0x7b, 0xd3, 0x53, 0x13, 0x0f, 0xe6, 0xef, 0xb6, 0x52, 0x5d, 0x64, 0xbf, 0xfe, 0x5b,
	0x7b, 0x09, 0xf4, 0x33, 0xaf, 0x91, 0xdd, 0xbb, 0xf6, 0x63, 0x88, 0xee, 0x5f, 0x85, 0x6d, 0x58,
	0x29, 0x39, 0x21, 0xec, 0x19, 0x50, 0x1d, 0x20, 0xbb, 0x3b, 0x8d, 0xd3, 0x0a, 0x8f, 0xae, 0xcb,
	0x98, 0x65, 0xe0, 0x3b, 0x2f, 0xec, 0x81, 0x63, 0x5c, 0x71, 0x17, 0x1d, 0xb8, 0x17, 0xe6, 0xd8,
	0xf5, 0xc2, 0xfc, 0x07, 0xcf, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x89, 0xb1, 0x7d, 0xdd, 0x2b,
	0x03, 0x00, 0x00,
}
