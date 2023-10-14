package pgrepo

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/murat96k/kitaptar.kz/api"
	"github.com/murat96k/kitaptar.kz/internal/entity"
	"github.com/murat96k/kitaptar.kz/pkg/util"
)

func (p *Postgres) CreateUser(ctx context.Context, u *entity.User) (string, error) {
	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return "", err
	}

	var userId string
	query := fmt.Sprintf(`
			INSERT INTO %s (
			                email, -- 1 
			                first_name, -- 2
			                last_name, -- 3
			                password, -- 4,
							created_at
			                )
			VALUES ($1, $2, $3, $4, $5) RETURNING id
			`, usersTable)

	err = p.Pool.QueryRow(ctx, query, u.Email, u.FirstName, u.LastName, u.Password, time.Now()).Scan(&userId)
	if err != nil {
		tx.Rollback(ctx)
		return "", err
	}

	return userId, tx.Commit(ctx)
}

func (p *Postgres) GetUser(ctx context.Context, email string) (*entity.User, error) {
	user := new(entity.User)
	query := fmt.Sprintf("SELECT id, email, first_name, last_name, password FROM %s WHERE email = '%s'", usersTable, email)

	err := pgxscan.Get(ctx, p.Pool, user, query)
	if err != nil {
		log.Println("Error after pgx get")
		return nil, err
	}

	return user, nil
}

func (p *Postgres) UpdateUser(ctx context.Context, id string, user *api.UpdateUserRequest) error {

	values := make([]string, 0)

	if user.FirstName != "" {
		values = append(values, fmt.Sprintf("first_name='%s'", user.FirstName))
	}
	if user.LastName != "" {
		values = append(values, fmt.Sprintf("last_name='%s'", user.LastName))
	}
	if user.Email != "" {
		values = append(values, fmt.Sprintf("email='%s'", user.Email))
	}
	if user.Password != "" {
		password, err := util.HashPassword(user.Password)
		if err != nil {
			//fmt.Println(err)
			return err
		}
		values = append(values, fmt.Sprintf("password='%s'", password))
	}

	setQuery := strings.Join(values, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = '%s';", usersTable, setQuery, id)
	fmt.Println(query)

	_, err := p.Pool.Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) DeleteUser(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id='%s'", usersTable, id)

	_, err := p.Pool.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
