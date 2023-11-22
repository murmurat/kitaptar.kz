package cache

import (
	"context"
	"encoding/json"
	"github.com/murat96k/kitaptar.kz/internal/kitaptar/entity"
)

type AuthorCacher interface {
	GetAuthor(ctx context.Context, key string) (*entity.Author, error)
	SetAuthor(ctx context.Context, value *entity.Author) error
	DeleteAuthor(ctx context.Context, key string) error
}

func (c *Cache) GetAuthor(ctx context.Context, key string) (*entity.Author, error) {

	value, err := c.redisCli.Get(ctx, key).Result()
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

func (c *Cache) SetAuthor(ctx context.Context, value *entity.Author) error {

	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.redisCli.Set(ctx, value.Id.String(), string(jsonValue), c.Expiration).Err()
}

func (c *Cache) DeleteAuthor(ctx context.Context, key string) error {
	return c.redisCli.Del(ctx, key).Err()
}
