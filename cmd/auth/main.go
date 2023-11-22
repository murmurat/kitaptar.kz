package main

import (
	"fmt"
	"github.com/murat96k/kitaptar.kz/internal/auth/app"
	"github.com/murat96k/kitaptar.kz/internal/auth/config"
	"github.com/spf13/viper"
	"log"
)

func main() {
	cfg, err := loadConfig("config/auth")

	if err != nil {
		log.Printf("config init err %s", err)
	}

	err = app.Run(cfg)
	if err != nil {
		log.Printf("config init err %s", err)
	}
}

func loadConfig(path string) (config *config.Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("failed to ReadInConfig err: %w", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("failed to Unmarshal config err: %w", err)
	}

	return config, nil
}
