package middleware

import (
	"github.com/bhaskarvb28/facefind/internal/server"
)

type Middlewares struct {
	Global *GlobalMiddlewares
	Internal *InternalMiddleware
	// Auth            *AuthMiddleware
	// ContextEnhancer *ContextEnhancer
	// Tracing         *TracingMiddleware
	// RateLimit       *RateLimitMiddleware
}

func NewMiddlewares(s *server.Server) *Middlewares {

	return &Middlewares{
		Global: NewGlobalMiddlewares(s),
		Internal: NewInternalMiddleware(s),
		// Auth:            NewAuthMiddleware(s),
		// ContextEnhancer: NewContextEnhancer(s),
		// Tracing:         NewTracingMiddleware(s, nrApp),
		// RateLimit:       NewRateLimitMiddleware(s),
	}
}
