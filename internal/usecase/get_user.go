package usecase

import (
	"fmt"

	"github.com/njslxve/avito-shop/internal/model"
	"github.com/njslxve/avito-shop/internal/validation"
)

func (u *Usecase) User(username, password string) (model.User, error) { //get or create
	user, err := u.repo.User.FindUser(username)
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
