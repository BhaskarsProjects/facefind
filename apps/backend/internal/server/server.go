package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bhaskarvb28/facefind/internal/config"
	"github.com/bhaskarvb28/facefind/internal/database"
)

type Server struct {
	Config     *config.Config
	httpServer *http.Server
	DB         *database.Database
	// Redis      *redis.Client
}

func New(cfg *config.Config) (*Server, error) {
	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr: cfg.Redis.Address,
	// })

	// ctx := context.Background()

	// pong, err := redisClient.Ping(ctx).Result()
	// if err != nil {
	// 	return nil, fmt.Errorf("redis ping failed: %w", err)
	// }

	// fmt.Println("[REDIS] ping:", pong)

	// err = redisClient.Set(
	// 	ctx,
	// 	"debug:test",
	// 	"hello",
	// 	10*time.Minute,
	// ).Err()

	// if err != nil {
	// 	return nil, fmt.Errorf("redis debug set failed: %w", err)
	// }

	// fmt.Println("[REDIS] debug key created")

	db, err := database.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("database initialization failed: %w", err)
	}

	server := &Server{
		Config: cfg,
		DB:     db,
		// Redis:    redisClient,
	}

	return server, nil
}

func (s *Server) SetupHTTPServer(handler http.Handler) {
	s.httpServer = &http.Server{
		Addr:    ":" + s.Config.Server.Port,
		Handler: handler,
	}
}

func (s *Server) Start() error {
	if s.httpServer == nil {
		return fmt.Errorf("HTTP server is not initialized")
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	if s.httpServer != nil {
		if err := s.httpServer.Shutdown(ctx); err != nil {
			return fmt.Errorf(
				"failed to shutdown HTTP server: %w",
				err,
			)
		}
	}

	if s.DB != nil {
		s.DB.Close()
	}

	// if s.Redis != nil {
	// 	if err := s.Redis.Close(); err != nil {
	// 		return fmt.Errorf(
	// 			"failed to close Redis connection: %w",
	// 			err,
	// 		)
	// 	}
	// }

	return nil
}
