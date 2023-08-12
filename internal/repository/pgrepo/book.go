package pgrepo

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"log"
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

	query := fmt.Sprintf("SELECT id, name,genre, annotation ,author_id, image_path FROM %s WHERE id=$1", bookTable)
	err := pgxscan.Get(ctx, p.Pool, book, query, strings.TrimSpace(id))
	if err != nil {
		return nil, err
	}

	return book, nil
}
