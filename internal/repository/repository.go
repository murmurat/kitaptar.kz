package repository

import (
	"context"
	"one-lab/api"
	"one-lab/internal/entity"
)

type Repository interface {
	CreateUser(ctx context.Context, u *entity.User) error
	GetUser(ctx context.Context, email string) (*entity.User, error)
	UpdateUser(ctx context.Context, id string, req *api.UpdateUserRequest) error
	DeleteUser(ctx context.Context, id string) error

	CreateBook(ctx context.Context, req *api.BookRequest) error
	GetUserBooks(email string) ([]entity.Book, error)
	GetAllBooks(ctx context.Context) ([]entity.Book, error)
	GetBookById(ctx context.Context, id string) (*entity.Book, error)
	DeleteBook(ctx context.Context, id string) error                       //test
	UpdateBook(ctx context.Context, id string, req *api.BookRequest) error //test

	CreateAuthor(ctx context.Context, req *api.AuthorRequest) error            //test
	GetAllAuthors(ctx context.Context) ([]entity.Author, error)                //test
	GetAuthorById(ctx context.Context, id string) (*entity.Author, error)      //test
	DeleteAuthor(ctx context.Context, id string) error                         //test
	UpdateAuthor(ctx context.Context, id string, req *api.AuthorRequest) error //test

	CreateFilePath(ctx context.Context, req *api.FilePathRequest) error            //test
	GetAllFilePaths(ctx context.Context) ([]entity.FilePath, error)                //test
	GetFilePathById(ctx context.Context, id string) (*entity.FilePath, error)      //test
	DeleteFilePath(ctx context.Context, id string) error                           //test
	UpdateFilePath(ctx context.Context, id string, req *api.FilePathRequest) error //test
	//UpdateUser(ctx context.Context, u *entity.User) error
	//DeleteUser(ctx context.Context, id int64) error
	//VerifyToken(token string) error
	//
	//CreateArticle(ctx context.Context, a *entity.Article) error
	//UpdateArticle(ctx context.Context, a *entity.Article) error
	//DeleteArticle(ctx context.Context, id int64) error
	//GetArticleByID(ctx context.Context, id int64) (*entity.Article, error)
	//GetAllArticles(ctx context.Context) ([]entity.Article, error)
	//GetArticlesByUserID(ctx context.Context, userID int64) ([]entity.Article, error)
	//
	//GetCategories(ctx context.Context) ([]entity.Category, error)
}
