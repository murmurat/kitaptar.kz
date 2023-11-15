package pgrepo

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
)

import "github.com/jackc/pgx/v4/pgxpool"

const (
	usersTable     = "users"
	bookTable      = "books"
	authorTable    = "authors"
	filePathsTable = "file_paths"
)

type Postgres struct {
	host     string
	username string
	password string
	port     string
	dbName   string
	Pool     *pgxpool.Pool
	SQLDB    *sql.DB
}

func New(opts ...Option) (*Postgres, error) {
	p := new(Postgres)

	for _, opt := range opts {
		opt(p)
	}

	u := url.URL{
		Scheme: "postgresql",
		User:   url.UserPassword(p.username, p.password),
		Host:   fmt.Sprintf("%s:%s", p.host, p.port),
		Path:   p.dbName,
	}

	poolConfig, err := pgxpool.ParseConfig(u.String())
	if err != nil {
		return nil, fmt.Errorf("pgxpool config error %s", err.Error())
	}
	p.Pool, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("pgxpool connect err: %s", err)
	}
	return p, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
