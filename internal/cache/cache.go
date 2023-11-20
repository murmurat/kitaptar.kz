package cache

import (
	"github.com/redis/go-redis/v9"
	"time"
)

type AppCache struct {
	UserCache     UserCacher
	AuthorCache   AuthorCacher
	BookCache     BookCacher
	FilePathCache FilePathCacher
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
