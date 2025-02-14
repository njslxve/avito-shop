package usecase

import (
	"fmt"

	"github.com/njslxve/avito-shop/internal/model"
	"github.com/njslxve/avito-shop/internal/validation"
)

func (u *Usecase) UserByName(username, password string) (model.User, error) { //get or create
	user, err := u.repo.User.FindUserByName(username)
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

func (u *Usecase) UserByID(useerID string) (model.User, error) {
	user, err := u.repo.User.FindUserByID(useerID)
	if err != nil {
		return model.User{}, err //TODO
	}

	return user, nil
}
