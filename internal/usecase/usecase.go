package usecase

import (
	"fmt"
	"log/slog"

	"github.com/njslxve/avito-shop/internal/auth"
	"github.com/njslxve/avito-shop/internal/model"
	"github.com/njslxve/avito-shop/internal/repository"
	"github.com/njslxve/avito-shop/internal/validation"
)

type Usecase struct {
	logger *slog.Logger
	auth   *auth.Auth
	repo   *repository.Repository
}

func New(logger *slog.Logger, auth *auth.Auth, repo *repository.Repository) *Usecase {
	return &Usecase{
		logger: logger,
		auth:   auth,
		repo:   repo,
	}
}

func (u *Usecase) User(username, password string) (model.User, error) { //get or create
	user, err := u.repo.FindUser(username)
	if err != nil {
		user, err = u.createUser(username, password)
		if err != nil {
			return model.User{}, err
		}

		return user, nil
	}

	if !validation.ValidatePassword(user, password) {
		return model.User{}, fmt.Errorf("invalid password") //TODO
	}

	return user, nil
}

func (u *Usecase) Token(user model.User) (string, error) {
	token, err := u.auth.GenerateToken(user.Username, user.Password)
	if err != nil {
		u.logger.Error("failed to generate token",
			slog.String("username", user.Username),
			slog.String("error", err.Error()),
		)
		return "", err
	}

	return token, nil
}

func (u *Usecase) ValidateItem(item string) bool {
	if err := u.repo.FindItem(item); err != nil {
		u.logger.Error("failed to find item",
			slog.String("item", item),
			slog.String("error", err.Error()),
		)

		return false
	}

	return true
}

func (u *Usecase) BuyItem(user model.User, item string) error {
	// TODO
	return nil
}

func (u *Usecase) createUser(username, password string) (model.User, error) {
	user := model.User{
		Username: username,
		Password: password,
	}

	err := u.repo.CreateUser(user)
	if err != nil {
		u.logger.Error("failed to create user",
			slog.String("username", username),
			slog.String("error", err.Error()),
		)
	}

	return user, err
}
