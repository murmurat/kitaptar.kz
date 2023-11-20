package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/murat96k/kitaptar.kz/api"
	"github.com/murat96k/kitaptar.kz/internal/entity"
	"github.com/murat96k/kitaptar.kz/pkg/util"
	"github.com/redis/go-redis/v9"
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

	user, err := m.GetUserByEmail(ctx, email)
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

	user, err := m.Repository.GetUserById(ctx, id)
	if err != nil {
		return err
	}

	if req.Email != "" {
		user.Email = req.Email
	}
	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Password != "" {
		req.Password, err = util.HashPassword(req.Password)
		if err != nil {
			return err
		}
		user.Password = req.Password
	}

	err = m.Cache.UserCache.DeleteUser(ctx, user.Id.String())
	if err != nil {
		return err
	}

	err = m.Repository.UpdateUser(ctx, id, req)
	if err != nil {
		return err
	}

	_ = m.Cache.UserCache.SetUser(ctx, user)

	return nil
}

func (m *Manager) GetUserById(ctx context.Context, id string) (*entity.User, error) {

	user, err := m.Cache.UserCache.GetUser(ctx, id)
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}

	user, err = m.Repository.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	err = m.Cache.UserCache.SetUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *Manager) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {

	user, err := m.Repository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	err = m.Cache.UserCache.SetUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *Manager) DeleteUser(ctx context.Context, id string) error {

	err := m.Cache.UserCache.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return m.Repository.DeleteUser(ctx, id)
}
