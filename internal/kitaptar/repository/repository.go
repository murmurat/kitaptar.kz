package repository

import (
	"context"
	"github.com/murat96k/kitaptar.kz/api"
	"github.com/murat96k/kitaptar.kz/internal/kitaptar/entity"
)

type Repository interface {
	CreateBook(ctx context.Context, req *api.BookRequest) (string, error)
	GetUserBooks(email string) ([]entity.Book, error)
	GetAllBooks(ctx context.Context) ([]entity.Book, error)
	GetBookById(ctx context.Context, id string) (*entity.Book, error)
	GetBookByName(ctx context.Context, name string) ([]entity.Book, error)
	DeleteBook(ctx context.Context, id string) error
	UpdateBook(ctx context.Context, id string, req *api.BookRequest) error

	CreateAuthor(ctx context.Context, req *api.AuthorRequest) (string, error)
	GetAllAuthors(ctx context.Context) ([]entity.Author, error)
	GetAuthorById(ctx context.Context, id string) (*entity.Author, error)
	GetAuthorByName(ctx context.Context, name string) ([]entity.Author, error)
	DeleteAuthor(ctx context.Context, id string) error
	UpdateAuthor(ctx context.Context, id string, req *api.AuthorRequest) error

	CreateFilePath(ctx context.Context, req *api.FilePathRequest) (string, error)
	GetAllFilePaths(ctx context.Context) ([]entity.FilePath, error)
	GetFilePathById(ctx context.Context, id string) (*entity.FilePath, error)
	DeleteFilePath(ctx context.Context, id string) error
	UpdateFilePath(ctx context.Context, id string, req *api.FilePathRequest) error
}
