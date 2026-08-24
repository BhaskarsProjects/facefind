package auth

import "time"

type CreateUserParams struct {
	Auth0Sub          string
	Email             string
	EmailVerified     bool
	DisplayName       string
	ProfilePictureURL string
}

type AuthenticatedUser struct {
	Auth0UserID string
	UserID      int64
	Role        string
}

type User struct {
	ID        int64
	Auth0Sub  string
	Role      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
