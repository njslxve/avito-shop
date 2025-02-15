package shop

import (
	"log/slog"

	"github.com/njslxve/avito-shop/internal/model"
	"github.com/njslxve/avito-shop/internal/repository"
)

type ShopService struct {
	logger *slog.Logger
	repo   *repository.Repository
}

func New(logger *slog.Logger, repo *repository.Repository) *ShopService {
	return &ShopService{
		logger: logger,
		repo:   repo,
	}
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
