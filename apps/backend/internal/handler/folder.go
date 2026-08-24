package handler

import (
	"net/http"

	"github.com/bhaskarvb28/facefind/internal/model/folder"
	"github.com/bhaskarvb28/facefind/internal/server"
	"github.com/bhaskarvb28/facefind/internal/service"
	"github.com/labstack/echo/v4"
)

type FolderHandler struct {
	Handler
	folderService *service.FolderService
}

func NewFolderHandler(
	s *server.Server,
	folderService *service.FolderService,
) *FolderHandler {
	return &FolderHandler{
		Handler:       NewHandler(s),
		folderService: folderService,
	}
}

func (h *FolderHandler) Create(c echo.Context) error {
	var payload folder.CreatePayload

	// if err := c.Bind(&payload); err != nil {
	// 	return echo.NewHTTPError(
	// 		http.StatusBadRequest,
	// 		"invalid request body",
	// 	)
	// }

	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			err.Error(),
		)
	}

	if err := payload.Validate(); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			err.Error(),
		)
	}

	ownerSub, ok := c.Get("user_sub").(string)
	if !ok || ownerSub == "" {
		return echo.NewHTTPError(
			http.StatusUnauthorized,
			"unauthorized",
		)
	}

	f, err := h.folderService.Create(
		c.Request().Context(),
		ownerSub,
		payload.Name,
	)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to create folder",
		)
	}

	return c.JSON(http.StatusCreated, f)
}

func (h *FolderHandler) Get(c echo.Context) error {
	ownerSub, ok := c.Get("user_sub").(string)
	if !ok || ownerSub == "" {
		return echo.NewHTTPError(
			http.StatusUnauthorized,
			"Unauthorized",
		)
	}

	folders, err := h.folderService.Get(
		c.Request().Context(),
		ownerSub,
	)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to get folders",
		)
	}

	return c.JSON(http.StatusOK, folders)
}

func (h *FolderHandler) UpdateName(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"folder id is required",
		)
	}

	var req folder.UpdatePayload

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"invalid request body",
		)
	}

	if err := req.Validate(); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"folder name is required",
		)
	}

	ownerSub, ok := c.Get("user_sub").(string)

	if !ok || ownerSub == "" {
		return echo.NewHTTPError(
			http.StatusUnauthorized,
			"unauthorized",
		)
	}

	f, err := h.folderService.UpdateName(
		c.Request().Context(),
		ownerSub,
		id,
		req.Name,
	)

	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to update folder",
		)
	}

	return c.JSON(http.StatusOK, f)
}
