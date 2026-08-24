package v1

import (
	"github.com/bhaskarvb28/facefind/internal/handler"
	"github.com/bhaskarvb28/facefind/internal/service"
	"github.com/labstack/echo/v4"
)

func registerFolderRoutes(
	router *echo.Group,
	h *handler.Handlers,
	services *service.Services,
) {
	router.POST("/folders", h.Folder.Create)
	router.GET("/folders", h.Folder.Get)
	router.PATCH("/folders/:id", h.Folder.UpdateName)

	// Photos
	router.POST("/folders/:folderId/photos", h.Photo.Upload)
	router.GET("/:folderId/photos", h.Photo.GetFolderPhotos)
	router.DELETE("/:folderId/photos/:photoId", h.Photo.DeletePhoto)
}
