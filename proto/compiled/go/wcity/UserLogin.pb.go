// Code generated by protoc-gen-go. DO NOT EDIT.
// source: UserLogin.proto

package wcity

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type UserRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	HashPassword         string   `protobuf:"bytes,2,opt,name=hashPassword,proto3" json:"hashPassword,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62064063b9af919c, []int{0}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserRequest) GetHashPassword() string {
	if m != nil {
		return m.HashPassword
	}
	return ""
}

func (m *UserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UserResponse struct {
	IsUser               bool     `protobuf:"varint,1,opt,name=isUser,proto3" json:"isUser,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62064063b9af919c, []int{1}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetIsUser() bool {
	if m != nil {
		return m.IsUser
	}
	return false
}

func (m *UserResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LogRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogRequest) Reset()         { *m = LogRequest{} }
func (m *LogRequest) String() string { return proto.CompactTextString(m) }
func (*LogRequest) ProtoMessage()    {}
func (*LogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62064063b9af919c, []int{2}
}

func (m *LogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRequest.Unmarshal(m, b)
}
func (m *LogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRequest.Marshal(b, m, deterministic)
}
func (m *LogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRequest.Merge(m, src)
}
func (m *LogRequest) XXX_Size() int {
	return xxx_messageInfo_LogRequest.Size(m)
}
func (m *LogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogRequest proto.InternalMessageInfo

func (m *LogRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LogRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type LogResponse struct {
	Sucess               bool     `protobuf:"varint,1,opt,name=sucess,proto3" json:"sucess,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogResponse) Reset()         { *m = LogResponse{} }
func (m *LogResponse) String() string { return proto.CompactTextString(m) }
func (*LogResponse) ProtoMessage()    {}
func (*LogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62064063b9af919c, []int{3}
}

func (m *LogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogResponse.Unmarshal(m, b)
}
func (m *LogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogResponse.Marshal(b, m, deterministic)
}
func (m *LogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogResponse.Merge(m, src)
}
func (m *LogResponse) XXX_Size() int {
	return xxx_messageInfo_LogResponse.Size(m)
}
func (m *LogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogResponse proto.InternalMessageInfo

func (m *LogResponse) GetSucess() bool {
	if m != nil {
		return m.Sucess
	}
	return false
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "wcity.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "wcity.UserResponse")
	proto.RegisterType((*LogRequest)(nil), "wcity.LogRequest")
	proto.RegisterType((*LogResponse)(nil), "wcity.LogResponse")
}

func init() { proto.RegisterFile("UserLogin.proto", fileDescriptor_62064063b9af919c) }

var fileDescriptor_62064063b9af919c = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x25, 0xa9, 0x0d, 0x76, 0x5a, 0x14, 0xb7, 0x22, 0xa1, 0x27, 0x09, 0x08, 0x5e, 0x8c, 0x60,
	0x15, 0x3c, 0x78, 0xb1, 0xbd, 0x7a, 0x08, 0xc1, 0x9e, 0x3c, 0xad, 0xe9, 0xd0, 0x2c, 0x6d, 0x77,
	0x63, 0x66, 0x43, 0xf1, 0x67, 0x3c, 0xf8, 0xa5, 0x66, 0x37, 0xab, 0x49, 0x6f, 0xf6, 0xb6, 0x6f,
	0x66, 0xde, 0xbc, 0xc7, 0xdb, 0x81, 0xd3, 0x05, 0x61, 0xf9, 0xa2, 0x56, 0x42, 0xc6, 0x45, 0xa9,
	0xb4, 0x62, 0xfd, 0x5d, 0x26, 0xf4, 0x67, 0xf4, 0x06, 0x43, 0xd3, 0x49, 0xf1, 0xa3, 0x42, 0xd2,
	0xec, 0x1c, 0xfa, 0xb8, 0xe5, 0x62, 0x13, 0x7a, 0x97, 0xde, 0xf5, 0x20, 0x6d, 0x00, 0x8b, 0x60,
	0x94, 0x73, 0xca, 0x13, 0x4e, 0xb4, 0x53, 0xe5, 0x32, 0xf4, 0x6d, 0x73, 0xaf, 0xc6, 0x18, 0x1c,
	0x49, 0xbe, 0xc5, 0xb0, 0x67, 0x7b, 0xf6, 0x1d, 0x3d, 0xc1, 0xa8, 0x59, 0x4e, 0x85, 0x92, 0x84,
	0xec, 0x02, 0x02, 0x41, 0xa6, 0x62, 0xd7, 0x1f, 0xa7, 0x0e, 0x19, 0x55, 0xad, 0xd6, 0x28, 0xdd,
	0xe2, 0x06, 0x44, 0x8f, 0x00, 0xb5, 0xe1, 0x8e, 0xb3, 0x66, 0xc6, 0xeb, 0xcc, 0xb4, 0x7e, 0xfd,
	0x8e, 0xdf, 0xe8, 0x0a, 0x86, 0x96, 0xd9, 0xca, 0x52, 0x95, 0x21, 0xd1, 0xaf, 0x6c, 0x83, 0xee,
	0xbe, 0x7c, 0x60, 0x46, 0xff, 0xb9, 0xd2, 0x39, 0x4a, 0x2d, 0x32, 0xae, 0x85, 0x92, 0xec, 0x01,
	0x60, 0x5e, 0x22, 0xd7, 0x68, 0xbd, 0xb1, 0xd8, 0x06, 0x15, 0x77, 0x52, 0x9a, 0x8c, 0xf7, 0x6a,
	0x4e, 0xa5, 0xa6, 0x2d, 0x8a, 0xe5, 0xc1, 0xb4, 0x7b, 0x18, 0xd8, 0x6f, 0x39, 0x8c, 0x35, 0xad,
	0x3d, 0xe6, 0x98, 0xad, 0x5f, 0x6d, 0x0a, 0x67, 0x6e, 0xa4, 0x8d, 0x6b, 0xc2, 0xba, 0x25, 0x47,
	0xba, 0x85, 0xa0, 0x86, 0xaa, 0xd2, 0xff, 0x24, 0xcc, 0x6e, 0x60, 0x2c, 0x54, 0xbc, 0x2a, 0x8b,
	0xcc, 0x35, 0x37, 0xc6, 0xe9, 0xec, 0xe4, 0xef, 0x96, 0x12, 0x73, 0x4a, 0x89, 0xf7, 0xed, 0xf7,
	0xd2, 0x64, 0xfe, 0x1e, 0xd8, 0xcb, 0x9a, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x29, 0xa9,
	0x66, 0x6c, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserAuthenticationClient is the client API for UserAuthentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserAuthenticationClient interface {
	// created user and return login token
	CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	//update user
	UpdateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	// used for login
	LoginUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	// check if user is logged
	CheckToken(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error)
	// for logout
	Logout(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error)
}

type userAuthenticationClient struct {
	cc *grpc.ClientConn
}

func NewUserAuthenticationClient(cc *grpc.ClientConn) UserAuthenticationClient {
	return &userAuthenticationClient{cc}
}

func (c *userAuthenticationClient) CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/wcity.UserAuthentication/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAuthenticationClient) UpdateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/wcity.UserAuthentication/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAuthenticationClient) LoginUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/wcity.UserAuthentication/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAuthenticationClient) CheckToken(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := c.cc.Invoke(ctx, "/wcity.UserAuthentication/CheckToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAuthenticationClient) Logout(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := c.cc.Invoke(ctx, "/wcity.UserAuthentication/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAuthenticationServer is the server API for UserAuthentication service.
type UserAuthenticationServer interface {
	// created user and return login token
	CreateUser(context.Context, *UserRequest) (*UserResponse, error)
	//update user
	UpdateUser(context.Context, *UserRequest) (*UserResponse, error)
	// used for login
	LoginUser(context.Context, *UserRequest) (*UserResponse, error)
	// check if user is logged
	CheckToken(context.Context, *LogRequest) (*LogResponse, error)
	// for logout
	Logout(context.Context, *LogRequest) (*LogResponse, error)
}

// UnimplementedUserAuthenticationServer can be embedded to have forward compatible implementations.
type UnimplementedUserAuthenticationServer struct {
}

func (*UnimplementedUserAuthenticationServer) CreateUser(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserAuthenticationServer) UpdateUser(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedUserAuthenticationServer) LoginUser(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (*UnimplementedUserAuthenticationServer) CheckToken(ctx context.Context, req *LogRequest) (*LogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckToken not implemented")
}
func (*UnimplementedUserAuthenticationServer) Logout(ctx context.Context, req *LogRequest) (*LogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}

func RegisterUserAuthenticationServer(s *grpc.Server, srv UserAuthenticationServer) {
	s.RegisterService(&_UserAuthentication_serviceDesc, srv)
}

func _UserAuthentication_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthenticationServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wcity.UserAuthentication/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthenticationServer).CreateUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAuthentication_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthenticationServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wcity.UserAuthentication/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthenticationServer).UpdateUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAuthentication_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthenticationServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wcity.UserAuthentication/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthenticationServer).LoginUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAuthentication_CheckToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthenticationServer).CheckToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wcity.UserAuthentication/CheckToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthenticationServer).CheckToken(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAuthentication_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthenticationServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wcity.UserAuthentication/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthenticationServer).Logout(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserAuthentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wcity.UserAuthentication",
	HandlerType: (*UserAuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserAuthentication_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserAuthentication_UpdateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _UserAuthentication_LoginUser_Handler,
		},
		{
			MethodName: "CheckToken",
			Handler:    _UserAuthentication_CheckToken_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _UserAuthentication_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "UserLogin.proto",
}
