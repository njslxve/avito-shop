package usecase

import (
	"log/slog"

	"github.com/njslxve/avito-shop/internal/model"
)

func (u *Usecase) createUser(username, password string) (model.User, error) {
	user := model.User{
		Username: username,
		Password: password,
	}

	err := u.repo.User.Create(user)
	if err != nil {
		u.logger.Error("failed to create user",
			slog.String("username", username),
			slog.String("error", err.Error()),
		)
	}

	return user, err
}
