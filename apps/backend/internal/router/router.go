package router

import (
	"github.com/bhaskarvb28/facefind/internal/handler"
	"github.com/bhaskarvb28/facefind/internal/middleware"
	v1 "github.com/bhaskarvb28/facefind/internal/router/v1"
	"github.com/bhaskarvb28/facefind/internal/server"
	"github.com/bhaskarvb28/facefind/internal/service"
	"github.com/labstack/echo/v4"
)

func NewRouter(
	s *server.Server,
	h *handler.Handlers,
	services *service.Services,
) (*echo.Echo, error) {

	middlewares := middleware.NewMiddlewares(s)

	router := echo.New()

	// Global middleware
	router.Use(
		middlewares.Global.CORS(),
	)

	// API version
	v1Router := router.Group("/api/v1")

	// Public routes
	public := v1Router.Group("")
	v1.RegisterPublicRoutes(public, h, services)

	// Internal routes
	internal := v1Router.Group("/internal")
	internal.Use(middlewares.Internal.RequireWorkerAuth)
	v1.RegisterInternalRoutes(internal, h, services)

	// Authentication middleware
	authMiddleware, err := middleware.NewAuthMiddleware(s)
	if err != nil {
		return nil, err
	}

	// Protected routes
	protected := v1Router.Group("")
	protected.Use(authMiddleware.RequireAuth)
	v1.RegisterProtectedRoutes(protected, h, services)

	return router, nil
}