package service

import (
	"context"
	"one-lab/api"
	"one-lab/internal/entity"
)

func (m *Manager) GetUserBooks(email string) ([]entity.Book, error) {
	return m.Repository.GetUserBooks(email)
}

func (m *Manager) GetAllBooks(ctx context.Context) ([]entity.Book, error) {
	return m.Repository.GetAllBooks(ctx)
}

func (m *Manager) GetBookById(ctx context.Context, id string) (*entity.Book, error) {
	return m.Repository.GetBookById(ctx, id)
}

func (m *Manager) CreateBook(ctx context.Context, req *api.BookRequest) error {
	_, err := m.Repository.GetAuthorById(ctx, req.AuthorId.String())
	if err != nil {
		return err
	}

	return m.Repository.CreateBook(ctx, req)
}

func (m *Manager) DeleteBook(ctx context.Context, id string) error {
	return m.Repository.DeleteBook(ctx, id)
}

func (m *Manager) UpdateBook(ctx context.Context, id string, req *api.BookRequest) error {
	book, err := m.Repository.GetBookById(ctx, id)
	bookID := book.Id
	if err != nil {
		return err
	}
	return m.Repository.UpdateBook(ctx, bookID.String(), req)
}
