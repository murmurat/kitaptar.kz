package service

import (
	"context"
	"github.com/murat96k/kitaptar.kz/internal/auth/entity"
)

//go:generate mockgen -source=service.go -destination=mock/mock_service.go
type Service interface {
	CreateUser(ctx context.Context, u *entity.User) (string, error)
	Login(ctx context.Context, email, password string) (string, error)
}
