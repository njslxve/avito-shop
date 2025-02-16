package shop

import (
	"fmt"
	"log/slog"

	"github.com/njslxve/avito-shop/internal/model"
)

func (ss *ShopService) SendCoin(sender model.User, receiverUsername string, amount int64) error {
	const op = "shop.SendCoin"

	if amount > sender.Coins {
		return fmt.Errorf("not enough coins")
	}

	err := ss.repo.User.UpdateUserCoins(sender, -amount)
	if err != nil {
		ss.logger.Error("failed to update user coins",
			slog.String("operation", op),
			slog.String("username", sender.Username),
			slog.String("error", err.Error()),
		)

		return err
	}

	reciver, err := ss.repo.User.FindUserByName(receiverUsername)
	if err != nil {
		ss.logger.Error("failed to find user",
			slog.String("operation", op),
			slog.String("username", receiverUsername),
			slog.String("error", err.Error()),
		)
		return err
	}

	err = ss.repo.User.UpdateUserCoins(reciver, amount)
	if err != nil {
		ss.logger.Error("failed to update user coins",
			slog.String("operation", op),
			slog.String("username", sender.Username),
			slog.String("error", err.Error()),
		)

		return err
	}

	err = ss.repo.Coin.CreateTransfer(sender.ID, reciver.ID, amount)
	if err != nil {
		ss.logger.Error("failed to create transfer",
			slog.String("operation", op),
			slog.String("sender", sender.Username),
			slog.String("receiver", reciver.Username),
			slog.String("error", err.Error()),
		)

		return err
	}

	return nil
}
