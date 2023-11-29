package config

import "time"

type Config struct {
	HttpServer HttpServer  `mapstructure:"HttpServer"`
	Database   Database    `mapstructure:"Database"`
	Auth       Auth        `mapstructure:"Auth"`
	Transport  Transport   `mapstructure:"Transport"`
	Redis      RedisConfig `mapstructure:"Redis"`
	//	Kafka      Kafka      `mapstructure:"Kafka"`
}

type HttpServer struct {
	Port            string        `mapstructure:"port"`
	Timeout         time.Duration `mapstructure:"timeout"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DBName   string `mapstructure:"db_name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Auth struct {
	Access  Access  `mapstructure:"Access"`
	Refresh Refresh `mapstructure:"Refresh"`
}

type Access struct {
	PasswordSecretKey string        `mapstructure:"PasswordSecretKey"`
	JwtSecretKey      string        `mapstructure:"JwtSecretKey"`
	TimeToLive        time.Duration `mapstructure:"TimeToLive"`
}

type Refresh struct {
	PasswordSecretKey string        `mapstructure:"PasswordSecretKey"`
	JwtSecretKey      string        `mapstructure:"JwtSecretKey"`
	TimeToLive        time.Duration `mapstructure:"TimeToLive"`
}

type Transport struct {
	User     UserTransport     `mapstructure:"user"`
	UserGrpc UserGrpcTransport `mapstructure:"userGrpc"`
}

type UserTransport struct {
	Host    string        `mapstructure:"host"`
	Timeout time.Duration `mapstructure:"timeout"`
}

type UserGrpcTransport struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type RedisConfig struct {
	Address        string        `mapstructure:"address"`
	DB             int           `mapstructure:"db"`
	ExpirationTime time.Duration `mapstructure:"TimeToLive"`
}

// TODO Implement Kafka and uncomment
//type Kafka struct {
//	Brokers  []string `mapstructure:"brokers"`
//	Producer Producer `mapstructure:"producer"`
//	Consumer Consumer `mapstructure:"consumer"`
//}
//
//type Producer struct {
//	Topic string `mapstructure:"topic"`
//}
//
//type Consumer struct {
//	Topics []string `mapstructure:"topics"`
//}
