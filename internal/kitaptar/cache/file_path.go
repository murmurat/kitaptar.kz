package cache

import (
	"context"
	"encoding/json"
	"github.com/murat96k/kitaptar.kz/internal/kitaptar/entity"
)

type FilePathCacher interface {
	GetFilePath(ctx context.Context, key string) (*entity.FilePath, error)
	SetFilePath(ctx context.Context, value *entity.FilePath) error
	DeleteFilePath(ctx context.Context, key string) error
}

func (c *Cache) GetFilePath(ctx context.Context, key string) (*entity.FilePath, error) {

	value, err := c.redisCli.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if value == "" {
		return nil, nil
	}

	var filePath *entity.FilePath

	err = json.Unmarshal([]byte(value), &filePath)
	if err != nil {
		return nil, err
	}

	return filePath, nil
}

func (c *Cache) SetFilePath(ctx context.Context, value *entity.FilePath) error {

	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.redisCli.Set(ctx, value.Id.String(), string(jsonValue), c.Expiration).Err()
}

func (c *Cache) DeleteFilePath(ctx context.Context, key string) error {
	return c.redisCli.Del(ctx, key).Err()
}
