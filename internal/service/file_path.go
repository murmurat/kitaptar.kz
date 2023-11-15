package service

import (
	"context"
	"github.com/murat96k/kitaptar.kz/api"
	"github.com/murat96k/kitaptar.kz/internal/entity"
	"github.com/redis/go-redis/v9"
)

func (m *Manager) GetAllFilePaths(ctx context.Context) ([]entity.FilePath, error) {
	return m.Repository.GetAllFilePaths(ctx)
}

func (m *Manager) GetFilePathById(ctx context.Context, id string) (*entity.FilePath, error) {

	filePath, err := m.Cache.FilePathCache.Get(ctx, id)
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if filePath != nil {
		return filePath, nil
	}

	filePath, err = m.Repository.GetFilePathById(ctx, id)
	if err != nil {
		return nil, err
	}

	err = m.Cache.FilePathCache.Set(ctx, id, filePath)
	if err != nil {
		return nil, err
	}

	return filePath, nil
}

func (m *Manager) CreateFilePath(ctx context.Context, req *api.FilePathRequest) (string, error) {
	return m.Repository.CreateFilePath(ctx, req)
}

func (m *Manager) DeleteFilePath(ctx context.Context, id string) error {

	err := m.Cache.FilePathCache.Delete(ctx, id)
	if err != nil {
		return err
	}

	return m.Repository.DeleteFilePath(ctx, id)
}

func (m *Manager) UpdateFilePath(ctx context.Context, id string, req *api.FilePathRequest) error {

	filePath, err := m.Repository.GetFilePathById(ctx, id)
	if err != nil {
		return err
	}

	if req.Mobi != "" {
		filePath.Mobi = req.Mobi
	}
	if req.Fb2 != "" {
		filePath.Fb2 = req.Fb2
	}
	if req.Epub != "" {
		filePath.Epub = req.Epub
	}
	if req.Docx != "" {
		filePath.Docx = req.Docx
	}

	err = m.Cache.FilePathCache.Set(ctx, id, filePath)
	if err != nil {
		return err
	}

	return m.Repository.UpdateFilePath(ctx, id, req)
}
