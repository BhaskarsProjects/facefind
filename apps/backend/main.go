package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/static", "static")

	e.GET("/docs", func(c echo.Context) error {
		return c.File("static/openapi.html")
	})

	e.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"status":      "healthy",
			"timestamp":   "2026-08-15T00:00:00Z",
			"environment": "development",
			"checks": map[string]any{
				"database": map[string]any{
					"status":        "healthy",
					"response_time": "10ms",
				},
			},
		})
	})

	e.GET("/name/:name", func(c echo.Context) error {
		name := c.Param("name")
		return c.JSON(http.StatusOK, "Hello " + name)
	})

	e.Logger.Fatal(e.Start(":8080"))
}