package v1

import (
	"github.com/bhaskarvb28/facefind/internal/handler"
	"github.com/bhaskarvb28/facefind/internal/service"
	"github.com/labstack/echo/v4"
)

func registerFaceRoutes(
	router *echo.Group,
	h *handler.Handlers,
	services *service.Services,
) {
	router.POST(
		"/photos/:photoId/faces",
		h.Face.Create,
	)
}

func registerFaceSearchRoutes(
	router *echo.Group,
	h *handler.Handlers,
	services *service.Services,
) {
	router.POST(
		"/folders/:folderId/faces/search",
		h.Face.Search,
	)
}