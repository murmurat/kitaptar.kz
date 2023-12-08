package service

import (
	"context"

	"github.com/murat96k/kitaptar.kz/api"
	"github.com/murat96k/kitaptar.kz/internal/user/entity"
)

//go:generate mockgen -source=service.go -destination=mock/mock_service.go
type Service interface {
	CreateUser(ctx context.Context, u *entity.User) (string, error)
	VerifyToken(token string) (string, error)
	UpdateUser(ctx context.Context, id string, req *api.UpdateUserRequest) error
	GetUserById(ctx context.Context, id string) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	DeleteUser(ctx context.Context, id string) error
	ConfirmUser(ctx context.Context, userID, code string) error
	GetAllUsers(ctx context.Context, userID string) ([]entity.User, error)
	SetUserRoleById(ctx context.Context, userID, targetUserId, role string) error
}
