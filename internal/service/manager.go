package service

import (
	"github.com/murat96k/kitaptar.kz/internal/config"
	"github.com/murat96k/kitaptar.kz/internal/repository"
	"github.com/murat96k/kitaptar.kz/pkg/jwttoken"
)

type Manager struct {
	Repository repository.Repository
	Config     *config.Config
	Token      *jwttoken.JWTToken
}

func New(repository repository.Repository, config *config.Config, token *jwttoken.JWTToken) *Manager {
	return &Manager{Repository: repository, Config: config, Token: token}
}
