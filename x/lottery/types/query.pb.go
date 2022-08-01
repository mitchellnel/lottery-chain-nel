// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lottery/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_559935ceda7539d5, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_559935ceda7539d5, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetOwnerRequest struct {
}

func (m *QueryGetOwnerRequest) Reset()         { *m = QueryGetOwnerRequest{} }
func (m *QueryGetOwnerRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetOwnerRequest) ProtoMessage()    {}
func (*QueryGetOwnerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_559935ceda7539d5, []int{2}
}
func (m *QueryGetOwnerRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetOwnerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetOwnerRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetOwnerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetOwnerRequest.Merge(m, src)
}
func (m *QueryGetOwnerRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetOwnerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetOwnerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetOwnerRequest proto.InternalMessageInfo

type QueryGetOwnerResponse struct {
	Owner Owner `protobuf:"bytes,1,opt,name=Owner,proto3" json:"Owner"`
}

func (m *QueryGetOwnerResponse) Reset()         { *m = QueryGetOwnerResponse{} }
func (m *QueryGetOwnerResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetOwnerResponse) ProtoMessage()    {}
func (*QueryGetOwnerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_559935ceda7539d5, []int{3}
}
func (m *QueryGetOwnerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetOwnerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetOwnerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetOwnerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetOwnerResponse.Merge(m, src)
}
func (m *QueryGetOwnerResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetOwnerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetOwnerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetOwnerResponse proto.InternalMessageInfo

func (m *QueryGetOwnerResponse) GetOwner() Owner {
	if m != nil {
		return m.Owner
	}
	return Owner{}
}

type QueryGetEntranceFeeRequest struct {
}

func (m *QueryGetEntranceFeeRequest) Reset()         { *m = QueryGetEntranceFeeRequest{} }
func (m *QueryGetEntranceFeeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetEntranceFeeRequest) ProtoMessage()    {}
func (*QueryGetEntranceFeeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_559935ceda7539d5, []int{4}
}
func (m *QueryGetEntranceFeeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetEntranceFeeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetEntranceFeeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetEntranceFeeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetEntranceFeeRequest.Merge(m, src)
}
func (m *QueryGetEntranceFeeRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetEntranceFeeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetEntranceFeeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetEntranceFeeRequest proto.InternalMessageInfo

type QueryGetEntranceFeeResponse struct {
	EntranceFee EntranceFee `protobuf:"bytes,1,opt,name=EntranceFee,proto3" json:"EntranceFee"`
}

func (m *QueryGetEntranceFeeResponse) Reset()         { *m = QueryGetEntranceFeeResponse{} }
func (m *QueryGetEntranceFeeResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetEntranceFeeResponse) ProtoMessage()    {}
func (*QueryGetEntranceFeeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_559935ceda7539d5, []int{5}
}
func (m *QueryGetEntranceFeeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetEntranceFeeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetEntranceFeeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetEntranceFeeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetEntranceFeeResponse.Merge(m, src)
}
func (m *QueryGetEntranceFeeResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetEntranceFeeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetEntranceFeeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetEntranceFeeResponse proto.InternalMessageInfo

func (m *QueryGetEntranceFeeResponse) GetEntranceFee() EntranceFee {
	if m != nil {
		return m.EntranceFee
	}
	return EntranceFee{}
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "lotterychainnel.lottery.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "lotterychainnel.lottery.QueryParamsResponse")
	proto.RegisterType((*QueryGetOwnerRequest)(nil), "lotterychainnel.lottery.QueryGetOwnerRequest")
	proto.RegisterType((*QueryGetOwnerResponse)(nil), "lotterychainnel.lottery.QueryGetOwnerResponse")
	proto.RegisterType((*QueryGetEntranceFeeRequest)(nil), "lotterychainnel.lottery.QueryGetEntranceFeeRequest")
	proto.RegisterType((*QueryGetEntranceFeeResponse)(nil), "lotterychainnel.lottery.QueryGetEntranceFeeResponse")
}

func init() { proto.RegisterFile("lottery/query.proto", fileDescriptor_559935ceda7539d5) }

var fileDescriptor_559935ceda7539d5 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0x3d, 0x40, 0xb3, 0x98, 0xee, 0xa6, 0xe6, 0x47, 0xa6, 0x72, 0x5b, 0x83, 0x44, 0x0b,
	0xc4, 0xa3, 0xb6, 0xac, 0x40, 0x6c, 0x2a, 0x01, 0x1b, 0x24, 0x20, 0xb0, 0x62, 0x83, 0x26, 0xd1,
	0xc5, 0x58, 0x38, 0x33, 0x8e, 0x67, 0x02, 0x64, 0xcb, 0x03, 0xa0, 0x48, 0x3c, 0x05, 0x6f, 0x92,
	0x65, 0x24, 0x16, 0xb0, 0x42, 0x28, 0xe1, 0x41, 0x50, 0x66, 0xae, 0x23, 0x27, 0xc1, 0x21, 0xdd,
	0x25, 0x67, 0xce, 0x3d, 0xdf, 0xf1, 0xfc, 0xd0, 0x9d, 0x4c, 0x19, 0x03, 0xc5, 0x80, 0xf7, 0xfa,
	0x50, 0x0c, 0xe2, 0xbc, 0x50, 0x46, 0xb1, 0xab, 0x28, 0x76, 0xde, 0x89, 0x54, 0x4a, 0xc8, 0x62,
	0xfc, 0x1f, 0xf8, 0x89, 0x4a, 0x94, 0xf5, 0xf0, 0xd9, 0x2f, 0x67, 0x0f, 0x76, 0x13, 0xa5, 0x92,
	0x0c, 0xb8, 0xc8, 0x53, 0x2e, 0xa4, 0x54, 0x46, 0x98, 0x54, 0x49, 0x8d, 0xab, 0xb7, 0x3b, 0x4a,
	0x77, 0x95, 0xe6, 0x6d, 0xa1, 0xc1, 0x51, 0xf8, 0x87, 0xe3, 0x36, 0x18, 0x71, 0xcc, 0x73, 0x91,
	0xa4, 0xd2, 0x9a, 0xd1, 0xeb, 0x97, 0x6d, 0x72, 0x51, 0x88, 0x6e, 0x99, 0x30, 0xef, 0xa8, 0x3e,
	0x4a, 0x28, 0x50, 0x0c, 0x4a, 0x11, 0xa4, 0x29, 0x84, 0xec, 0xc0, 0x9b, 0xb7, 0x00, 0x6e, 0x2d,
	0xf2, 0x29, 0x7b, 0x31, 0x03, 0x3d, 0xb7, 0x29, 0x2d, 0xe8, 0xf5, 0x41, 0x9b, 0xe8, 0x15, 0xdd,
	0x59, 0x50, 0x75, 0xae, 0xa4, 0x06, 0xf6, 0x90, 0x36, 0x1c, 0xed, 0x1a, 0xd9, 0x27, 0x87, 0xdb,
	0x27, 0x7b, 0x71, 0xcd, 0xd7, 0xc7, 0x6e, 0xf0, 0xec, 0xd2, 0xe8, 0xd7, 0x9e, 0xd7, 0xc2, 0xa1,
	0xe8, 0x0a, 0xf5, 0x6d, 0xea, 0x13, 0x30, 0xcf, 0x66, 0xf5, 0x4a, 0xda, 0x4b, 0x7a, 0x79, 0x49,
	0x47, 0xde, 0x7d, 0xba, 0x65, 0x05, 0xc4, 0x85, 0xb5, 0x38, 0xeb, 0x42, 0x9a, 0x1b, 0x89, 0x76,
	0x69, 0x50, 0x86, 0x3e, 0xc2, 0xcf, 0x7e, 0x0c, 0x50, 0x22, 0xdf, 0xd3, 0xeb, 0xff, 0x5c, 0x45,
	0xf0, 0x53, 0xba, 0x5d, 0x91, 0x11, 0x7f, 0xb3, 0x16, 0x5f, 0xf1, 0x62, 0x89, 0xea, 0xf8, 0xc9,
	0x8f, 0x8b, 0x74, 0xcb, 0xd2, 0xd8, 0x17, 0x42, 0x1b, 0x6e, 0x6b, 0xd8, 0x9d, 0xda, 0xb4, 0xd5,
	0xf3, 0x08, 0xee, 0x6e, 0x66, 0x76, 0xed, 0xa3, 0xa3, 0xcf, 0xdf, 0xff, 0x7c, 0xbd, 0x70, 0x83,
	0x1d, 0x70, 0x74, 0x35, 0xed, 0x58, 0x53, 0x42, 0xc6, 0x17, 0x6f, 0x0d, 0x1b, 0x12, 0xdc, 0x62,
	0xd6, 0x5c, 0x8f, 0x58, 0x3a, 0xb3, 0x20, 0xde, 0xd4, 0x8e, 0x9d, 0x0e, 0x6d, 0xa7, 0x88, 0xed,
	0xaf, 0xe9, 0x64, 0xef, 0x2c, 0xfb, 0x46, 0x16, 0x36, 0x9f, 0x9d, 0xfe, 0x97, 0xb4, 0x7a, 0xbe,
	0xc1, 0xbd, 0xf3, 0x0d, 0x61, 0x49, 0x6e, 0x4b, 0x1e, 0xb1, 0x5b, 0x6b, 0x4a, 0x56, 0xdf, 0xd0,
	0xd9, 0x83, 0xd1, 0x24, 0x24, 0xe3, 0x49, 0x48, 0x7e, 0x4f, 0x42, 0x32, 0x9c, 0x86, 0xde, 0x78,
	0x1a, 0x7a, 0x3f, 0xa7, 0xa1, 0xf7, 0xfa, 0x60, 0x35, 0xe1, 0xd3, 0x3c, 0xc3, 0x0c, 0x72, 0xd0,
	0xed, 0x86, 0x7d, 0x81, 0xa7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x81, 0xe5, 0xa6, 0x76, 0x58,
	0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a Owner by index.
	Owner(ctx context.Context, in *QueryGetOwnerRequest, opts ...grpc.CallOption) (*QueryGetOwnerResponse, error)
	// Queries a EntranceFee by index.
	EntranceFee(ctx context.Context, in *QueryGetEntranceFeeRequest, opts ...grpc.CallOption) (*QueryGetEntranceFeeResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/lotterychainnel.lottery.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Owner(ctx context.Context, in *QueryGetOwnerRequest, opts ...grpc.CallOption) (*QueryGetOwnerResponse, error) {
	out := new(QueryGetOwnerResponse)
	err := c.cc.Invoke(ctx, "/lotterychainnel.lottery.Query/Owner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EntranceFee(ctx context.Context, in *QueryGetEntranceFeeRequest, opts ...grpc.CallOption) (*QueryGetEntranceFeeResponse, error) {
	out := new(QueryGetEntranceFeeResponse)
	err := c.cc.Invoke(ctx, "/lotterychainnel.lottery.Query/EntranceFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a Owner by index.
	Owner(context.Context, *QueryGetOwnerRequest) (*QueryGetOwnerResponse, error)
	// Queries a EntranceFee by index.
	EntranceFee(context.Context, *QueryGetEntranceFeeRequest) (*QueryGetEntranceFeeResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Owner(ctx context.Context, req *QueryGetOwnerRequest) (*QueryGetOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Owner not implemented")
}
func (*UnimplementedQueryServer) EntranceFee(ctx context.Context, req *QueryGetEntranceFeeRequest) (*QueryGetEntranceFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EntranceFee not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lotterychainnel.lottery.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Owner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Owner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lotterychainnel.lottery.Query/Owner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Owner(ctx, req.(*QueryGetOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EntranceFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetEntranceFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EntranceFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lotterychainnel.lottery.Query/EntranceFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EntranceFee(ctx, req.(*QueryGetEntranceFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lotterychainnel.lottery.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Owner",
			Handler:    _Query_Owner_Handler,
		},
		{
			MethodName: "EntranceFee",
			Handler:    _Query_EntranceFee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lottery/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetOwnerRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetOwnerRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetOwnerRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryGetOwnerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetOwnerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetOwnerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Owner.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetEntranceFeeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetEntranceFeeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetEntranceFeeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryGetEntranceFeeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetEntranceFeeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetEntranceFeeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.EntranceFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetOwnerRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryGetOwnerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Owner.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetEntranceFeeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryGetEntranceFeeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EntranceFee.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetOwnerRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetOwnerRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetOwnerRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetOwnerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetOwnerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetOwnerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Owner.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetEntranceFeeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetEntranceFeeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetEntranceFeeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetEntranceFeeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetEntranceFeeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetEntranceFeeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntranceFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EntranceFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
