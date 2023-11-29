package config

import "time"

type Config struct {
	HttpServer HttpServer `mapstructure:"HttpServer"`
	//	GrpcServer GrpcServer  `mapstructure:"GrpcServer"`
	Database    DBConfig    `mapstructure:"Database"`
	Auth        Auth        `mapstructure:"Auth"`
	Redis       RedisConfig `mapstructure:"Redis"`
	DebugServer DebugServer `mapstructure:"DebugServer"`
}

type Auth struct {
	PasswordSecretKey string        `mapstructure:"PasswordSecretKey"`
	JwtSecretKey      string        `mapstructure:"JwtSecretKey"`
	TimeToLive        time.Duration `mapstructure:"TimeToLive"`
}

type HttpServer struct {
	Port            string        `mapstructure:"port"`
	Timeout         time.Duration `mapstructure:"timeout"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DBName   string `mapstructure:"db_name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type RedisConfig struct {
	Address        string        `mapstructure:"address"`
	DB             int           `mapstructure:"db"`
	ExpirationTime time.Duration `mapstructure:"TimeToLive"`
}

type DebugServer struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}
