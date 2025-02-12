package usecase

import (
	"fmt"
	"log/slog"

	"github.com/njslxve/avito-shop/internal/auth"
	"github.com/njslxve/avito-shop/internal/model"
	"github.com/njslxve/avito-shop/internal/repository"
	"github.com/njslxve/avito-shop/internal/validation"
)

type Usecase struct {
	logger *slog.Logger
	auth   *auth.Auth
	repo   *repository.Repository
}

func New(logger *slog.Logger, auth *auth.Auth, repo *repository.Repository) *Usecase {
	return &Usecase{
		logger: logger,
		auth:   auth,
		repo:   repo,
	}
}

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

func (u *Usecase) Token(user model.User) (string, error) {
	token, err := u.auth.GenerateToken(user.Username, user.Password)
	if err != nil {
		u.logger.Error("failed to generate token",
			slog.String("username", user.Username),
			slog.String("error", err.Error()),
		)
		return "", err
	}

	return token, nil
}

func (u *Usecase) ValidateItem(itemname string) bool {
	item, err := u.repo.Item.FindItem(itemname)
	if err != nil {
		u.logger.Error("failed to find item",
			slog.String("item", item.Type),
			slog.String("error", err.Error()),
		)

		return false
	}

	return true
}

func (u *Usecase) BuyItem(user model.User, itemname string) error {
	item, err := u.repo.Item.FindItem(itemname)
	if err != nil {
		u.logger.Error("failed to find item",
			slog.String("item", itemname),
			slog.String("error", err.Error()),
		)
	}

	if item.Price > user.Coins {
		return fmt.Errorf("not enough coins")
	}

	err = u.repo.User.UpdateUserCoins(user, -item.Price)
	if err != nil {
		u.logger.Error("failed to update user coins",
			slog.String("username", user.Username),
			slog.String("error", err.Error()),
		)
	}

	err = u.repo.Transaction.Create(user.ID, item.ID)
	if err != nil {
		u.logger.Error("failed to create transaction",
			slog.String("username", user.Username),
			slog.String("item", itemname),
			slog.String("error", err.Error()),
		)

		err = u.repo.User.UpdateUserCoins(user, item.Price)
		if err != nil {
			u.logger.Error("failed to refund user coins",
				slog.String("username", user.Username),
				slog.String("error", err.Error()),
			)
		}

		return err
	}

	u.logger.Info("item bought",
		slog.String("username", user.Username),
		slog.String("item", itemname),
	)

	return nil
}

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

func (u *Usecase) createUser(username, password string) (model.User, error) {
	user := model.User{
		Username: username,
		Password: password,
	}

	err := u.repo.User.Create(user)
	if err != nil {
		u.logger.Error("failed to create user",
			slog.String("username", username),
			slog.String("error", err.Error()),
		)
	}

	return user, err
}
