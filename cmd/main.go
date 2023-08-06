package main

import (
	"log"
	"one-lab/internal/app"
	"one-lab/internal/config"
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
	cfg, err := config.InitConfig("config.yaml")

	if err != nil {
		log.Printf("config init err %w", err)
		//panic(err)
	}

	err = app.Run(cfg)
	if err != nil {
		log.Printf("config init err %w", err)
		//panic(err)
	}
}
