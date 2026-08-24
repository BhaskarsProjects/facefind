package handler

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/bhaskarvb28/facefind/internal/model/face"
	"github.com/bhaskarvb28/facefind/internal/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type FaceHandler struct {
	service *service.FaceService
}

func NewFaceHandler(
	service *service.FaceService,
) *FaceHandler {
	return &FaceHandler{
		service: service,
	}
}

func (h *FaceHandler) Create(c echo.Context) error {
	fmt.Println("reached here")

	photoID, err := uuid.Parse(c.Param("photoId"))
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"invalid photo id",
		)
	}

	var req face.CreateFacesRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"invalid request body",
		)
	}

	if err := h.service.Create(
		c.Request().Context(),
		photoID,
		req.Faces,
	); err != nil {
		fmt.Println(err)

		return echo.NewHTTPError(
			http.StatusBadRequest,
			err.Error(),
		)
	}

	return c.NoContent(http.StatusCreated)
}

func (h *FaceHandler) Search(c echo.Context) error {
	fmt.Println("reached search handler")
	folderID, err := uuid.Parse(c.Param("folderId"))
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"invalid folder id",
		)
	}

	file, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"image is required",
		)
	}

	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"failed to open image",
		)
	}
	defer src.Close()

	imageBytes, err := io.ReadAll(src)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"failed to read image",
		)
	}

	limit := 20

	if value := c.FormValue("limit"); value != "" {
		parsedLimit, err := strconv.Atoi(value)
		if err != nil || parsedLimit <= 0 {
			return echo.NewHTTPError(
				http.StatusBadRequest,
				"invalid limit",
			)
		}

		limit = parsedLimit
	}

	results, err := h.service.Search(
		c.Request().Context(),
		folderID,
		imageBytes,
		file.Filename,
		limit,
	)
	if err != nil {
		fmt.Println(err)

		return echo.NewHTTPError(
			http.StatusBadRequest,
			err.Error(),
		)
	}

	return c.JSON(
		http.StatusOK,
		map[string]any{
			"results": results,
		},
	)
}