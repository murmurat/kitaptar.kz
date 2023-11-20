package config

import "time"
import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	HTTP  ServerConfig `yaml:"http"`
	DB    DBConfig     `yaml:"db"`
	Token TokenConfig  `yaml:"token"`
	Redis RedisConfig  `yaml:"redis"`
}

type TokenConfig struct {
	SecretKey  string        `env:"TOKEN_SECRET_KEY" env-default:"one-lab"`
	TimeToLive time.Duration `yaml:"time_to_live"`
}
type ServerConfig struct {
	Port            string        `yaml:"port"`
	Timeout         time.Duration `yaml:"timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
}
type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"db_name"`
	Username string `yaml:"username"`
	//MigrationPath    string `yaml:"migration_path"`
	//MigrationVersion uint   `yaml:"migration_version"`
	Password string `env:"DB_PASSWORD" env-default:"postgres"`
}

type RedisConfig struct {
	Address        string        `yaml:"address"`
	ExpirationTime time.Duration `yaml:"expiration_time"`
}

func InitConfig(path string) (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)

	return cfg, nil
}
