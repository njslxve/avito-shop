package shop

import (
	"fmt"
	"log/slog"

	"github.com/njslxve/avito-shop/internal/model"
)

func (ss *ShopService) BuyItem(user model.User, itemname string) error {
	item, err := ss.repo.Item.FindItem(itemname)
	if err != nil {
		ss.logger.Error("failed to find item",
			slog.String("item", itemname),
			slog.String("error", err.Error()),
		)
	}

	if item.Price > user.Coins {
		return fmt.Errorf("not enough coins")
	}

	err = ss.repo.User.UpdateUserCoins(user, -item.Price)
	if err != nil {
		ss.logger.Error("failed to update user coins",
			slog.String("username", user.Username),
			slog.String("error", err.Error()),
		)
	}

	err = ss.repo.Transaction.Create(user.ID, item.ID)
	if err != nil {
		ss.logger.Error("failed to create transaction",
			slog.String("username", user.Username),
			slog.String("item", itemname),
			slog.String("error", err.Error()),
		)

		err = ss.repo.User.UpdateUserCoins(user, item.Price)
		if err != nil {
			ss.logger.Error("failed to refund user coins",
				slog.String("username", user.Username),
				slog.String("error", err.Error()),
			)
		}

		return err
	}

	return nil
}
