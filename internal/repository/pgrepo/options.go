package pgrepo

type Option func(*Postgres)

func WithHost(host string) Option {
	return func(postgres *Postgres) {
		postgres.host = host
	}
}

func WithPort(port string) Option {
	return func(postgres *Postgres) {
		postgres.port = port
	}
}

func WithUsername(username string) Option {
	return func(postgres *Postgres) {
		postgres.username = username
	}
}

func WithPassword(password string) Option {
	return func(postgres *Postgres) {
		postgres.password = password
	}
}

func WithDBName(dbName string) Option {
	return func(postgres *Postgres) {
		postgres.dbName = dbName
	}
}
