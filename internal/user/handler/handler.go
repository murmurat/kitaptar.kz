package handler

import "github.com/murat96k/kitaptar.kz/internal/user/service"

type Handler struct {
	srvs service.Service
}

func New(srvc service.Service) *Handler {
	return &Handler{
		srvs: srvc,
	}
}
