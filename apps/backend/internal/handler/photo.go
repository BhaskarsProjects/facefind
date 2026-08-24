package handler

import (
	"net/http"

	"github.com/bhaskarvb28/facefind/internal/server"
	"github.com/bhaskarvb28/facefind/internal/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PhotoHandler struct {
	server  *server.Server
	service *service.PhotoService
}

func NewPhotoHandler(
	s *server.Server,
	service *service.PhotoService,
) *PhotoHandler {
	return &PhotoHandler{
		server:  s,
		service: service,
	}
}

func (h *PhotoHandler) Upload(c echo.Context) error {
	folderID := c.Param("folderId")

	if folderID == "" {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"folder id is required",
		)
	}

	// We'll get this from the authentication middleware.
	sub, ok := c.Get("user_sub").(string)
	if !ok || sub == "" {
		return echo.NewHTTPError(
			http.StatusUnauthorized,
			"unauthorized",
		)
	}

	form, err := c.MultipartForm()
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"invalid multipart form",
		)
	}

	files := form.File["photos"]

	if len(files) == 0 {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"no photos provided",
		)
	}

	if len(files) > 100 {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"maximum 100 photos can be uploaded at once",
		)
	}

	// Validation and processing will be moved into the service.
	for _, file := range files {
		if file.Size > 25*1024*1024 {
			return echo.NewHTTPError(
				http.StatusBadRequest,
				"each photo must be 25 MB or smaller",
			)
		}
	}

	err = h.service.Upload(
		c.Request().Context(),
		sub,
		folderID,
		files,
	)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "photos uploaded successfully",
	})
}

func (h *PhotoHandler) GetFolderPhotos(c echo.Context) error {
	folderID, err := uuid.Parse(c.Param("folderId"))
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"invalid folder id",
		)
	}

	userSub, ok := c.Get("user_sub").(string)
	if !ok || userSub == "" {
		return echo.NewHTTPError(
			http.StatusUnauthorized,
			"unauthorized",
		)
	}

	photos, err := h.service.GetFolderPhotos(
		c.Request().Context(),
		folderID,
		userSub,
	)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to get folder photos",
		)
	}

	return c.JSON(http.StatusOK, map[string]any{
		"photos": photos,
	})
}

func (h *PhotoHandler) DeletePhoto(c echo.Context) error {
	folderID, err := uuid.Parse(c.Param("folderId"))
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"invalid folder id",
		)
	}

	photoID, err := uuid.Parse(c.Param("photoId"))
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"invalid photo id",
		)
	}

	userSub, ok := c.Get("user_sub").(string)
	if !ok || userSub == "" {
		return echo.NewHTTPError(
			http.StatusUnauthorized,
			"unauthorized",
		)
	}

	err = h.service.Delete(
		c.Request().Context(),
		folderID,
		photoID,
		userSub,
	)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to delete photo",
		)
	}

	return c.NoContent(http.StatusNoContent)
}