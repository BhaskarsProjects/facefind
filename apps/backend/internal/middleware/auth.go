package middleware

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v3/jwks"
	"github.com/auth0/go-jwt-middleware/v3/validator"
	"github.com/bhaskarvb28/facefind/internal/server"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	server         *server.Server
	tokenValidator *validator.Validator
}

func NewAuthMiddleware(s *server.Server) (*AuthMiddleware, error) {
	issuerURL, err := url.Parse(
		fmt.Sprintf("https://%s/", s.Config.Auth.Domain),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Auth0 issuer URL: %w", err)
	}

	provider, err := jwks.NewCachingProvider(
		jwks.WithIssuerURL(issuerURL),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Auth0 JWKS provider: %w", err)
	}

	tokenValidator, err := validator.New(
		validator.WithKeyFunc(provider.KeyFunc),
		validator.WithAlgorithm(validator.RS256),
		validator.WithIssuer(issuerURL.String()),
		validator.WithAudience(s.Config.Auth.Audience),
		validator.WithAllowedClockSkew(10*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create access token validator: %w", err)
	}

	return &AuthMiddleware{
		server:         s,
		tokenValidator: tokenValidator,
	}, nil
}

func (auth *AuthMiddleware) RequireAuth(
	next echo.HandlerFunc,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("AUTH MIDDLEWARE HIT")
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			fmt.Println("No auth header")
			return echo.NewHTTPError(
				http.StatusUnauthorized,
				"Unauthorized",
			)
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 ||
			!strings.EqualFold(parts[0], "Bearer") {
			fmt.Println("No valid syntax of token")
			return echo.NewHTTPError(
				http.StatusUnauthorized,
				"Unauthorized",
			)
		}

		claims, err := auth.tokenValidator.ValidateToken(
			c.Request().Context(),
			parts[1],
		)
		if err != nil {
			fmt.Println("Invalid token", err)
			return echo.NewHTTPError(
			
				http.StatusUnauthorized,
				"Unauthorized",
			)
		}

		validatedClaims, ok := claims.(*validator.ValidatedClaims)
		if !ok {
			fmt.Println("Claims error")
			return echo.NewHTTPError(
				http.StatusUnauthorized,
				"Unauthorized",
			)
		}

		sub := validatedClaims.RegisteredClaims.Subject

		if sub == "" {
			fmt.Println("No sub")

			return echo.NewHTTPError(
				http.StatusUnauthorized,
				"Unauthorized",
			)
		}

		c.Set("user_sub", sub)

		fmt.Println("AUTH MIDDLEWARE SUCCESS")

		return next(c)
	}
}
