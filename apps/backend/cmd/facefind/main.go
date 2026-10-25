package main

import (
	"fmt"

	"github.com/bhaskarvb28/facefind/internal/config"
	"github.com/bhaskarvb28/facefind/internal/handler"
	"github.com/bhaskarvb28/facefind/internal/repository"
	"github.com/bhaskarvb28/facefind/internal/router"
	"github.com/bhaskarvb28/facefind/internal/server"
	"github.com/bhaskarvb28/facefind/internal/service"
)

func main() {
	fmt.Println("Starting Application")
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	s, err := server.New(cfg)
	if err != nil {
		panic(err)
	}

	repos := repository.NewRepositories(s)
	services, err := service.NewServices(s, repos)
	if err != nil {
		panic(err)
	}

	handlers := handler.NewHandlers(s, services)

	httpRouter, err := router.NewRouter(s, handlers, services)
	if err != nil {
		fmt.Println(err)
	}

	s.SetupHTTPServer(httpRouter)

	if err := s.Start(); err != nil {
		panic(err)
	}
}
