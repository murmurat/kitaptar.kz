package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/murat96k/kitaptar.kz/api"
	"github.com/murat96k/kitaptar.kz/internal/entity"
	"github.com/murat96k/kitaptar.kz/pkg/util"
)

func (m *Manager) CreateUser(ctx context.Context, u *entity.User) (string, error) {
	hashPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return "", fmt.Errorf("hash password error %w", err)
	}

	u.Password = hashPassword

	return m.Repository.CreateUser(ctx, u)
}

func (m *Manager) Login(ctx context.Context, email, password string) (string, error) {
	user, err := m.Repository.GetUser(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("user not found")
		}

		return "", fmt.Errorf("get user err: %w", err)
	}

	err = util.CheckPassword(password, user.Password)
	if err != nil {
		return "", fmt.Errorf("incorrect password: %w", err)
	}

	accessToken, err := m.Token.CreateToken(user.Id.String(), user.Email, m.Config.Token.TimeToLive)
	if err != nil {
		return "", fmt.Errorf("create token err: %w", err)
	}

	return accessToken, nil
}

func (m *Manager) VerifyToken(token string) (string, error) {
	claim, err := m.Token.ValidateToken(token)
	if err != nil {
		return "", fmt.Errorf("validate token err: %w", err)
	}

	return claim.UserID, nil
}

func (m *Manager) UpdateUser(ctx context.Context, id string, req *api.UpdateUserRequest) error {

	//user, err := m.Repository.GetUser(ctx, id)
	//userID := user.Id
	//if err != nil {
	//	return err
	//}
	return m.Repository.UpdateUser(ctx, id, req)
}

func (m *Manager) GetUser(ctx context.Context, id string) (*entity.User, error) {
	return m.Repository.GetUser(ctx, id)
}

func (m *Manager) DeleteUser(ctx context.Context, id string) error {
	return m.Repository.DeleteUser(ctx, id)
}
