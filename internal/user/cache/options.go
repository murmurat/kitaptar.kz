package cache

type Option func(cache *AppCache)

func WithUserCache(user UserCacher) Option {
	return func(cache *AppCache) {
		cache.UserCache = user
	}
}

func WithCodeCache(code CodeCacher) Option {
	return func(cache *AppCache) {
		cache.CodeCache = code
	}
}
