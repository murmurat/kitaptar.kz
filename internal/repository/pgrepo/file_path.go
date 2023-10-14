package pgrepo

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/murat96k/kitaptar.kz/api"
	"github.com/murat96k/kitaptar.kz/internal/entity"
)

func (p *Postgres) GetAllFilePaths(ctx context.Context) ([]entity.FilePath, error) {
	var filePaths []entity.FilePath
	query := fmt.Sprintf("SELECT id, mobi, fb2, epub, docx FROM %s", filePathsTable)
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
			log.Printf("Scan filePath values error %s", err.Error())
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

	query := fmt.Sprintf("SELECT id, mobi,fb2, epub, docx FROM %s WHERE id='%s'", filePathsTable, strings.TrimSpace(id))
	err := pgxscan.Get(ctx, p.Pool, filePath, query)
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
			INSERT INTO %s (mobi, fb2, epub ,docx, created_at)
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
	query := fmt.Sprintf("DELETE FROM %s WHERE id='%s'", filePathsTable, id)

	_, err := p.Pool.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) UpdateFilePath(ctx context.Context, id string, req *api.FilePathRequest) error {
	values := make([]string, 0)

	if req.Mobi != "" {
		values = append(values, fmt.Sprintf("mobi='%s'", req.Mobi))
	}
	if req.Fb2 != "" {
		values = append(values, fmt.Sprintf("fb2='%s'", req.Fb2))
	}
	if req.Epub != "" {
		values = append(values, fmt.Sprintf("epub='%s'", req.Epub))
	}
	if req.Docx != "" {
		values = append(values, fmt.Sprintf("docx='%s'", req.Docx))
	}

	setQuery := strings.Join(values, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = '%s';", filePathsTable, setQuery, id)
	fmt.Println(query)

	_, err := p.Pool.Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}
