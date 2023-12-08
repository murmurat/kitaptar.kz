package repository

import (
	"context"

	"github.com/murat96k/kitaptar.kz/api"
	"github.com/murat96k/kitaptar.kz/internal/user/entity"
)

type Repository interface {
	CreateUser(ctx context.Context, u *entity.User) (string, error)
	GetUserById(ctx context.Context, id string) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	UpdateUser(ctx context.Context, id string, req *api.UpdateUserRequest) error
	DeleteUser(ctx context.Context, id string) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)
	SetUserRoleById(ctx context.Context, userID, role string) error
}
