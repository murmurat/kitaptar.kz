package v1

import (
	"context"
	"github.com/murat96k/kitaptar.kz/internal/user/entity"
	"github.com/murat96k/kitaptar.kz/internal/user/service"
	"github.com/uristemov/auth-user-grpc/protobuf"
	"log"
)

type Service struct {
	protobuf.UnimplementedUserServer
	service service.Service
}

func NewService(service service.Service) *Service {
	return &Service{
		service: service,
	}
}

func (s *Service) GetUserByEmail(ctx context.Context, req *protobuf.GetUserByEmailRequest) (*protobuf.GetUserByEmailResponse, error) {
	user, err := s.service.GetUserByEmail(ctx, req.Email)
	if err != nil {
		log.Printf("failed to GetUserByEmail err %v", err)
		return nil, err
	}

	return &protobuf.GetUserByEmailResponse{
		Id:        user.Id.String(),
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Password:  user.Password,
		Email:     user.Email,
	}, nil
}

func (s *Service) CreateUser(ctx context.Context, req *protobuf.CreateUserRequest) (*protobuf.CreateUserResponse, error) {

	user := &entity.User{
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Password:  req.Password,
		Email:     req.Email,
	}

	id, err := s.service.CreateUser(ctx, user)
	if err != nil {
		log.Printf("failed to CreateUser %v", err)
		return nil, err
	}

	return &protobuf.CreateUserResponse{
		Id:    id,
		Error: nil,
	}, nil
}
