package auth

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (p *LoginPayload) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

type SignupPayload struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=8"`
}

func (p *SignupPayload) Validate() error {
	validate := validator.New()

	if err := validate.Struct(p); err != nil {
		return err
	}

	if p.Password != p.ConfirmPassword {
		return errors.New("passwords do not match")
	}

	return nil
}

type LoginResponse struct {
	Message string `json:"message"`
}

type SignupResponse struct {
	Message string `json:"message"`
}
