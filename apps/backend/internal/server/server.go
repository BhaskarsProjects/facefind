package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bhaskarvb28/facefind/internal/config"
)

type Server struct {
	Config     *config.Config
	httpServer *http.Server
}

func New(cfg *config.Config) (*Server, error) {
	server := &Server{
		Config: cfg,
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
	if s.httpServer == nil {
		return nil
	}

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown HTTP server: %w", err)
	}

	return nil
}
