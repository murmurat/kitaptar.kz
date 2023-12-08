package service

import (
	"github.com/murat96k/kitaptar.kz/internal/kafka"
	"github.com/murat96k/kitaptar.kz/internal/user/cache"
	"github.com/murat96k/kitaptar.kz/internal/user/config"
	"github.com/murat96k/kitaptar.kz/internal/user/repository"
	"github.com/murat96k/kitaptar.kz/pkg/jwttoken"
)

type Manager struct {
	Repository               repository.Repository
	Config                   *config.Config
	Token                    *jwttoken.JWTToken
	Cache                    cache.AppCache
	userVerificationProducer *kafka.Producer
}

func New(repository repository.Repository, config *config.Config, token *jwttoken.JWTToken, cache cache.AppCache, userVerificationProducer *kafka.Producer) *Manager {
	return &Manager{Repository: repository, Config: config, Token: token, Cache: cache, userVerificationProducer: userVerificationProducer}
}
