package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/murat96k/kitaptar.kz/internal/auth/entity"
	"github.com/murat96k/kitaptar.kz/internal/auth/metrics"
	"github.com/murat96k/kitaptar.kz/pkg/util"
	"github.com/uristemov/auth-user-grpc/models"
)

func (m *Manager) CreateUser(ctx context.Context, u *entity.User) (string, error) {

	metrics.Auth.Add(1)

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

func (m *Manager) Login(ctx context.Context, email, password string) (string, string, error) {

	user, err := m.UserClient.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", fmt.Errorf("user not found")
		}

		return "", "", fmt.Errorf("get user err: %w", err)
	}

	err = util.CheckPassword(password, user.Password)
	if err != nil {
		return "", "", fmt.Errorf("incorrect password: %w", err)
	}

	accessToken, err := m.Token.CreateToken(user.Id, user.Email, m.Config.Auth.Access.TimeToLive)
	if err != nil {
		return "", "", fmt.Errorf("create access token err: %w", err)
	}

	refreshTokenClaims, err := m.Token.CreateRefreshToken(user.Id, user.Email, m.Config.Auth.Refresh.TimeToLive)
	if err != nil {
		return "", "", fmt.Errorf("create refresh token err: %w", err)
	}

	err = m.Cache.TokenCache.SetToken(ctx, user.Id, refreshTokenClaims.RefreshId, m.Config.Auth.Refresh.TimeToLive)
	if err != nil {
		return "", "", err
	}

	fmt.Println("Refresh token claims: ", refreshTokenClaims.RefreshId)

	return accessToken, refreshTokenClaims.TokenString, nil
}

func (m *Manager) Refresh(ctx context.Context, refreshToken string) (string, string, error) {

	claim, err := m.Token.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", "", fmt.Errorf("validate token err: %w", err)
	}

	fmt.Println("Refresh token claims: ", claim.RefreshId)

	refreshId, err := m.Cache.TokenCache.GetToken(ctx, claim.UserID)
	if err != nil {
		return "", "", err
	}

	if refreshId == "" {
		return "", "", nil
	}

	// TODO check user with user_id for existence

	accessToken, err := m.Token.CreateToken(claim.UserID, claim.Email, m.Config.Auth.Access.TimeToLive)
	if err != nil {
		return "", "", fmt.Errorf("create access token err: %w", err)
	}

	refreshClaims, err := m.Token.CreateRefreshToken(claim.UserID, claim.Email, m.Config.Auth.Refresh.TimeToLive)
	if err != nil {
		return "", "", fmt.Errorf("create refresh token err: %w", err)
	}

	err = m.Cache.TokenCache.SetToken(ctx, claim.UserID, refreshClaims.RefreshId, m.Config.Auth.Refresh.TimeToLive)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshClaims.TokenString, nil
}
