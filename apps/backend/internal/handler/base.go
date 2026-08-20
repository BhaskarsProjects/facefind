package handler

import "github.com/bhaskarvb28/facefind/internal/server"

// Handler provides base functionality for all handlers
type Handler struct {
	server *server.Server
}

// NewHandler creates a new base handler
func NewHandler(s *server.Server) Handler {
	return Handler{server: s}
}
