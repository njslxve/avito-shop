package usecase

import (
	"log/slog"

	"github.com/njslxve/avito-shop/internal/auth"
	"github.com/njslxve/avito-shop/internal/model"
	"github.com/njslxve/avito-shop/internal/storage"
)

type Usecase struct {
	logger  *slog.Logger
	auth    *auth.Auth
	storage *storage.Storage
}

func New(logger *slog.Logger, auth *auth.Auth, storage *storage.Storage) *Usecase {
	return &Usecase{
		logger:  logger,
		auth:    auth,
		storage: storage,
	}
}

func (u *Usecase) User(username, password string) (model.User, error) {

	return model.User{}, nil
}

func (u *Usecase) Token(user model.User) (string, error) {

	return "", nil
}
