package main

import (
	"net/http"

	"github.com/bhaskarvb28/facefind/internal/config"
	"github.com/bhaskarvb28/facefind/internal/handler"
	"github.com/bhaskarvb28/facefind/internal/server"
	"github.com/bhaskarvb28/facefind/internal/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	s, err := server.New(cfg)
	if err != nil {
		panic(err)
	}

	services, err := service.NewServices(s)
	if err != nil {
		panic(err)
	}

	handlers := handler.NewHandlers(s, services)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /auth/login", handlers.Auth.Login)
	mux.HandleFunc("GET /auth/callback", handlers.Auth.Callback)

	s.SetupHTTPServer(mux)

	if err := s.Start(); err != nil {
		panic(err)
	}

}
