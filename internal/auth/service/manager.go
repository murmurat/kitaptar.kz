package service

import (
	"github.com/murat96k/kitaptar.kz/internal/auth/cache"
	"github.com/murat96k/kitaptar.kz/internal/auth/config"
	"github.com/murat96k/kitaptar.kz/internal/auth/repository"
	"github.com/murat96k/kitaptar.kz/pkg/jwttoken"
	userClient "github.com/uristemov/auth-user-grpc/client/grpc"
)

type Manager struct {
	Repository repository.Repository
	Config     *config.Config
	Token      *jwttoken.JWTToken
	UserClient *userClient.Client
	Cache      cache.AppCache
}

func New(repository repository.Repository, config *config.Config, token *jwttoken.JWTToken, userClient *userClient.Client, cache cache.AppCache) *Manager {
	return &Manager{Repository: repository, Config: config, Token: token, UserClient: userClient, Cache: cache}
}
