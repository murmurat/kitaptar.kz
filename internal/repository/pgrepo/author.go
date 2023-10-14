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
)

func (p *Postgres) GetAllAuthors(ctx context.Context) ([]entity.Author, error) {
	var authors []entity.Author
	query := fmt.Sprintf("SELECT id, firstname,lastname, image_path ,about_author FROM %s", authorTable)
	rows, err := p.Pool.Query(ctx, query)
	//rows, err := p.SQLDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		author := entity.Author{}
		err = rows.Scan(&author.Id, &author.Firstname, &author.Lastname, &author.ImagePath, &author.AboutAuthor)
		authors = append(authors, author)
		if err != nil {
			log.Printf("Scan author values error %s", err.Error())
			return nil, err
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (p *Postgres) GetAuthorById(ctx context.Context, id string) (*entity.Author, error) {

	author := new(entity.Author)

	query := fmt.Sprintf("SELECT id, firstname,lastname, image_path ,about_author FROM %s WHERE id='%s'", authorTable, strings.TrimSpace(id))
	err := pgxscan.Get(ctx, p.Pool, author, query)
	if err != nil {
		return nil, err
	}

	return author, nil
}

func (p *Postgres) CreateAuthor(ctx context.Context, req *api.AuthorRequest) (string, error) {
	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return "", err
	}

	var authorId string
	// Need to flexible code (Like Update queries)
	query := fmt.Sprintf(`
			INSERT INTO %s (firstname, lastname, image_path ,about_author, created_at)
			VALUES ($1, $2, $3, $4, $5) RETURNING id
			`, authorTable)

	err = p.Pool.QueryRow(ctx, query, req.Firstname, req.Lastname, req.ImagePath, req.AboutAuthor, time.Now()).Scan(&authorId)
	if err != nil {
		tx.Rollback(ctx)
		return "", err
	}

	return authorId, tx.Commit(ctx)
}

func (p *Postgres) DeleteAuthor(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id='%s'", authorTable, id)

	_, err := p.Pool.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) UpdateAuthor(ctx context.Context, id string, req *api.AuthorRequest) error {
	values := make([]string, 0)

	if req.Firstname != "" {
		values = append(values, fmt.Sprintf("firstname='%s'", req.Firstname))
	}
	if req.Lastname != "" {
		values = append(values, fmt.Sprintf("lastname='%s'", req.Lastname))
	}
	if req.AboutAuthor != "" {
		// check for existing author
		values = append(values, fmt.Sprintf("about_author='%s'", req.AboutAuthor))
	}
	if req.ImagePath != "" {
		values = append(values, fmt.Sprintf("image_path='%s'", req.ImagePath))
	}

	setQuery := strings.Join(values, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = '%s';", authorTable, setQuery, id)

	_, err := p.Pool.Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}
