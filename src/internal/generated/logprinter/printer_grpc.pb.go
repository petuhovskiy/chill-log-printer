// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: public/printer.proto

package logprinter

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

// PrinterServiceClient is the client API for PrinterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrinterServiceClient interface {
	PrintLog(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error)
}

type printerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrinterServiceClient(cc grpc.ClientConnInterface) PrinterServiceClient {
	return &printerServiceClient{cc}
}

func (c *printerServiceClient) PrintLog(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := c.cc.Invoke(ctx, "/PrinterService/PrintLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrinterServiceServer is the server API for PrinterService service.
// All implementations must embed UnimplementedPrinterServiceServer
// for forward compatibility
type PrinterServiceServer interface {
	PrintLog(context.Context, *LogRequest) (*LogResponse, error)
	mustEmbedUnimplementedPrinterServiceServer()
}

// UnimplementedPrinterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPrinterServiceServer struct {
}

func (UnimplementedPrinterServiceServer) PrintLog(context.Context, *LogRequest) (*LogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrintLog not implemented")
}
func (UnimplementedPrinterServiceServer) mustEmbedUnimplementedPrinterServiceServer() {}

// UnsafePrinterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrinterServiceServer will
// result in compilation errors.
type UnsafePrinterServiceServer interface {
	mustEmbedUnimplementedPrinterServiceServer()
}

func RegisterPrinterServiceServer(s grpc.ServiceRegistrar, srv PrinterServiceServer) {
	s.RegisterService(&PrinterService_ServiceDesc, srv)
}

func _PrinterService_PrintLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrinterServiceServer).PrintLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PrinterService/PrintLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrinterServiceServer).PrintLog(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PrinterService_ServiceDesc is the grpc.ServiceDesc for PrinterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrinterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PrinterService",
	HandlerType: (*PrinterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PrintLog",
			Handler:    _PrinterService_PrintLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "public/printer.proto",
}