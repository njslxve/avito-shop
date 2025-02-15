package auth

import (
	"crypto/sha256"
	"fmt"
	"log/slog"

	"github.com/njslxve/avito-shop/internal/model"
)

func (a *Auth) createUser(username, password string) (model.User, error) {
	pwd := passwordHash(password)

	user := model.User{
		Username:     username,
		PasswordHash: pwd,
	}

	id, err := a.repo.User.Create(user)
	if err != nil {
		a.logger.Error("failed to create user",
			slog.String("username", username),
			slog.String("error", err.Error()),
		)
	}

	user.ID = id

	return user, err
}

func passwordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum(nil))
}
