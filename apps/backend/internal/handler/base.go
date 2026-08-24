package handler

import (
	"fmt"
	"time"

	"github.com/bhaskarvb28/facefind/internal/server"
	"github.com/bhaskarvb28/facefind/internal/validation"
	"github.com/labstack/echo/v4"
)

// Handler provides base functionality for all handlers
type Handler struct {
	server *server.Server
}

// HandlerFunc represents a typed handler function
// that processes a request and returns a response.
type HandlerFunc[Req validation.Validatable, Res any] func(
	c echo.Context,
	req Req,
) (Res, error)

type ResponseHandler interface {
	Handle(c echo.Context, result interface{}) error
}

type JSONResponseHandler struct {
	status int
}

func (h JSONResponseHandler) Handle(
	c echo.Context,
	result interface{},
) error {
	return c.JSON(h.status, result)
}

// NewHandler creates a new base handler
func NewHandler(s *server.Server) Handler {
	return Handler{server: s}
}

func handleRequest[Req validation.Validatable, Res any](
	c echo.Context,
	req Req,
	handler HandlerFunc[Req, Res],
	responseHandler ResponseHandler,
) error {
	start := time.Now()

	method := c.Request().Method
	path := c.Request().URL.Path
	route := c.Path()

	fmt.Printf(
		"[HTTP] %s %s route=%s\n",
		method,
		path,
		route,
	)

	validationStart := time.Now()

	if err := validation.BindAndValidate(c, req); err != nil {
		validationDuration := time.Since(validationStart)

		fmt.Printf(
			"[VALIDATION] failed duration=%s error=%v\n",
			validationDuration,
			err,
		)

		return err
	}

	validationDuration := time.Since(validationStart)

	fmt.Printf(
		"[VALIDATION] success duration=%s\n",
		validationDuration,
	)

	handlerStart := time.Now()

	result, err := handler(c, req)

	handlerDuration := time.Since(handlerStart)

	if err != nil {
		totalDuration := time.Since(start)

		fmt.Printf(
			"[HANDLER] failed duration=%s total=%s error=%v\n",
			handlerDuration,
			totalDuration,
			err,
		)

		return err
	}

	totalDuration := time.Since(start)

	fmt.Printf(
		"[HANDLER] success duration=%s total=%s\n",
		handlerDuration,
		totalDuration,
	)

	return responseHandler.Handle(c, result)
}

func Handle[Req validation.Validatable, Res any](
	h Handler,
	handler HandlerFunc[Req, Res],
	status int,
	req Req,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		return handleRequest(
			c,
			req,
			handler,
			JSONResponseHandler{
				status: status,
			},
		)
	}
}
