package main

import (
	"fmt"

	"github.com/bhaskarvb28/go_backend_boilerplate/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	fmt.Println(cfg.Observability.GetLogLevel())
	fmt.Println(cfg.Observability.IsProduction())
}