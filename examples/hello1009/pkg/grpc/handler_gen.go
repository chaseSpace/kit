// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	pb "hello1009/pb/gen-go/pb"
	endpoint "hello1009/pkg/endpoint"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	foo            grpc.Handler
	updateUserInfo grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.Hello1009Server {
	return &grpcServer{
		foo:            makeFooHandler(endpoints, options["Foo"]),
		updateUserInfo: makeUpdateUserInfoHandler(endpoints, options["UpdateUserInfo"]),
	}
}