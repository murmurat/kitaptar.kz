package grpc

import (
	"fmt"
	"net"

	v1 "github.com/murat96k/kitaptar.kz/internal/user/handler/grpc/v1"
	"github.com/uristemov/auth-user-grpc/protobuf"
	"google.golang.org/grpc"
)

type Server struct {
	port       string
	grpcServer *grpc.Server
	service    *v1.Service
}

func NewServer(port string, service *v1.Service) *Server {
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	return &Server{
		port:       port,
		service:    service,
		grpcServer: grpcServer,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.port)
	if err != nil {
		return fmt.Errorf("failed to listen grpc port: %s", s.port)
	}

	protobuf.RegisterUserServer(s.grpcServer, s.service)

	//nolint
	go s.grpcServer.Serve(listener)

	return nil
}

func (s *Server) Close() {
	s.grpcServer.GracefulStop()
}
