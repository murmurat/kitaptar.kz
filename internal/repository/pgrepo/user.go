package pgrepo

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"one-lab/internal/entity"
	"strings"
)

func (p *Postgres) CreateUser(ctx context.Context, u *entity.User) error {
	query := fmt.Sprintf(`
			INSERT INTO %s (
			                email, -- 1 
			                first_name, -- 2
			                last_name, -- 3
			                password -- 4
			                )
			VALUES ($1, $2, $3, $4)
			`, usersTable)

	fmt.Println(u)
	_, err := p.Pool.Exec(ctx, query, u.Email, u.FirstName, u.LastName, u.Password)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetUser(ctx context.Context, email string) (*entity.User, error) {
	user := new(entity.User)

	query := fmt.Sprintf("SELECT id, email, first_name, last_name, password FROM %s WHERE email = $1", usersTable)

	//rows, err := p.SQLDB.Query(query, username)
	//if err != nil {
	//	return nil, err
	//}
	//defer rows.Close()
	//
	//for rows.Next() {
	//	err := rows.Scan(&user.ID, &user.Username, &user.LastName, &user.LastName, &user.Password)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	//err = rows.Err()
	//if err != nil {
	//	return nil, err
	//}

	err := pgxscan.Get(ctx, p.Pool, user, query, strings.TrimSpace(email))
	if err != nil {
		return nil, err
	}

	return user, nil
}
