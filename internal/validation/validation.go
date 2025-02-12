package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/njslxve/avito-shop/internal/model"
)

func ValdateAuthRequest(r model.AuthRequest) error {
	v := validator.New(validator.WithRequiredStructEnabled())

	return v.Struct(r)
}

func ValidateSendCoinRequest(r model.SendCoinRequest) error {
	v := validator.New(validator.WithRequiredStructEnabled())

	return v.Struct(r)
}

func ValidatePassword(u model.User, p string) bool {
	return u.Password == p
}
