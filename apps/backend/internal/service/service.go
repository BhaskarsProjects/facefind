package service

import "github.com/bhaskarvb28/facefind/internal/server"

type Services struct {
	Auth *AuthService
}

func NewServices(s *server.Server) (*Services, error) {
	authService, err := NewAuthService(s)
	if err != nil {
		return nil, err
	}

	return &Services{
		Auth: authService,
	}, nil
}
