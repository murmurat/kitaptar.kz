package service

import (
	"context"

	"github.com/murat96k/kitaptar.kz/api"
	"github.com/murat96k/kitaptar.kz/internal/kitaptar/entity"
	"github.com/redis/go-redis/v9"
)

func (m *Manager) GetAllAuthors(ctx context.Context) ([]entity.Author, error) {
	return m.Repository.GetAllAuthors(ctx)
}

func (m *Manager) GetAuthorById(ctx context.Context, id string) (*entity.Author, error) {

	author, err := m.Cache.AuthorCache.GetAuthor(ctx, id)
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if author != nil {
		return author, nil
	}

	author, err = m.Repository.GetAuthorById(ctx, id)
	if err != nil {
		return nil, err
	}

	err = m.Cache.AuthorCache.SetAuthor(ctx, author)
	if err != nil {
		return nil, err
	}

	return author, nil
}

func (m *Manager) GetAuthorByName(ctx context.Context, name string) ([]entity.Author, error) {
	return m.Repository.GetAuthorByName(ctx, name)
}

func (m *Manager) CreateAuthor(ctx context.Context, req *api.AuthorRequest) (string, error) {
	return m.Repository.CreateAuthor(ctx, req)
}

func (m *Manager) DeleteAuthor(ctx context.Context, id string) error {

	err := m.Cache.AuthorCache.DeleteAuthor(ctx, id)
	if err != nil {
		return err
	}

	return m.Repository.DeleteAuthor(ctx, id)
}

func (m *Manager) UpdateAuthor(ctx context.Context, id string, req *api.AuthorRequest) error {

	author, err := m.Repository.GetAuthorById(ctx, id)
	if err != nil {
		return err
	}

	if req.Firstname != "" {
		author.Firstname = req.Firstname
	}
	if req.Lastname != "" {
		author.Lastname = req.Lastname
	}
	if req.AboutAuthor != "" {
		author.AboutAuthor = req.AboutAuthor
	}
	if req.ImagePath != "" {
		author.ImagePath = req.ImagePath
	}

	err = m.Cache.AuthorCache.DeleteAuthor(ctx, id)
	if err != nil {
		return err
	}

	err = m.Repository.UpdateAuthor(ctx, id, req)
	if err != nil {
		return err
	}

	_ = m.Cache.AuthorCache.SetAuthor(ctx, author)

	return nil
}
