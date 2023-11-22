package cache

type Option func(cache *AppCache)

func WithAuthorCache(author AuthorCacher) Option {
	return func(cache *AppCache) {
		cache.AuthorCache = author
	}
}

func WithBookCache(book BookCacher) Option {
	return func(cache *AppCache) {
		cache.BookCache = book
	}
}

func WithFilePathCache(filePath FilePathCacher) Option {
	return func(cache *AppCache) {
		cache.FilePathCache = filePath
	}
}
