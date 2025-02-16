package auth

import (
	"fmt"

	"github.com/njslxve/avito-shop/internal/model"
	"github.com/njslxve/avito-shop/internal/validation"
)

func (a *Auth) User(username, password string) (model.User, error) { //get or create
	user, err := a.repo.User.FindUserByName(username)
	if err != nil {
		user, err = a.createUser(username, password)
		if err != nil {
			return model.User{}, err
		}

		return user, nil
	}

	if !validation.ValidatePassword(user, password) {
		return model.User{}, fmt.Errorf("invalid password")
	}

	return user, nil
}
