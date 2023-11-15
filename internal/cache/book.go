package cache

import (
	"context"
	"encoding/json"
	"github.com/murat96k/kitaptar.kz/internal/entity"
	"github.com/redis/go-redis/v9"
	"time"
)

type Book interface {
	Get(ctx context.Context, key string) (*entity.Book, error)
	Set(ctx context.Context, key string, value *entity.Book) error
	Delete(ctx context.Context, key string) error
}

type BookCache struct {
	Expiration time.Duration
	redisCli   *redis.Client
}

func NewBookCache(redisCli *redis.Client, expiration time.Duration) Book {
	return &BookCache{
		redisCli:   redisCli,
		Expiration: expiration,
	}
}

func (b *BookCache) Get(ctx context.Context, key string) (*entity.Book, error) {

	value, err := b.redisCli.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if value == "" {
		return nil, nil
	}

	var book *entity.Book

	err = json.Unmarshal([]byte(value), &book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (b *BookCache) Set(ctx context.Context, key string, value *entity.Book) error {

	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return b.redisCli.Set(ctx, key, string(jsonValue), b.Expiration).Err()
}

func (b *BookCache) Delete(ctx context.Context, key string) error {
	return b.redisCli.Del(ctx, key).Err()
}
