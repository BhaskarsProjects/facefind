package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/auth0/go-auth0/v2/authentication"
	"github.com/auth0/go-auth0/v2/authentication/database"
	"github.com/auth0/go-auth0/v2/authentication/oauth"
	"github.com/bhaskarvb28/facefind/internal/model/auth"
	"github.com/bhaskarvb28/facefind/internal/repository"
	"github.com/bhaskarvb28/facefind/internal/server"

	"github.com/auth0/go-jwt-middleware/v3/jwks"
	"github.com/auth0/go-jwt-middleware/v3/validator"
)

type AuthService struct {
	server             *server.Server
	client             *authentication.Authentication
	domain             string
	databaseConnection string
	clientID           string
	clientSecret       string
	callbackURL        string
	idTokenValidator   *validator.Validator
	userRepo           *repository.UserRepository
}

func NewIDTokenValidator(
	domain string,
	clientID string,
) (*validator.Validator, error) {
	issuerURL, err := url.Parse(
		fmt.Sprintf("https://%s/", domain),
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to parse Auth0 issuer URL: %w",
			err,
		)
	}

	provider, err := jwks.NewCachingProvider(
		jwks.WithIssuerURL(issuerURL),
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create Auth0 JWKS provider: %w",
			err,
		)
	}

	jwtValidator, err := validator.New(
		validator.WithKeyFunc(provider.KeyFunc),
		validator.WithAlgorithm(validator.RS256),
		validator.WithIssuer(issuerURL.String()),
		validator.WithAudience(clientID),

		// Allow small differences between Auth0 and our server clocks.
		validator.WithAllowedClockSkew(10*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create Auth0 JWT validator: %w",
			err,
		)
	}

	return jwtValidator, nil
}

func NewAuthService(
	s *server.Server,
	userRepo *repository.UserRepository,
) (*AuthService, error) {
	cfg := s.Config.Auth

	client, err := authentication.New(
		context.Background(),
		cfg.Domain,
		authentication.WithClientID(cfg.ClientID),
		authentication.WithClientSecret(cfg.ClientSecret),
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to initialize Auth0 client: %w",
			err,
		)
	}

	idTokenValidator, err := NewIDTokenValidator(
		cfg.Domain,
		cfg.ClientID,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to initialize ID token validator: %w",
			err,
		)
	}

	return &AuthService{
		server:             s,
		client:             client,
		domain:             cfg.Domain,
		databaseConnection: cfg.DatabaseConnection,
		clientID:           cfg.ClientID,
		clientSecret:       cfg.ClientSecret,
		callbackURL:        cfg.CallbackURL,
		idTokenValidator:   idTokenValidator,
		userRepo:           userRepo,
	}, nil
}

func (a *AuthService) Signup(
	ctx context.Context,
	email string,
	password string,
) (*auth.SignupResponse, error) {

	res, err := a.client.Database.Signup(
		ctx,
		database.SignupRequest{
			Connection: a.databaseConnection,
			Email:      email,
			Password:   password,
		},
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create Auth0 user: %w",
			err,
		)
	}

	err = a.userRepo.CreateUser(
		ctx,
		res.ID,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create local user: %w",
			err,
		)
	}

	return &auth.SignupResponse{
		Message: "Account created successfully",
	}, nil
}

func (a *AuthService) Login(
	ctx context.Context,
	email string,
	password string,
) (*auth.AuthenticatedUser, error) {

	// 1. Authenticate credentials with Auth0.
	tokenSet, err := a.client.OAuth.LoginWithPassword(
		ctx,
		oauth.LoginWithPasswordRequest{
			Username: email,
			Password: password,
		},
		oauth.IDTokenValidationOptions{},
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to authenticate with Auth0: %w",
			err,
		)
	}

	// 2. Validate the Auth0 ID token.
	claims, err := a.idTokenValidator.ValidateToken(
		ctx,
		tokenSet.IDToken,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to validate Auth0 ID token: %w",
			err,
		)
	}

	validatedClaims, ok := claims.(*validator.ValidatedClaims)
	if !ok {
		return nil, fmt.Errorf(
			"unexpected Auth0 ID token claims type",
		)
	}

	auth0Sub := validatedClaims.RegisteredClaims.Subject

	user, err := a.userRepo.GetByAuth0Sub(ctx, auth0Sub)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to get local user: %w",
			err,
		)
	}

	if user.Status != "active" {
		return nil, fmt.Errorf("user account is not active")
	}

	return &auth.AuthenticatedUser{
		Auth0UserID: auth0Sub,
		UserID:      user.ID,
		Role:        user.Role,
	}, nil
}

func (a *AuthService) GetUserInfo(
	ctx context.Context,
	accessToken string,
) (map[string]any, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("https://%s/userinfo", a.domain),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create userinfo request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to request user info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)

		return nil, fmt.Errorf(
			"userinfo request failed: status=%d body=%s",
			resp.StatusCode,
			string(body),
		)
	}

	var userInfo map[string]any

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %w", err)
	}

	return userInfo, nil
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
