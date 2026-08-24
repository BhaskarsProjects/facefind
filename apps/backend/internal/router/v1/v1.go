package v1

import (
	"github.com/bhaskarvb28/facefind/internal/handler"
	"github.com/bhaskarvb28/facefind/internal/service"
	"github.com/labstack/echo/v4"
)

func RegisterPublicRoutes(
	router *echo.Group,
	h *handler.Handlers,
	services *service.Services,
) {
	registerAuthRoutes(router, h)
}

func RegisterProtectedRoutes(
	router *echo.Group,
	h *handler.Handlers,
	services *service.Services,
) {
	registerFolderRoutes(router, h, services)
	registerFaceSearchRoutes(router, h, services)
}

func RegisterInternalRoutes(
	router *echo.Group,
	h *handler.Handlers,
	services *service.Services,
) {
	registerFaceRoutes(router, h, services)
}
