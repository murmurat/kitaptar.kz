package cache

import (
	"context"
	"encoding/json"
	"github.com/murat96k/kitaptar.kz/internal/entity"
	"github.com/redis/go-redis/v9"
	"time"
)

type Author interface {
	Get(ctx context.Context, key string) (*entity.Author, error)
	Set(ctx context.Context, key string, value *entity.Author) error
	Delete(ctx context.Context, key string) error
}

type AuthorCache struct {
	Expiration time.Duration
	redisCli   *redis.Client
}

func NewAuthorCache(redisCli *redis.Client, expiration time.Duration) Author {
	return &AuthorCache{
		redisCli:   redisCli,
		Expiration: expiration,
	}
}

func (a *AuthorCache) Get(ctx context.Context, key string) (*entity.Author, error) {

	value, err := a.redisCli.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if value == "" {
		return nil, nil
	}

	var author *entity.Author

	err = json.Unmarshal([]byte(value), &author)
	if err != nil {
		return nil, err
	}

	return author, nil
}

func (a *AuthorCache) Set(ctx context.Context, key string, value *entity.Author) error {

	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return a.redisCli.Set(ctx, key, string(jsonValue), a.Expiration).Err()
}

func (a *AuthorCache) Delete(ctx context.Context, key string) error {
	return a.redisCli.Del(ctx, key).Err()
}
