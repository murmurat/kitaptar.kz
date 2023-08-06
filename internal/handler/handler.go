package handler

import "one-lab/internal/service"

type Handler struct {
	srvs service.Service
}

func New(srvc service.Service) *Handler {
	return &Handler{
		srvs: srvc,
	}
}
