package cache

import (
	"context"
	"encoding/json"
	"github.com/murat96k/kitaptar.kz/internal/entity"
)

type BookCacher interface {
	GetBook(ctx context.Context, key string) (*entity.Book, error)
	SetBook(ctx context.Context, value *entity.Book) error
	DeleteBook(ctx context.Context, key string) error
}

func (c *Cache) GetBook(ctx context.Context, key string) (*entity.Book, error) {

	value, err := c.redisCli.Get(ctx, key).Result()
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

func (c *Cache) SetBook(ctx context.Context, value *entity.Book) error {

	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.redisCli.Set(ctx, value.Id.String(), string(jsonValue), c.Expiration).Err()
}

func (c *Cache) DeleteBook(ctx context.Context, key string) error {
	return c.redisCli.Del(ctx, key).Err()
}
