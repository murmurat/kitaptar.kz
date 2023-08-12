package service

import (
	"context"
	"one-lab/api"
	"one-lab/internal/entity"
)

type Service interface {
	CreateUser(ctx context.Context, u *entity.User) error
	Login(ctx context.Context, email, password string) (string, error)
	VerifyToken(token string) (string, error)
	UpdateUser(ctx context.Context, id string, req api.UpdateUserRequest) error //test
	//DeleteUser(ctx context.Context, id int64) error

	GetUserBooks(email string) ([]entity.Book, error)                 //test
	GetAllBooks(ctx context.Context) ([]entity.Book, error)           //test
	GetBookById(ctx context.Context, id string) (*entity.Book, error) //test

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
