package shop

import (
	"log/slog"

	"github.com/njslxve/avito-shop/internal/model"
)

func (ss *ShopService) User(userID string) (model.User, error) {
	const op = "shop.User"
	user, err := ss.repo.User.FindUserByID(userID)
	if err != nil {
		ss.logger.Error("failed to find user",
			slog.String("operation", op),
			slog.String("error", err.Error()),
		)

		return model.User{}, err
	}

	return user, nil
}
