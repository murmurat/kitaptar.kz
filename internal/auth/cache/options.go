package cache

type Option func(cache *AppCache)

func WithUserCache(user TokenCacher) Option {
	return func(cache *AppCache) {
		cache.TokenCache = user
	}
}
