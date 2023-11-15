package cache

type Cache struct {
	UserCache     User
	AuthorCache   Author
	BookCache     Book
	FilePathCache FilePath
}

func NewCache(opts ...Option) (*Cache, error) {
	cache := new(Cache)

	for _, opt := range opts {
		opt(cache)
	}

	return cache, nil
}
