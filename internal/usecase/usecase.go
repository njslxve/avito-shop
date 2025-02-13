package usecase

import (
	"log/slog"

	"github.com/njslxve/avito-shop/internal/auth"
	"github.com/njslxve/avito-shop/internal/model"
	"github.com/njslxve/avito-shop/internal/repository"
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

func aggregatePurshases(purshases []string) []model.ItemInfo {
	m := make(map[string]int64)

	for _, purshase := range purshases {
		m[purshase]++
	}

	var itemInfo []model.ItemInfo

	for item, quantity := range m {
		itemInfo = append(itemInfo, model.ItemInfo{
			Type:     item,
			Quantity: quantity,
		})
	}

	return itemInfo
}

func aggregateSender(senderHist []model.Transaction) []model.Sent {
	var sent []model.Sent

	for _, hist := range senderHist {
		sent = append(sent, model.Sent{
			ToUser: hist.Username,
			Amount: hist.Amount,
		})
	}

	return sent
}

func aggregateReceiver(receiverHist []model.Transaction) []model.Received {
	var received []model.Received

	for _, hist := range receiverHist {
		received = append(received, model.Received{
			FromUser: hist.Username,
			Amount:   hist.Amount,
		})
	}

	return received
}
