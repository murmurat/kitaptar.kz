package service

import (
	"context"
	"one-lab/api"
	"one-lab/internal/entity"
)

func (m *Manager) GetAllFilePaths(ctx context.Context) ([]entity.FilePath, error) {
	return m.Repository.GetAllFilePaths(ctx)
}

func (m *Manager) GetFilePathById(ctx context.Context, id string) (*entity.FilePath, error) {
	return m.Repository.GetFilePathById(ctx, id)
}

func (m *Manager) CreateFilePath(ctx context.Context, req *api.FilePathRequest) error {
	return m.Repository.CreateFilePath(ctx, req)
}

func (m *Manager) DeleteFilePath(ctx context.Context, id string) error {
	// implement checking for existing book of that file path
	return m.Repository.DeleteFilePath(ctx, id)
}

func (m *Manager) UpdateFilePath(ctx context.Context, id string, req *api.FilePathRequest) error {
	//filePath, err := m.Repository.GetFilePathById(ctx, id)
	//if err != nil {
	//	return err
	//}
	//filePathID := filePath.Id
	return m.Repository.UpdateFilePath(ctx, id, req)
}
