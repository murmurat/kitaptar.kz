package repository

import (
	"context"
	"one-lab/api"
	"one-lab/internal/entity"
)

type Repository interface {
	CreateUser(ctx context.Context, u *entity.User) error
	GetUser(ctx context.Context, email string) (*entity.User, error)
	UpdateUser(ctx context.Context, id string, req api.UpdateUserRequest) error

	GetUserBooks(email string) ([]entity.Book, error)
	GetAllBooks(ctx context.Context) ([]entity.Book, error)
	GetBookById(ctx context.Context, id string) (*entity.Book, error)
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
