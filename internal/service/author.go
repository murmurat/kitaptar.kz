package service

import (
	"context"
	"one-lab/api"
	"one-lab/internal/entity"
)

func (m *Manager) GetAllAuthors(ctx context.Context) ([]entity.Author, error) {
	return m.Repository.GetAllAuthors(ctx)
}

func (m *Manager) GetAuthorById(ctx context.Context, id string) (*entity.Author, error) {
	return m.Repository.GetAuthorById(ctx, id)
}

func (m *Manager) CreateAuthor(ctx context.Context, req *api.AuthorRequest) error {
	return m.Repository.CreateAuthor(ctx, req)
}

func (m *Manager) DeleteAuthor(ctx context.Context, id string) error {
	// implement checking for existing book of that author
	return m.Repository.DeleteAuthor(ctx, id)
}

func (m *Manager) UpdateAuthor(ctx context.Context, id string, req *api.AuthorRequest) error {
	author, err := m.Repository.GetAuthorById(ctx, id)
	if err != nil {
		return err
	}
	authorID := author.Id
	return m.Repository.UpdateAuthor(ctx, authorID.String(), req)
}
