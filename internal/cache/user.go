package cache

import (
	"context"
	"encoding/json"
	"github.com/murat96k/kitaptar.kz/internal/entity"
	"github.com/redis/go-redis/v9"
	"time"
)

type User interface {
	Get(ctx context.Context, key string) (*entity.User, error)
	Set(ctx context.Context, key string, value *entity.User) error
	Delete(ctx context.Context, key string) error
}

type UserCache struct {
	Expiration time.Duration
	redisCli   *redis.Client
}

func NewUserCache(redisCli *redis.Client, expiration time.Duration) User {
	return &UserCache{
		redisCli:   redisCli,
		Expiration: expiration,
	}
}

func (u *UserCache) Get(ctx context.Context, key string) (*entity.User, error) {

	value, err := u.redisCli.Get(ctx, key).Result()
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

func (u *UserCache) Set(ctx context.Context, key string, value *entity.User) error {

	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return u.redisCli.Set(ctx, key, string(jsonValue), u.Expiration).Err()
}

func (u *UserCache) Delete(ctx context.Context, key string) error {
	return u.redisCli.Del(ctx, key).Err()
}
