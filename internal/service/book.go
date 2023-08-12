package service

import (
	"context"
	"one-lab/internal/entity"
)

func (m *Manager) GetUserBooks(email string) ([]entity.Book, error) {
	return m.Repository.GetUserBooks(email)
}

func (m *Manager) GetAllBooks(ctx context.Context) ([]entity.Book, error) {
	return m.Repository.GetAllBooks(ctx)
}

func (m *Manager) GetBookById(ctx context.Context, id string) (entity.Book, error) {
	return m.Repository.GetBookById(ctx, id)
}
