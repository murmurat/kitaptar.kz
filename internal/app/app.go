package app

import (
	"github.com/joho/godotenv"
	"github.com/murat96k/kitaptar.kz/internal/config"
	"github.com/murat96k/kitaptar.kz/internal/handler"
	"github.com/murat96k/kitaptar.kz/internal/repository/pgrepo"
	"github.com/murat96k/kitaptar.kz/internal/service"
	"github.com/murat96k/kitaptar.kz/pkg/httpserver"
	"github.com/murat96k/kitaptar.kz/pkg/jwttoken"
	"log"
	"os"
	"os/signal"
)

//func Init(config config.Config) {
//	server := httpserver.New()
//}

func Run(cfg *config.Config) error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := pgrepo.New(
		pgrepo.WithHost(cfg.DB.Host),
		pgrepo.WithUsername(cfg.DB.Username),
		pgrepo.WithPort(cfg.DB.Port),
		pgrepo.WithDBName(cfg.DB.DBName),
		pgrepo.WithPassword(cfg.DB.Password),
	)
	if err != nil {
		log.Printf("connection to DB error %s", err)
	}

	token := jwttoken.New(cfg.Token.SecretKey)

	service := service.New(db, cfg, token)
	handler := handler.New(service)

	server := httpserver.New(handler.InitRouter(),
		httpserver.WithReadTimeout(cfg.HTTP.ReadTimeout),
		httpserver.WithWriteTimeout(cfg.HTTP.WriteTimeout),
		httpserver.WithPort(cfg.HTTP.Port),
		httpserver.WithShutdownTimeout(cfg.HTTP.ShutdownTimeout),
	)
	log.Println("server started")
	server.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		log.Printf("signal received: %s", s.String())
	case err = <-server.Notify():
		log.Printf("server notify: %s", err.Error())
	}

	err = server.Shutdown()
	if err != nil {
		log.Printf("server shutdown err: %s", err)
	}

	return nil
}
