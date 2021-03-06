// Code generated by protoc-gen-go.
// source: rannu.proto
// DO NOT EDIT!

/*
Package rannu is a generated protocol buffer package.

It is generated from these files:
	rannu.proto

It has these top-level messages:
	Unit
	DataFile
	Size
	Vector
	Matrix
*/
package rannu

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

type Unit struct {
}

func (m *Unit) Reset()                    { *m = Unit{} }
func (m *Unit) String() string            { return proto.CompactTextString(m) }
func (*Unit) ProtoMessage()               {}
func (*Unit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DataFile struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DataFile) Reset()                    { *m = DataFile{} }
func (m *DataFile) String() string            { return proto.CompactTextString(m) }
func (*DataFile) ProtoMessage()               {}
func (*DataFile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Size struct {
	Rows int32 `protobuf:"varint,1,opt,name=rows" json:"rows,omitempty"`
	Cols int32 `protobuf:"varint,2,opt,name=cols" json:"cols,omitempty"`
}

func (m *Size) Reset()                    { *m = Size{} }
func (m *Size) String() string            { return proto.CompactTextString(m) }
func (*Size) ProtoMessage()               {}
func (*Size) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Vector struct {
	Elements []float64 `protobuf:"fixed64,1,rep,packed,name=elements" json:"elements,omitempty"`
}

func (m *Vector) Reset()                    { *m = Vector{} }
func (m *Vector) String() string            { return proto.CompactTextString(m) }
func (*Vector) ProtoMessage()               {}
func (*Vector) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Matrix struct {
	Elements []*Vector `protobuf:"bytes,1,rep,name=elements" json:"elements,omitempty"`
}

func (m *Matrix) Reset()                    { *m = Matrix{} }
func (m *Matrix) String() string            { return proto.CompactTextString(m) }
func (*Matrix) ProtoMessage()               {}
func (*Matrix) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Matrix) GetElements() []*Vector {
	if m != nil {
		return m.Elements
	}
	return nil
}

func init() {
	proto.RegisterType((*Unit)(nil), "rannu.Unit")
	proto.RegisterType((*DataFile)(nil), "rannu.DataFile")
	proto.RegisterType((*Size)(nil), "rannu.Size")
	proto.RegisterType((*Vector)(nil), "rannu.Vector")
	proto.RegisterType((*Matrix)(nil), "rannu.Matrix")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Worker service

type WorkerClient interface {
	LoadData(ctx context.Context, in *DataFile, opts ...grpc.CallOption) (*Size, error)
	GetSum(ctx context.Context, in *Unit, opts ...grpc.CallOption) (*Vector, error)
	GetScatterMatrix(ctx context.Context, in *Vector, opts ...grpc.CallOption) (*Matrix, error)
	ComputeScores(ctx context.Context, in *Matrix, opts ...grpc.CallOption) (*DataFile, error)
}

type workerClient struct {
	cc *grpc.ClientConn
}

func NewWorkerClient(cc *grpc.ClientConn) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) LoadData(ctx context.Context, in *DataFile, opts ...grpc.CallOption) (*Size, error) {
	out := new(Size)
	err := grpc.Invoke(ctx, "/rannu.Worker/LoadData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) GetSum(ctx context.Context, in *Unit, opts ...grpc.CallOption) (*Vector, error) {
	out := new(Vector)
	err := grpc.Invoke(ctx, "/rannu.Worker/GetSum", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) GetScatterMatrix(ctx context.Context, in *Vector, opts ...grpc.CallOption) (*Matrix, error) {
	out := new(Matrix)
	err := grpc.Invoke(ctx, "/rannu.Worker/GetScatterMatrix", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) ComputeScores(ctx context.Context, in *Matrix, opts ...grpc.CallOption) (*DataFile, error) {
	out := new(DataFile)
	err := grpc.Invoke(ctx, "/rannu.Worker/ComputeScores", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Worker service

type WorkerServer interface {
	LoadData(context.Context, *DataFile) (*Size, error)
	GetSum(context.Context, *Unit) (*Vector, error)
	GetScatterMatrix(context.Context, *Vector) (*Matrix, error)
	ComputeScores(context.Context, *Matrix) (*DataFile, error)
}

func RegisterWorkerServer(s *grpc.Server, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

func _Worker_LoadData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataFile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).LoadData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rannu.Worker/LoadData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).LoadData(ctx, req.(*DataFile))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_GetSum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Unit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).GetSum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rannu.Worker/GetSum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).GetSum(ctx, req.(*Unit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_GetScatterMatrix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Vector)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).GetScatterMatrix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rannu.Worker/GetScatterMatrix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).GetScatterMatrix(ctx, req.(*Vector))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_ComputeScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Matrix)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).ComputeScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rannu.Worker/ComputeScores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).ComputeScores(ctx, req.(*Matrix))
	}
	return interceptor(ctx, in, info, handler)
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rannu.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoadData",
			Handler:    _Worker_LoadData_Handler,
		},
		{
			MethodName: "GetSum",
			Handler:    _Worker_GetSum_Handler,
		},
		{
			MethodName: "GetScatterMatrix",
			Handler:    _Worker_GetScatterMatrix_Handler,
		},
		{
			MethodName: "ComputeScores",
			Handler:    _Worker_ComputeScores_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("rannu.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x91, 0x41, 0x4f, 0x83, 0x40,
	0x10, 0x85, 0x4b, 0xa5, 0x9b, 0x3a, 0xa4, 0xb1, 0x99, 0x53, 0xc3, 0xc1, 0x98, 0x3d, 0x18, 0xf4,
	0x40, 0x22, 0xfd, 0x07, 0x6a, 0xf4, 0xa2, 0x17, 0x88, 0x7a, 0x5e, 0x71, 0x0e, 0x44, 0xd8, 0x6d,
	0x96, 0x21, 0x1a, 0xff, 0x9e, 0x7f, 0x4c, 0x76, 0xa1, 0x35, 0xf4, 0x36, 0xbc, 0xf7, 0xbd, 0xe1,
	0x4d, 0x16, 0x22, 0xab, 0xb4, 0xee, 0xd2, 0x9d, 0x35, 0x6c, 0x70, 0xe1, 0x3f, 0xa4, 0x80, 0xf0,
	0x45, 0x57, 0x2c, 0xcf, 0x61, 0x79, 0xaf, 0x58, 0x3d, 0x54, 0x35, 0x21, 0x42, 0xa8, 0x55, 0x43,
	0x9b, 0xe0, 0x22, 0x48, 0x4e, 0x73, 0x3f, 0xcb, 0x14, 0xc2, 0xa2, 0xfa, 0xf1, 0x9e, 0x35, 0x5f,
	0xad, 0xf7, 0x16, 0xb9, 0x9f, 0x9d, 0x56, 0x9a, 0xba, 0xdd, 0xcc, 0x07, 0xcd, 0xcd, 0x32, 0x01,
	0xf1, 0x4a, 0x25, 0x1b, 0x8b, 0xfd, 0x66, 0xaa, 0xa9, 0x21, 0xcd, 0x2e, 0x75, 0x92, 0x04, 0xb7,
	0xf3, 0x75, 0x90, 0x1f, 0x34, 0xb9, 0x05, 0xf1, 0xac, 0xd8, 0x56, 0xdf, 0x78, 0x75, 0x44, 0x46,
	0xd9, 0x2a, 0x1d, 0x2a, 0x0f, 0xab, 0xfe, 0x43, 0xd9, 0x6f, 0x00, 0xe2, 0xcd, 0xd8, 0x4f, 0xb2,
	0x78, 0x0d, 0xcb, 0x27, 0xa3, 0x3e, 0x5c, 0x7b, 0x3c, 0x1b, 0xf9, 0xfd, 0x29, 0x71, 0x34, 0x0a,
	0xae, 0xbb, 0x9c, 0xe1, 0x25, 0x88, 0x47, 0xe2, 0xa2, 0x6b, 0x70, 0x6f, 0xb8, 0xe3, 0xe3, 0xe9,
	0x6f, 0x7a, 0x2e, 0x83, 0xb5, 0xe3, 0x4a, 0xc5, 0x4c, 0x76, 0x6c, 0x37, 0x85, 0x0e, 0x99, 0xc1,
	0xed, 0x33, 0x37, 0xb0, 0xba, 0x33, 0xcd, 0xae, 0x63, 0x2a, 0x4a, 0x63, 0xa9, 0xc5, 0x29, 0x11,
	0x1f, 0x77, 0x93, 0xb3, 0x77, 0xe1, 0x9f, 0x62, 0xfb, 0x17, 0x00, 0x00, 0xff, 0xff, 0x68, 0xf9,
	0xa8, 0xd7, 0x99, 0x01, 0x00, 0x00,
}
