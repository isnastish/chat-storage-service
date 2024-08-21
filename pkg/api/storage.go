package api

import (
	"github.com/isnastish/chat-backend/pkg/service"
	"github.com/isnastish/chat-backend/proto/api"
	"google.golang.org/grpc"
)

type storageServer struct {
	api.UnimplementedChatStorageServiceServer
	service     *service.Service
	serviceDesc *grpc.ServiceDesc
}

func NewStorageServer(service *service.Service) *storageServer {
	return &storageServer{
		service:     service,
		serviceDesc: &api.ChatStorageService_ServiceDesc,
	}
}

func (s *storageServer) ServiceDesc() *grpc.ServiceDesc {
	return s.serviceDesc
}
