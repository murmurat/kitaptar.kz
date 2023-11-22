package repository

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/murat96k/kitaptar.kz/api"
	"github.com/murat96k/kitaptar.kz/internal/kitaptar/entity"
	"strings"
	"time"
)

func (p *Postgres) GetAllFilePaths(ctx context.Context) ([]entity.FilePath, error) {

	var filePaths []entity.FilePath
	query := fmt.Sprintf("SELECT id, mobi, fb2, epub ,docx FROM %s", filePathsTable)

	rows, err := p.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		filePath := entity.FilePath{}
		err = rows.Scan(&filePath.Id, &filePath.Mobi, &filePath.Fb2, &filePath.Epub, &filePath.Docx)
		filePaths = append(filePaths, filePath)
		if err != nil {
			return nil, err
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return filePaths, nil
}

func (p *Postgres) GetFilePathById(ctx context.Context, id string) (*entity.FilePath, error) {

	filePath := new(entity.FilePath)

	query := fmt.Sprintf("SELECT id, mobi, fb2, epub, docx FROM %s WHERE id=$1", filePathsTable)
	err := pgxscan.Get(ctx, p.Pool, filePath, query, strings.TrimSpace(id))
	if err != nil {
		return nil, err
	}

	return filePath, nil
}

func (p *Postgres) CreateFilePath(ctx context.Context, req *api.FilePathRequest) (string, error) {

	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return "", err
	}

	var filePathId string

	query := fmt.Sprintf(`
			INSERT INTO %s (mobi, fb2, epub, docx, created_at)
			VALUES ($1, $2, $3, $4, $5) RETURNING id
			`, filePathsTable)

	err = p.Pool.QueryRow(ctx, query, req.Mobi, req.Fb2, req.Epub, req.Docx, time.Now()).Scan(&filePathId)
	if err != nil {
		tx.Rollback(ctx)
		return "", err
	}

	return filePathId, tx.Commit(ctx)
}

func (p *Postgres) DeleteFilePath(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", filePathsTable)

	_, err := p.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) UpdateFilePath(ctx context.Context, id string, req *api.FilePathRequest) error {

	values := make([]string, 0)
	paramCount := 2
	params := make([]interface{}, 0)

	if req.Mobi != "" {
		values = append(values, fmt.Sprintf("mobi=$%d", paramCount))
		params = append(params, req.Mobi)
		paramCount++
	}
	if req.Fb2 != "" {
		values = append(values, fmt.Sprintf("fb2=$%d", paramCount))
		params = append(params, req.Fb2)
		paramCount++
	}
	if req.Epub != "" {
		values = append(values, fmt.Sprintf("epub=$%d", paramCount))
		params = append(params, req.Epub)
		paramCount++
	}
	if req.Docx != "" {
		values = append(values, fmt.Sprintf("docx=$%d", paramCount))
		params = append(params, req.Docx)
		paramCount++
	}

	setQuery := strings.Join(values, ", ")
	setQuery = fmt.Sprintf("UPDATE %s SET ", filePathsTable) + setQuery + " WHERE id=$1"

	params = append([]interface{}{id}, params...)

	_, err := p.Pool.Exec(ctx, setQuery, params...)
	if err != nil {
		return err
	}

	return nil
}
