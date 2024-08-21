package api

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/isnastish/chat-backend/pkg/log"
)

type GrpcServer struct {
	server *grpc.Server
}

type GrpcService interface {
	ServiceDesc() *grpc.ServiceDesc
}

func NewGrpcServer(services ...GrpcService) *GrpcServer {
	server := grpc.NewServer()

	for _, service := range services {
		desc := service.ServiceDesc()
		server.RegisterService(desc, service)
	}

	return &GrpcServer{server: server}
}

func (s *GrpcServer) Serve(port uint) error {
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", port))
	if err != nil {
		return fmt.Errorf("Failed to listen on port %d, error %v", port, err)
	}

	log.Logger.Info("grpc: Listening on port %v", fmt.Sprintf("0.0.0.0:%v", port))

	err = s.server.Serve(listener)
	if err != nil {
		fmt.Errorf("grpc: Faile to serve %v", err)
	}

	return nil
}

func (s *GrpcServer) Close() {
	s.server.Stop()
}
