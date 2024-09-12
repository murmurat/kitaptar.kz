package main

import (
	"fmt"
	"log"

	"github.com/murat96k/kitaptar.kz/internal/kitaptar/app"
	"github.com/murat96k/kitaptar.kz/internal/kitaptar/config"
	"github.com/spf13/viper"
)

// @title           ONE LAB Kitaptar
// @version         0.0.1
// @description     API for Book application

// @contact.name   Meiirzhan
// @contact.email  admin@kitaptar.kz

// @host      localhost:8080
// @BasePath  /main

// @securitydefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	cfg, err := loadConfig("config/kitaptar")

	if err != nil {
		log.Fatalf("config init err %s", err)
	}

	err = app.Run(cfg)
	if err != nil {
		log.Fatalf("app run err %s", err)
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
