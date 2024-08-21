package api

import (
	"google.golang.org/grpc"
)

type GrpcServer struct {
	server *grpc.Server
}

type GrpcService interface {
	ServiceDesc() *grpc.ServiceDesc
}

func NewGrpcServer(services ...GrpcService) *GrpcServer {
	server := grpc.NewServer()

	return &GrpcServer{server: server}
}
