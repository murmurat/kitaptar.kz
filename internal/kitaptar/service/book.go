package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/murat96k/kitaptar.kz/api"
	"github.com/murat96k/kitaptar.kz/internal/kitaptar/entity"
	"github.com/redis/go-redis/v9"
)

func (m *Manager) GetUserBooks(email string) ([]entity.Book, error) {
	// TODO implement service
	return m.Repository.GetUserBooks(email)
}

func (m *Manager) GetAllBooks(ctx context.Context, sortBy string) ([]entity.Book, error) {
	return m.Repository.GetAllBooks(ctx, sortBy)
}

func (m *Manager) GetBookById(ctx context.Context, id string) (*entity.Book, error) {

	book, err := m.Cache.BookCache.GetBook(ctx, id)
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if book != nil {
		return book, nil
	}

	book, err = m.Repository.GetBookById(ctx, id)
	if err != nil {
		return nil, err
	}

	err = m.Cache.BookCache.SetBook(ctx, book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (m *Manager) GetBookByName(ctx context.Context, name string) ([]entity.Book, error) {
	return m.Repository.GetBookByName(ctx, name)
}

func (m *Manager) CreateBook(ctx context.Context, req *api.BookRequest) (string, error) {
	return m.Repository.CreateBook(ctx, req)
}

func (m *Manager) DeleteBook(ctx context.Context, id string) error {

	err := m.Cache.BookCache.DeleteBook(ctx, id)
	if err != nil {
		return err
	}

	return m.Repository.DeleteBook(ctx, id)
}

func (m *Manager) UpdateBook(ctx context.Context, id string, req *api.BookRequest) error {

	book, err := m.Repository.GetBookById(ctx, id)
	if err != nil {
		return err
	}

	if req.Name != "" {
		book.Name = req.Name
	}
	if req.Annotation != "" {
		book.Annotation = req.Annotation
	}
	if req.Genre != "" {
		book.Genre = req.Genre
	}
	if req.ImagePath != "" {
		book.ImagePath = req.ImagePath
	}
	if req.FilePathId != uuid.Nil {
		book.FilePathId = req.FilePathId
	}
	if req.AuthorId != uuid.Nil {
		book.AuthorId = req.AuthorId
	}

	err = m.Cache.BookCache.DeleteBook(ctx, book.Id.String())
	if err != nil {
		return err
	}

	err = m.Repository.UpdateBook(ctx, id, req)
	if err != nil {
		return err
	}

	_ = m.Cache.BookCache.SetBook(ctx, book)

	return nil
}

func (m *Manager) AddToFavorites(ctx context.Context, userId, bookId string) (string, error) {

	favoriteId, err := m.Repository.AddToFavorites(ctx, userId, bookId)
	if err != nil {
		return "", err
	}

	return favoriteId, nil
}

func (m *Manager) DeleteFromFavorites(ctx context.Context, userId, bookId string) error {

	favorite, err := m.GetFromFavorites(ctx, userId, bookId)
	if err != nil {
		return err
	}

	err = m.Repository.DeleteFromFavorites(ctx, favorite.Id.String())
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) GetFromFavorites(ctx context.Context, userId, bookId string) (*entity.FavoriteBook, error) {
	return m.Repository.GetFromFavorites(ctx, userId, bookId)
}
