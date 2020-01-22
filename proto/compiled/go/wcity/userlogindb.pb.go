// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userlogindb.proto

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

// The user request.
type UserDBRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	HashedPassword       []byte   `protobuf:"bytes,2,opt,name=hashedPassword,proto3" json:"hashedPassword,omitempty"`
	Salt                 []byte   `protobuf:"bytes,3,opt,name=salt,proto3" json:"salt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDBRequest) Reset()         { *m = UserDBRequest{} }
func (m *UserDBRequest) String() string { return proto.CompactTextString(m) }
func (*UserDBRequest) ProtoMessage()    {}
func (*UserDBRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e0494eb4328d412, []int{0}
}

func (m *UserDBRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDBRequest.Unmarshal(m, b)
}
func (m *UserDBRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDBRequest.Marshal(b, m, deterministic)
}
func (m *UserDBRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDBRequest.Merge(m, src)
}
func (m *UserDBRequest) XXX_Size() int {
	return xxx_messageInfo_UserDBRequest.Size(m)
}
func (m *UserDBRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDBRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserDBRequest proto.InternalMessageInfo

func (m *UserDBRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserDBRequest) GetHashedPassword() []byte {
	if m != nil {
		return m.HashedPassword
	}
	return nil
}

func (m *UserDBRequest) GetSalt() []byte {
	if m != nil {
		return m.Salt
	}
	return nil
}

// the response
type UserDBResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	HashedPassword       []byte   `protobuf:"bytes,3,opt,name=hashedPassword,proto3" json:"hashedPassword,omitempty"`
	Salt                 []byte   `protobuf:"bytes,4,opt,name=salt,proto3" json:"salt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDBResponse) Reset()         { *m = UserDBResponse{} }
func (m *UserDBResponse) String() string { return proto.CompactTextString(m) }
func (*UserDBResponse) ProtoMessage()    {}
func (*UserDBResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e0494eb4328d412, []int{1}
}

func (m *UserDBResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDBResponse.Unmarshal(m, b)
}
func (m *UserDBResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDBResponse.Marshal(b, m, deterministic)
}
func (m *UserDBResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDBResponse.Merge(m, src)
}
func (m *UserDBResponse) XXX_Size() int {
	return xxx_messageInfo_UserDBResponse.Size(m)
}
func (m *UserDBResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDBResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserDBResponse proto.InternalMessageInfo

func (m *UserDBResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserDBResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserDBResponse) GetHashedPassword() []byte {
	if m != nil {
		return m.HashedPassword
	}
	return nil
}

func (m *UserDBResponse) GetSalt() []byte {
	if m != nil {
		return m.Salt
	}
	return nil
}

func init() {
	proto.RegisterType((*UserDBRequest)(nil), "wcity.UserDBRequest")
	proto.RegisterType((*UserDBResponse)(nil), "wcity.UserDBResponse")
}

func init() { proto.RegisterFile("userlogindb.proto", fileDescriptor_9e0494eb4328d412) }

var fileDescriptor_9e0494eb4328d412 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x2d, 0x4e, 0x2d,
	0xca, 0xc9, 0x4f, 0xcf, 0xcc, 0x4b, 0x49, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d,
	0x4f, 0xce, 0x2c, 0xa9, 0x54, 0x4a, 0xe4, 0xe2, 0x0d, 0x05, 0xca, 0xb9, 0x38, 0x05, 0xa5, 0x16,
	0x96, 0xa6, 0x16, 0x97, 0x08, 0x89, 0x70, 0xb1, 0xa6, 0xe6, 0x26, 0x66, 0xe6, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x42, 0x6a, 0x5c, 0x7c, 0x19, 0x89, 0xc5, 0x19, 0xa9, 0x29,
	0x01, 0x89, 0xc5, 0xc5, 0xe5, 0xf9, 0x45, 0x29, 0x12, 0x4c, 0x40, 0x69, 0x9e, 0x20, 0x34, 0x51,
	0x21, 0x21, 0x2e, 0x96, 0xe2, 0xc4, 0x9c, 0x12, 0x09, 0x66, 0xb0, 0x2c, 0x98, 0xad, 0x54, 0xc4,
	0xc5, 0x07, 0xb3, 0xa2, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x05,
	0x6c, 0x01, 0x6b, 0x10, 0x90, 0x85, 0xb0, 0x93, 0x09, 0xbf, 0x9d, 0xcc, 0x78, 0xed, 0x64, 0x41,
	0xd8, 0x69, 0xb4, 0x8d, 0x91, 0x8b, 0x13, 0x64, 0xa9, 0x4f, 0x7e, 0xba, 0x8b, 0x93, 0x90, 0x19,
	0x17, 0xbb, 0x63, 0x4a, 0x0a, 0x88, 0x2f, 0x24, 0xa2, 0x07, 0xf6, 0xb7, 0x1e, 0x8a, 0xa7, 0xa5,
	0x44, 0xd1, 0x44, 0xa1, 0xee, 0x04, 0xea, 0x73, 0x4f, 0x2d, 0x21, 0x5d, 0x9f, 0x25, 0x17, 0x57,
	0x68, 0x41, 0x4a, 0x62, 0x49, 0x2a, 0xc9, 0x5a, 0x9d, 0x74, 0xb9, 0x84, 0x33, 0xf3, 0xf5, 0xd2,
	0x8b, 0x0a, 0x92, 0xa1, 0xf2, 0xe0, 0x58, 0x73, 0xe2, 0x83, 0x7a, 0x26, 0x33, 0x2f, 0x00, 0x14,
	0x7b, 0x01, 0x8c, 0x8b, 0x98, 0x98, 0x83, 0x02, 0x9c, 0x93, 0xd8, 0xc0, 0x91, 0x69, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x2f, 0x60, 0x22, 0x8b, 0xe1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserLogDBClient is the client API for UserLogDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserLogDBClient interface {
	// add new user, response include id
	AddUser(ctx context.Context, in *UserDBRequest, opts ...grpc.CallOption) (*UserDBResponse, error)
	// get user details, request should include only the email
	GetUser(ctx context.Context, in *UserDBRequest, opts ...grpc.CallOption) (*UserDBResponse, error)
	//update a user.
	UpdateUser(ctx context.Context, in *UserDBRequest, opts ...grpc.CallOption) (*UserDBResponse, error)
}

type userLogDBClient struct {
	cc *grpc.ClientConn
}

func NewUserLogDBClient(cc *grpc.ClientConn) UserLogDBClient {
	return &userLogDBClient{cc}
}

func (c *userLogDBClient) AddUser(ctx context.Context, in *UserDBRequest, opts ...grpc.CallOption) (*UserDBResponse, error) {
	out := new(UserDBResponse)
	err := c.cc.Invoke(ctx, "/wcity.UserLogDB/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLogDBClient) GetUser(ctx context.Context, in *UserDBRequest, opts ...grpc.CallOption) (*UserDBResponse, error) {
	out := new(UserDBResponse)
	err := c.cc.Invoke(ctx, "/wcity.UserLogDB/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLogDBClient) UpdateUser(ctx context.Context, in *UserDBRequest, opts ...grpc.CallOption) (*UserDBResponse, error) {
	out := new(UserDBResponse)
	err := c.cc.Invoke(ctx, "/wcity.UserLogDB/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserLogDBServer is the server API for UserLogDB service.
type UserLogDBServer interface {
	// add new user, response include id
	AddUser(context.Context, *UserDBRequest) (*UserDBResponse, error)
	// get user details, request should include only the email
	GetUser(context.Context, *UserDBRequest) (*UserDBResponse, error)
	//update a user.
	UpdateUser(context.Context, *UserDBRequest) (*UserDBResponse, error)
}

// UnimplementedUserLogDBServer can be embedded to have forward compatible implementations.
type UnimplementedUserLogDBServer struct {
}

func (*UnimplementedUserLogDBServer) AddUser(ctx context.Context, req *UserDBRequest) (*UserDBResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (*UnimplementedUserLogDBServer) GetUser(ctx context.Context, req *UserDBRequest) (*UserDBResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserLogDBServer) UpdateUser(ctx context.Context, req *UserDBRequest) (*UserDBResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

func RegisterUserLogDBServer(s *grpc.Server, srv UserLogDBServer) {
	s.RegisterService(&_UserLogDB_serviceDesc, srv)
}

func _UserLogDB_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLogDBServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wcity.UserLogDB/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLogDBServer).AddUser(ctx, req.(*UserDBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLogDB_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLogDBServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wcity.UserLogDB/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLogDBServer).GetUser(ctx, req.(*UserDBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLogDB_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLogDBServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wcity.UserLogDB/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLogDBServer).UpdateUser(ctx, req.(*UserDBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserLogDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wcity.UserLogDB",
	HandlerType: (*UserLogDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _UserLogDB_AddUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserLogDB_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserLogDB_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userlogindb.proto",
}