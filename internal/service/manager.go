package service

import (
	"one-lab/internal/config"
	"one-lab/internal/repository"
	"one-lab/pkg/jwttoken"
)

type Manager struct {
	Repository repository.Repository
	Config     *config.Config
	Token      *jwttoken.JWTToken
}

func New(repository repository.Repository, config *config.Config, token *jwttoken.JWTToken) *Manager {
	return &Manager{Repository: repository, Config: config, Token: token}
}
