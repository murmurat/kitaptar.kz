package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/murat96k/kitaptar.kz/api"
	"github.com/murat96k/kitaptar.kz/internal/kitaptar/entity"
)

func (p *Postgres) GetAllAuthors(ctx context.Context) ([]entity.Author, error) {

	var authors []entity.Author

	query := fmt.Sprintf("SELECT id, firstname, lastname, image_path, about_author FROM %s", authorTable)

	rows, err := p.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		author := entity.Author{}
		err = rows.Scan(&author.Id, &author.Firstname, &author.Lastname, &author.ImagePath, &author.AboutAuthor)
		authors = append(authors, author)
		if err != nil {
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

	query := fmt.Sprintf("SELECT id, firstname, lastname, image_path, about_author FROM %s WHERE id=$1", authorTable)

	err := pgxscan.Get(ctx, p.Pool, author, query, id)
	if err != nil {
		return nil, err
	}

	return author, nil
}

func (p *Postgres) GetAuthorByName(ctx context.Context, name string) ([]entity.Author, error) {

	var authors []entity.Author

	query := fmt.Sprintf("SELECT firstname, lastname, image_path, about_author FROM %s WHERE firstname LIKE $1 OR lastname LIKE $1 OR CONCAT(firstname, ' ', lastname) LIKE $1;", authorTable)

	rows, err := p.Pool.Query(ctx, query, fmt.Sprint("%", name, "%"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		author := entity.Author{}
		err = rows.Scan(&author.Firstname, &author.Lastname, &author.ImagePath, &author.AboutAuthor)
		authors = append(authors, author)
		if err != nil {
			return nil, err
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return authors, nil
}

func (p *Postgres) CreateAuthor(ctx context.Context, req *api.AuthorRequest) (string, error) {

	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return "", err
	}

	var authorId string
	query := fmt.Sprintf(`
			INSERT INTO %s (firstname, lastname, image_path, about_author, created_at)
			VALUES ($1, $2, $3, $4, $5) RETURNING id
			`, authorTable)

	err = p.Pool.QueryRow(ctx, query, req.Firstname, req.Lastname, req.ImagePath, req.AboutAuthor, time.Now()).Scan(&authorId)
	if err != nil {
		//nolint
		tx.Rollback(ctx)
		return "", err
	}

	return authorId, tx.Commit(ctx)
}

func (p *Postgres) DeleteAuthor(ctx context.Context, id string) error {

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", authorTable)

	_, err := p.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) UpdateAuthor(ctx context.Context, id string, req *api.AuthorRequest) error {

	values := make([]string, 0)
	paramCount := 2
	params := make([]interface{}, 0)

	if req.Firstname != "" {
		values = append(values, fmt.Sprintf("firstname=$%d", paramCount))
		params = append(params, req.Firstname)
		paramCount++
	}
	if req.Lastname != "" {
		values = append(values, fmt.Sprintf("lastname=$%d", paramCount))
		params = append(params, req.Lastname)
		paramCount++
	}
	if req.AboutAuthor != "" {
		values = append(values, fmt.Sprintf("about_author=$%d", paramCount))
		params = append(params, req.AboutAuthor)
		paramCount++
	}
	if req.ImagePath != "" {
		values = append(values, fmt.Sprintf("image_path=$%d", paramCount))
		params = append(params, req.ImagePath)
	}

	setQuery := strings.Join(values, ", ")
	setQuery = fmt.Sprintf("UPDATE %s SET ", authorTable) + setQuery + " WHERE id=$1"

	params = append([]interface{}{id}, params...)

	_, err := p.Pool.Exec(ctx, setQuery, params...)
	if err != nil {
		return err
	}

	return nil
}
