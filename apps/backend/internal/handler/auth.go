package handler

import (
	"net/http"

	"github.com/bhaskarvb28/facefind/internal/server"
	"github.com/bhaskarvb28/facefind/internal/service"
	"github.com/bhaskarvb28/facefind/internal/utils"
)

type AuthHandler struct {
	Handler
	authService *service.AuthService
}

func NewAuthHandler(s *server.Server, authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		Handler: NewHandler(s),
		authService: authService,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	state, err := utils.GenerateState()
	if err != nil {
		http.Error(w, "failed to generate state", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "oauth_state",
		Value:    state,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // false locally if you're using plain HTTP
		SameSite: http.SameSiteLaxMode,
		MaxAge:   300, // 5 minutes
	})

	authURL := h.authService.AuthorizationURL(state)

	http.Redirect(w, r, authURL, http.StatusFound)
}

func (h *AuthHandler) Callback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	if code == "" {
		http.Error(w, "missing authorization code", http.StatusBadRequest)
		return
	}

	if state == "" {
		http.Error(w, "missing state", http.StatusBadRequest)
		return
	}

	stateCookie, err := r.Cookie("oauth_state")
	if err != nil {
		http.Error(w, "missing oauth state", http.StatusBadRequest)
		return
	}

	if state != stateCookie.Value {
		http.Error(w, "invalid oauth state", http.StatusBadRequest)
		return
	}

	tokenSet, err := h.authService.ExchangeCode(r.Context(), code)
	if err != nil {
		http.Error(w, "failed to authenticate", http.StatusUnauthorized)
		return
	}

	// Delete OAuth state cookie.
	http.SetCookie(w, &http.Cookie{
		Name:     "oauth_state",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	// Store tokens in cookies.
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    tokenSet.AccessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "id_token",
		Value:    tokenSet.IDToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	// Redirect to frontend.
	http.Redirect(
		w,
		r,
		"http://localhost:3000",
		http.StatusFound,
	)
}