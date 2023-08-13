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

func (p *Postgres) GetUserBooks(email string) ([]entity.Book, error) {
	var books []entity.Book

	// Implement all books which liked by user(Saved books)

	//query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
	//	todoListsTable, usersListsTable)
	//err := r.db.Select(&lists, query, userId)

	return books, nil
}

func (p *Postgres) GetAllBooks(ctx context.Context) ([]entity.Book, error) {
	var books []entity.Book
	query := fmt.Sprintf("SELECT id, name,genre, annotation ,author_id, image_path FROM %s", bookTable)
	rows, err := p.Pool.Query(ctx, query)
	//rows, err := p.SQLDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		book := entity.Book{}
		err = rows.Scan(&book.Id, &book.Name, &book.Genre, &book.Annotation, &book.AuthorId, &book.ImagePath)
		books = append(books, book)
		if err != nil {
			log.Printf("Scan book values error %s", err.Error())
			return nil, err
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (p *Postgres) GetBookById(ctx context.Context, id string) (*entity.Book, error) {

	book := new(entity.Book)

	query := fmt.Sprintf("SELECT id, name,genre, annotation ,author_id, image_path FROM %s WHERE id='$1'", bookTable)
	err := pgxscan.Get(ctx, p.Pool, book, query, strings.TrimSpace(id))
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (p *Postgres) CreateBook(ctx context.Context, req *api.BookRequest) error {
	query := fmt.Sprintf(`
			INSERT INTO %s (
			                author_id, -- 1 
			                annotation, -- 2
			                name, -- 3
			                genre, -- 4
							image_path
			                )
			VALUES ($1, $2, $3, $4, $5)
			`, bookTable)

	fmt.Println(req)
	_, err := p.Pool.Exec(ctx, query, req.AuthorId, req.Annotation, req.Name, req.Genre, req.ImagePath)
	if err != nil {
		return err
	}

	return nil
}
func (p *Postgres) DeleteBook(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id='%s'", bookTable, id)

	_, err := p.Pool.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) UpdateBook(ctx context.Context, id string, req *api.BookRequest) error {
	values := make([]string, 0)

	if req.Name != "" {
		values = append(values, fmt.Sprintf("name='%s'", req.Name))
	}
	if req.Annotation != "" {
		values = append(values, fmt.Sprintf("annotation='%s'", req.Annotation))
	}
	if req.AuthorId.String() != "" {
		// check for existing author
		values = append(values, fmt.Sprintf("author_id='%s'", req.AuthorId))
	}
	if req.ImagePath != "" {
		values = append(values, fmt.Sprintf("image_path='%s'", req.ImagePath))
	}

	setQuery := strings.Join(values, ", ")

	//fmt.Printf("Error dont have before query %s, query: '%s'", user.Password, setQuery)
	//query := fmt.Sprintf("UPDATE %s SET %s WHERE email = %s;", usersTable, setQuery, email)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = '%s';", bookTable, setQuery, id)
	fmt.Println(query)

	_, err := p.Pool.Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}
