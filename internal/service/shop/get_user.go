package shop

import (
	"github.com/njslxve/avito-shop/internal/model"
)

func (ss *ShopService) User(userID string) (model.User, error) {
	user, err := ss.repo.User.FindUserByID(userID)
	if err != nil {
		return model.User{}, err
	}

	return user, nil //TODO
}
