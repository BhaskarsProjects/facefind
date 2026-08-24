package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bhaskarvb28/facefind/internal/server"
	"github.com/labstack/echo/v4"
)

type InternalMiddleware struct {
	server *server.Server
}

func NewInternalMiddleware(
	server *server.Server,
) *InternalMiddleware {
	return &InternalMiddleware{
		server: server,
	}
}

func (m *InternalMiddleware) RequireWorkerAuth(
	next echo.HandlerFunc,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("===== WORKER MIDDLEWARE =====")
		fmt.Println("METHOD:", c.Request().Method)
		fmt.Println("PATH:", c.Request().URL.Path)

		authHeader := c.Request().Header.Get("Authorization")
		fmt.Printf("AUTH HEADER: %q\n", authHeader)

		const prefix = "Bearer "

		if !strings.HasPrefix(authHeader, prefix) {
			fmt.Println("STRINGS")
			return echo.NewHTTPError(
				http.StatusUnauthorized,
				"unauthorized",
			)
		}

		token := strings.TrimSpace(
			strings.TrimPrefix(authHeader, prefix),
		)

		fmt.Printf("TOKEN: %q\n", token)
		fmt.Printf("CONFIG TOKEN: %q\n", m.server.Config.AIWorker.Token)

		if token == "" {
			fmt.Println("TOKEN EMPTY")
			return echo.NewHTTPError(
				http.StatusUnauthorized,
				"unauthorized",
			)
		}

		if token != m.server.Config.AIWorker.Token {
			fmt.Println("TOKEN NOT EQUAL")
			return echo.NewHTTPError(
				http.StatusUnauthorized,
				"unauthorized",
			)
		}

		fmt.Println("WORKER AUTH PASSED")

		return next(c)
	}
}