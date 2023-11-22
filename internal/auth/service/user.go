package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/murat96k/kitaptar.kz/internal/auth/entity"
	"github.com/murat96k/kitaptar.kz/pkg/util"
	"github.com/uristemov/auth-user-grpc/models"
)

func (m *Manager) CreateUser(ctx context.Context, u *entity.User) (string, error) {

	hashPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return "", fmt.Errorf("hash password error %w", err)
	}

	u.Password = hashPassword

	req := &models.RegisterUser{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Password:  hashPassword,
		Email:     u.Email,
	}

	return m.UserClient.CreateUser(ctx, req)
}

func (m *Manager) Login(ctx context.Context, email, password string) (string, error) {

	user, err := m.UserClient.GetUserByEmail(ctx, email)
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

	accessToken, err := m.Token.CreateToken(user.Id, user.Email, m.Config.Auth.TimeToLive)
	if err != nil {
		return "", fmt.Errorf("create token err: %w", err)
	}

	return accessToken, nil
}
