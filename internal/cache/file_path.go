package cache

import (
	"context"
	"encoding/json"
	"github.com/murat96k/kitaptar.kz/internal/entity"
	"github.com/redis/go-redis/v9"
	"time"
)

type FilePath interface {
	Get(ctx context.Context, key string) (*entity.FilePath, error)
	Set(ctx context.Context, key string, value *entity.FilePath) error
	Delete(ctx context.Context, key string) error
}

type FilePathCache struct {
	Expiration time.Duration
	redisCli   *redis.Client
}

func NewFilePathCache(redisCli *redis.Client, expiration time.Duration) FilePath {
	return &FilePathCache{
		redisCli:   redisCli,
		Expiration: expiration,
	}
}

func (f *FilePathCache) Get(ctx context.Context, key string) (*entity.FilePath, error) {

	value, err := f.redisCli.Get(ctx, key).Result()
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

func (f *FilePathCache) Set(ctx context.Context, key string, value *entity.FilePath) error {

	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return f.redisCli.Set(ctx, key, string(jsonValue), f.Expiration).Err()
}

func (f *FilePathCache) Delete(ctx context.Context, key string) error {
	return f.redisCli.Del(ctx, key).Err()
}
