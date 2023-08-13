package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"one-lab/api"
	"one-lab/internal/entity"
	"one-lab/pkg/util"
)

func (m *Manager) CreateUser(ctx context.Context, u *entity.User) error {
	hashPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return fmt.Errorf("hash password error %w", err)
	}

	u.Password = hashPassword
	err = m.Repository.CreateUser(ctx, u)
	if err != nil {
		return err
	}

	return nil
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

	accessToken, err := m.Token.CreateToken(user.Email, m.Config.Token.TimeToLive)
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

	return claim.Email, nil
}

func (m *Manager) UpdateUser(ctx context.Context, email string, req *api.UpdateUserRequest) error {

	user, err := m.Repository.GetUser(ctx, email)
	userID := user.Id
	if err != nil {
		return err
	}
	return m.Repository.UpdateUser(ctx, userID.String(), req)
}

func (m *Manager) GetUser(ctx context.Context, email string) (*entity.User, error) {
	return m.Repository.GetUser(ctx, email)
}

func (m *Manager) DeleteUser(ctx context.Context, email string) error {
	return m.Repository.DeleteUser(ctx, email)
}
