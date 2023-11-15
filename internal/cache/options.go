package cache

type Option func(cache *Cache)

func WithUserCache(user User) Option {
	return func(cache *Cache) {
		cache.UserCache = user
	}
}

func WithAuthorCache(author Author) Option {
	return func(cache *Cache) {
		cache.AuthorCache = author
	}
}

func WithBookCache(book Book) Option {
	return func(cache *Cache) {
		cache.BookCache = book
	}
}

func WithFilePathCache(filePath FilePath) Option {
	return func(cache *Cache) {
		cache.FilePathCache = filePath
	}
}
