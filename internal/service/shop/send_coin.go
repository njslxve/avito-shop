package shop

import (
	"fmt"

	"github.com/njslxve/avito-shop/internal/model"
)

func (ss *ShopService) SendCoin(sender model.User, receiverUsername string, amount int64) error {
	if amount > sender.Coins {
		return fmt.Errorf("not enough coins")
	}

	err := ss.repo.User.UpdateUserCoins(sender, -amount)
	if err != nil {
		return err
	}

	reciver, err := ss.repo.User.FindUserByName(receiverUsername)
	if err != nil {
		return err
	}

	err = ss.repo.User.UpdateUserCoins(reciver, amount)
	if err != nil {
		return err
	}

	err = ss.repo.Coin.CreateTransfer(sender.ID, reciver.ID, amount)
	if err != nil {
		return err
	}

	return nil //TODO
}
