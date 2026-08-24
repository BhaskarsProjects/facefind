package handler

import (
	"net/http"

	"github.com/bhaskarvb28/facefind/internal/model/auth"
	"github.com/bhaskarvb28/facefind/internal/server"
	"github.com/bhaskarvb28/facefind/internal/service"
	"github.com/bhaskarvb28/facefind/internal/utils"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	Handler
	authService *service.AuthService
}

func NewAuthHandler(
	s *server.Server,
	authService *service.AuthService,
) *AuthHandler {
	return &AuthHandler{
		Handler:     NewHandler(s),
		authService: authService,
	}
}

func (h *AuthHandler) Signup(c echo.Context) error {
	return Handle(
		h.Handler,
		func(c echo.Context, payload *auth.SignupPayload) (*auth.SignupResponse, error) {

			response, err := h.authService.Signup(
				c.Request().Context(),
				payload.Email,
				payload.Password,
			)
			if err != nil {
				return nil, err
			}

			return response, nil
		},
		http.StatusCreated,
		&auth.SignupPayload{},
	)(c)
}

// 1. Get email/password
// 2. Authenticate user through authService
// 3. Create server-side session
// 4. Set session cookie
// 5. Return user

func (h *AuthHandler) Login(c echo.Context) error {
	return Handle(
		h.Handler,
		func(
			c echo.Context,
			payload *auth.LoginPayload,
		) (*auth.AuthenticatedUser, error) {

			response, err := h.authService.Login(
				c.Request().Context(),
				payload.Email,
				payload.Password,
			)
			if err != nil {
				return nil, err
			}

			return response, nil
		},
		http.StatusOK,
		&auth.LoginPayload{},
	)(c)
}

func (h *AuthHandler) OAuthLogin(c echo.Context) error {
	state, err := utils.GenerateState()
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"failed to generate state",
		)
	}

	cookie := new(http.Cookie)
	cookie.Name = "oauth_state"
	cookie.Value = state
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = false // false locally when using HTTP
	cookie.SameSite = http.SameSiteLaxMode
	cookie.MaxAge = 300

	c.SetCookie(cookie)

	authURL := h.authService.AuthorizationURL(state)

	return c.Redirect(http.StatusFound, authURL)
}

func (h *AuthHandler) Callback(c echo.Context) error {
	code := c.QueryParam("code")
	state := c.QueryParam("state")

	if code == "" {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"missing authorization code",
		)
	}

	if state == "" {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"missing state",
		)
	}

	stateCookie, err := c.Cookie("oauth_state")
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"missing oauth state",
		)
	}

	if state != stateCookie.Value {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"invalid oauth state",
		)
	}

	tokenSet, err := h.authService.ExchangeCode(
		c.Request().Context(),
		code,
	)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusUnauthorized,
			"failed to authenticate",
		)
	}

	// Delete OAuth state cookie.
	c.SetCookie(&http.Cookie{
		Name:     "oauth_state",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	// Store access token.
	c.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    tokenSet.AccessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	// Store ID token.
	c.SetCookie(&http.Cookie{
		Name:     "id_token",
		Value:    tokenSet.IDToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	// Redirect to frontend.
	return c.Redirect(
		http.StatusFound,
		"http://localhost:3000",
	)
}
