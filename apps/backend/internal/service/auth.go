package service

import (
	"context"
	"fmt"
	"net/url"

	"github.com/auth0/go-auth0/v2/authentication"
	"github.com/auth0/go-auth0/v2/authentication/oauth"
	"github.com/bhaskarvb28/facefind/internal/server"
)

type AuthService struct {
	server       *server.Server
	client       *authentication.Authentication
	domain       string
	clientID     string
	clientSecret string
	callbackURL  string
}

func NewAuthService(s *server.Server) (*AuthService, error) {
	cfg := s.Config.Auth

	client, err := authentication.New(
		context.Background(),
		cfg.Domain,
		authentication.WithClientID(cfg.ClientID),
		authentication.WithClientSecret(cfg.ClientSecret),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Auth0 client: %w", err)
	}

	return &AuthService{
		server:       s,
		client:       client,
		domain:       cfg.Domain,
		clientID:     cfg.ClientID,
		clientSecret: cfg.ClientSecret,
		callbackURL:  cfg.CallbackURL,
	}, nil
}

func (a *AuthService) AuthorizationURL(state string) string {
	u := url.URL{
		Scheme: "https",
		Host:   a.domain,
		Path:   "/authorize",
	}

	q := u.Query()
	q.Set("response_type", "code")
	q.Set("client_id", a.clientID)
	q.Set("redirect_uri", a.callbackURL)
	q.Set("scope", "openid profile email")
	q.Set("state", state)

	u.RawQuery = q.Encode()

	return u.String()
}

func (a *AuthService) ExchangeCode(ctx context.Context, code string) (*oauth.TokenSet, error) {
	tokenSet, err := a.client.OAuth.LoginWithAuthCode(
		ctx,
		oauth.LoginWithAuthCodeRequest{
			Code:        code,
			RedirectURI: a.callbackURL,
		},
		oauth.IDTokenValidationOptions{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange authorization code: %w", err)
	}

	return tokenSet, nil
}