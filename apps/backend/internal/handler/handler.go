package handler

import (
	"github.com/bhaskarvb28/facefind/internal/server"
	"github.com/bhaskarvb28/facefind/internal/service"
)

type Handlers struct {
	Auth *AuthHandler
}

func NewHandlers(s *server.Server, services *service.Services) *Handlers {
	return &Handlers{
		Auth: NewAuthHandler(s, services.Auth),
	}
}
