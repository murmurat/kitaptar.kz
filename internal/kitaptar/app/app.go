package app

import (
	"github.com/murat96k/kitaptar.kz/internal/kitaptar/cache"
	"github.com/murat96k/kitaptar.kz/internal/kitaptar/config"
	"github.com/murat96k/kitaptar.kz/internal/kitaptar/handler"
	"github.com/murat96k/kitaptar.kz/internal/kitaptar/repository"
	"github.com/murat96k/kitaptar.kz/internal/kitaptar/service"
	pkg_redis "github.com/murat96k/kitaptar.kz/pkg/cache/kitaptar"
	"github.com/murat96k/kitaptar.kz/pkg/httpserver"
	"github.com/murat96k/kitaptar.kz/pkg/jwttoken"
	"log"
	"os"
	"os/signal"
)

func Run(cfg *config.Config) error {

	db, err := repository.New(
		repository.WithHost(cfg.Database.Host),
		repository.WithUsername(cfg.Database.Username),
		repository.WithPort(cfg.Database.Port),
		repository.WithDBName(cfg.Database.DBName),
		repository.WithPassword(cfg.Database.Password),
	)
	if err != nil {
		log.Printf("connection to DB error %s", err)
	}

	token := jwttoken.New(cfg.Auth.JwtSecretKey)

	redisClient, err := pkg_redis.NewRedisClient(cfg)
	if err != nil {
		log.Fatal("[ERROR] connect to redis client error")
	}

	appCache := cache.NewCache(redisClient, cfg.Redis.ExpirationTime)

	cache, err := cache.NewAppCache(cache.WithAuthorCache(appCache), cache.WithBookCache(appCache), cache.WithFilePathCache(appCache))
	if err != nil {
		log.Fatalf("[ERROR] create cache error: %s", err.Error())
	}

	service := service.New(db, cfg, token, *cache)
	handler := handler.New(service)

	server := httpserver.New(handler.InitRouter(),
		httpserver.WithReadTimeout(cfg.HttpServer.ReadTimeout),
		httpserver.WithWriteTimeout(cfg.HttpServer.WriteTimeout),
		httpserver.WithPort(cfg.HttpServer.Port),
		httpserver.WithShutdownTimeout(cfg.HttpServer.ShutdownTimeout),
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
