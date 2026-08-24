package folder

import (
	"github.com/go-playground/validator/v10"
)

type CreatePayload struct {
	Name string `json:"name" validate:"required,min=1,max=50"`
}

func (p *CreatePayload) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

type UpdatePayload struct {
	Name string `json:"name" validate:"required"`
}

func (p *UpdatePayload) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
