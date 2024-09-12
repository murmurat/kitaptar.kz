package repository

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
)

const refreshTokensTable = "refresh_tokens"

func (p *Postgres) GetRefreshToken(ctx context.Context, refreshId string) (string, error) {

	var dbRefreshToken string

	query := fmt.Sprintf("SELECT refresh_token FROM %s WHERE id=$1", refreshTokensTable)

	err := pgxscan.Get(ctx, p.Pool, &dbRefreshToken, query, refreshId)
	if err != nil {
		return "", err
	}

	return dbRefreshToken, nil
}

func (p *Postgres) CreateRefreshToken(ctx context.Context, refreshToken, userId, refreshId string) (string, error) {

	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return "", err
	}

	var newRefreshToken string
	query := fmt.Sprintf(`
			INSERT INTO %s (
							id,
			                user_id,
			                refresh_token
			                )
			VALUES ($1, $2, $3) RETURNING refresh_token
			`, refreshTokensTable)

	err = p.Pool.QueryRow(ctx, query, refreshId, userId, refreshToken).Scan(&newRefreshToken)
	if err != nil {
		//nolint
		tx.Rollback(ctx)
		return "", err
	}

	return newRefreshToken, tx.Commit(ctx)
}

func (p *Postgres) DeleteRefreshToken(ctx context.Context, refreshToken string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE refresh_token=$1", refreshTokensTable)

	_, err := p.Pool.Exec(ctx, query, refreshToken)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) UpdateRefreshToken(ctx context.Context, refreshId, refreshToken string) error {

	query := fmt.Sprintf("UPDATE %s SET refresh_token=$1 WHERE id=$2", refreshTokensTable)

	_, err := p.Pool.Exec(ctx, query, refreshToken, refreshId)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetRefreshTokenByUserId(ctx context.Context, userId string) (string, error) {

	var dbRefreshToken string

	query := fmt.Sprintf("SELECT refresh_token FROM %s WHERE user_id=$1", refreshTokensTable)

	err := pgxscan.Get(ctx, p.Pool, &dbRefreshToken, query, userId)
	if err != nil {
		return "", err
	}

	return dbRefreshToken, nil
}
