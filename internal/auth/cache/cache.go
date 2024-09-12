package cache

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type AppCache struct {
	TokenCache TokenCacher
}

func NewAppCache(opts ...Option) (*AppCache, error) {
	cache := new(AppCache)

	for _, opt := range opts {
		opt(cache)
	}

	return cache, nil
}

type Cache struct {
	Expiration time.Duration
	redisCli   *redis.Client
}

func NewCache(redisCli *redis.Client, expiration time.Duration) *Cache {
	return &Cache{
		redisCli:   redisCli,
		Expiration: expiration,
	}
}
