package cache

import (
	"context"
	"encoding/json"
	"github.com/murat96k/kitaptar.kz/internal/auth/entity"
)

type UserCacher interface {
	GetUser(ctx context.Context, key string) (*entity.User, error)
	SetUser(ctx context.Context, value *entity.User) error
	DeleteUser(ctx context.Context, key string) error
}

func (c *Cache) GetUser(ctx context.Context, key string) (*entity.User, error) {

	value, err := c.redisCli.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if value == "" {
		return nil, nil
	}

	var user *entity.User

	err = json.Unmarshal([]byte(value), &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c *Cache) SetUser(ctx context.Context, value *entity.User) error {

	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.redisCli.Set(ctx, value.Id.String(), string(jsonValue), c.Expiration).Err()
}

func (c *Cache) DeleteUser(ctx context.Context, key string) error {
	return c.redisCli.Del(ctx, key).Err()
}
