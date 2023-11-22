package user

import (
	"context"
	"github.com/murat96k/kitaptar.kz/internal/user/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *config.Config) (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{Addr: cfg.Redis.Address, Password: "", DB: 0})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	return client, nil
}
