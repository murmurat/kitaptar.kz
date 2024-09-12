package user

import (
	"context"
	"github.com/murat96k/kitaptar.kz/internal/kitaptar/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *config.Config) (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{Addr: cfg.Redis.Address, Password: "", DB: cfg.Redis.DB})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	return client, nil
}
