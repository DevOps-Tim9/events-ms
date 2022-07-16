package dto

import (
	"github.com/go-playground/validator"
)

type EventRequestDTO struct {
	Timestamp string `validate:"required"`
	Message   string `validate:"required"`
}

func (u *EventRequestDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
