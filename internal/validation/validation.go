package validation

import (
	"crypto/sha256"
	"fmt"

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
	hash := sha256.New()
	hash.Write([]byte(p))

	phash := fmt.Sprintf("%x", hash.Sum(nil))

	return u.PasswordHash == phash
}
