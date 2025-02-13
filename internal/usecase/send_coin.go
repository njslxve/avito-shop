package usecase

import (
	"fmt"

	"github.com/njslxve/avito-shop/internal/model"
)

func (u *Usecase) SendCoin(sender model.User, receiverUsername string, amount int64) error {
	if amount > sender.Coins {
		return fmt.Errorf("not enough coins")
	}

	err := u.repo.User.UpdateUserCoins(sender, -amount)
	if err != nil {
		return err
	}

	reciver, err := u.repo.User.FindUser(receiverUsername)
	if err != nil {
		return err
	}

	err = u.repo.User.UpdateUserCoins(reciver, amount)
	if err != nil {
		return err
	}

	err = u.repo.Coin.CreateTransfer(sender.ID, reciver.ID, amount)
	if err != nil {
		return err
	}

	return nil //TODO
}
