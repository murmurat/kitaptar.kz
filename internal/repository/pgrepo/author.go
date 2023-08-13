package pgrepo

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"log"
	"one-lab/api"
	"one-lab/internal/entity"
	"strings"
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

	query := fmt.Sprintf("SELECT id, firstname,lastname, image_path ,about_author FROM %s WHERE id='$1'", authorTable)
	err := pgxscan.Get(ctx, p.Pool, author, query, strings.TrimSpace(id))
	if err != nil {
		return nil, err
	}

	return author, nil
}

func (p *Postgres) CreateAuthor(ctx context.Context, req *api.AuthorRequest) error {
	query := fmt.Sprintf(`
			INSERT INTO %s (firstname,lastname, image_path ,about_author)
			VALUES ($1, $2, $3, $4)
			`, authorTable)

	fmt.Println(req)
	_, err := p.Pool.Exec(ctx, query, req.Firstname, req.Lastname, req.ImagePath, req.AboutAuthor)
	if err != nil {
		return err
	}

	return nil
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

	//fmt.Printf("Error dont have before query %s, query: '%s'", user.Password, setQuery)
	//query := fmt.Sprintf("UPDATE %s SET %s WHERE email = %s;", usersTable, setQuery, email)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = '%s';", authorTable, setQuery, id)
	fmt.Println(query)

	_, err := p.Pool.Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}
