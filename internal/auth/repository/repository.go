package repository

import "context"

type Repository interface {
	GetRefreshToken(ctx context.Context, refreshId string) (string, error)
	GetRefreshTokenByUserId(ctx context.Context, userId string) (string, error)
	CreateRefreshToken(ctx context.Context, refreshToken, userId, refreshId string) (string, error)
	DeleteRefreshToken(ctx context.Context, refreshToken string) error
	UpdateRefreshToken(ctx context.Context, refreshId, refreshToken string) error
}
