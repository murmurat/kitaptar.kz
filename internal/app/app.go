package app

import (
	"github.com/joho/godotenv"
	"github.com/murat96k/kitaptar.kz/internal/cache"
	"github.com/murat96k/kitaptar.kz/internal/config"
	"github.com/murat96k/kitaptar.kz/internal/handler"
	"github.com/murat96k/kitaptar.kz/internal/repository/pgrepo"
	"github.com/murat96k/kitaptar.kz/internal/service"
	pkg_redis "github.com/murat96k/kitaptar.kz/pkg/cache"
	"github.com/murat96k/kitaptar.kz/pkg/httpserver"
	"github.com/murat96k/kitaptar.kz/pkg/jwttoken"
	"log"
	"os"
	"os/signal"
)

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

	redisClient, err := pkg_redis.NewRedisClient(cfg)
	if err != nil {
		log.Fatal("[ERROR] connect to redis client error")
	}

	appCache := cache.NewCache(redisClient, cfg.Redis.ExpirationTime)

	cache, err := cache.NewAppCache(cache.WithUserCache(appCache), cache.WithAuthorCache(appCache), cache.WithBookCache(appCache), cache.WithFilePathCache(appCache))
	if err != nil {
		log.Fatalf("[ERROR] create cache error: %s", err.Error())
	}

	service := service.New(db, cfg, token, *cache)
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
