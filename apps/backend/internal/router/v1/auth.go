package v1

import (
	"github.com/bhaskarvb28/facefind/internal/handler"
	"github.com/labstack/echo/v4"
)

func registerAuthRoutes(
	router *echo.Group,
	h *handler.Handlers,
) {
	auth := router.Group("/auth")

	// auth.POST("/login", h.Auth.Login)
	auth.POST("/signup", h.Auth.Signup)
	auth.POST("/login", h.Auth.Login)
	auth.GET("/callback", h.Auth.Callback)
}
