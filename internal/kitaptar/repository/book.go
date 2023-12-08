package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/murat96k/kitaptar.kz/api"
	"github.com/murat96k/kitaptar.kz/internal/kitaptar/entity"
)

func (p *Postgres) GetUserBooks(email string) ([]entity.Book, error) {

	var books []entity.Book

	// TODO Implement all books which liked by user(Saved books)

	return books, nil
}

func (p *Postgres) GetAllBooks(ctx context.Context, sortBy string) ([]entity.Book, error) {

	var books []entity.Book

	query := fmt.Sprintf("SELECT id, name, genre, annotation, author_id, image_path, file_path_id, created_at FROM %s", bookTable)

	if sortBy == "asc" {
		query += " ORDER BY created_at ASC"
	} else if sortBy == "desc" {
		query += " ORDER BY created_at DESC"
	}

	rows, err := p.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := entity.Book{}
		err = rows.Scan(&book.Id, &book.Name, &book.Genre, &book.Annotation, &book.AuthorId, &book.ImagePath, &book.FilePathId, &book.CreatedAt)
		books = append(books, book)
		if err != nil {
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

	query := fmt.Sprintf("SELECT id, name, genre, annotation, author_id, image_path, file_path_id FROM %s WHERE id=$1", bookTable)

	err := pgxscan.Get(ctx, p.Pool, book, query, strings.TrimSpace(id))
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (p *Postgres) GetBookByName(ctx context.Context, name string) ([]entity.Book, error) {

	var books []entity.Book

	query := fmt.Sprintf("SELECT name, genre, annotation, author_id, image_path, file_path_id FROM %s WHERE name LIKE $1 OR annotation LIKE $1 OR CONCAT(name, ' ', annotation) LIKE $1;", bookTable)

	rows, err := p.Pool.Query(ctx, query, fmt.Sprint("%", name, "%"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := entity.Book{}
		err = rows.Scan(&book.Name, &book.Genre, &book.Annotation, &book.AuthorId, &book.ImagePath, &book.FilePathId)
		books = append(books, book)
		if err != nil {
			return nil, err
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (p *Postgres) CreateBook(ctx context.Context, req *api.BookRequest) (string, error) {

	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return "", err
	}

	var bookId string
	query := fmt.Sprintf(`
			INSERT INTO %s (
			                author_id, 
			                annotation,
			                name,
			                genre,
							image_path, 
							file_path_id,
							created_at
			                )
			VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
			`, bookTable)

	err = p.Pool.QueryRow(ctx, query, req.AuthorId, req.Annotation, req.Name, req.Genre, req.ImagePath, req.FilePathId, time.Now()).Scan(&bookId)
	if err != nil {
		//nolint
		tx.Rollback(ctx)
		return "", err
	}

	return bookId, tx.Commit(ctx)
}

func (p *Postgres) DeleteBook(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", bookTable)

	_, err := p.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) UpdateBook(ctx context.Context, id string, req *api.BookRequest) error {

	values := make([]string, 0)
	paramCount := 2
	params := make([]interface{}, 0)

	if req.Name != "" {
		values = append(values, fmt.Sprintf("name=$%d", paramCount))
		params = append(params, req.Name)
		paramCount++
	}
	if req.Annotation != "" {
		values = append(values, fmt.Sprintf("annotation=$%d", paramCount))
		params = append(params, req.Annotation)
		paramCount++
	}
	if req.Genre != "" {
		values = append(values, fmt.Sprintf("genre=$%d", paramCount))
		params = append(params, req.Genre)
		paramCount++
	}
	if req.AuthorId != uuid.Nil {
		values = append(values, fmt.Sprintf("author_id=$%d", paramCount))
		params = append(params, req.AuthorId)
		paramCount++
	}
	if req.FilePathId != uuid.Nil {
		values = append(values, fmt.Sprintf("file_path_id=$%d", paramCount))
		params = append(params, req.FilePathId)
		paramCount++
	}
	if req.ImagePath != "" {
		values = append(values, fmt.Sprintf("image_path=$%d", paramCount))
		params = append(params, req.ImagePath)
	}

	setQuery := strings.Join(values, ", ")
	setQuery = fmt.Sprintf("UPDATE %s SET ", bookTable) + setQuery + " WHERE id=$1"

	params = append([]interface{}{id}, params...)

	_, err := p.Pool.Exec(ctx, setQuery, params...)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) AddToFavorites(ctx context.Context, userId, bookId string) (string, error) {

	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return "", err
	}

	var favoriteId string
	query := fmt.Sprintf("INSERT INTO %s (user_id, book_id) VALUES ($1, $2) RETURNING id", favoritesTable)

	err = p.Pool.QueryRow(ctx, query, userId, bookId).Scan(&favoriteId)
	if err != nil {
		//nolint
		tx.Rollback(ctx)
		return "", err
	}

	return favoriteId, tx.Commit(ctx)
}

func (p *Postgres) DeleteFromFavorites(ctx context.Context, favoriteId string) error {

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", favoritesTable)

	_, err := p.Pool.Exec(ctx, query, favoriteId)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetFromFavorites(ctx context.Context, userId, bookId string) (*entity.FavoriteBook, error) {

	favorite := new(entity.FavoriteBook)

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 AND book_id=$2", favoritesTable)

	err := pgxscan.Get(ctx, p.Pool, favorite, query, strings.TrimSpace(userId), strings.TrimSpace(bookId))
	if err != nil {
		return nil, err
	}

	return favorite, nil
}
