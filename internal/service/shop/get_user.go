package shop

import (
	"github.com/njslxve/avito-shop/internal/model"
)

func (ss *ShopService) User(userID string) (model.User, error) { //get or create
	user, err := ss.repo.User.FindUserByID(userID)
	if err != nil {
		return user, nil
	}

	// if !validation.ValidatePassword(user, password) {
	// 	return model.User{}, fmt.Errorf("invalid password") //TODO
	// }

	return user, nil
}
