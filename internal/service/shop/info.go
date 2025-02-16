package shop

import (
	"log/slog"

	"github.com/njslxve/avito-shop/internal/model"
)

func (ss *ShopService) Info(user model.User) (model.InfoResponse, error) {
	const op = "shop.Info"

	var infoResponse model.InfoResponse

	purshases, err := ss.repo.Transaction.UserHistory(user.ID)
	if err != nil {
		ss.logger.Error("failed to get user history",
			slog.String("operation", op),
			slog.String("error", err.Error()),
		)

		return model.InfoResponse{}, err
	}

	itemInfo := aggregatePurshases(purshases)

	senderHist, err := ss.repo.Coin.SenderHistory(user.ID)
	if err != nil {
		ss.logger.Error("failed to get sender history",
			slog.String("operation", op),
			slog.String("error", err.Error()),
		)

		return model.InfoResponse{}, err
	}

	receiverHist, err := ss.repo.Coin.ReceiverHistory(user.ID)
	if err != nil {
		ss.logger.Error("failed to get receiver history",
			slog.String("operation", op),
			slog.String("error", err.Error()),
		)

		return model.InfoResponse{}, err
	}

	sender := aggregateSender(senderHist)
	receiver := aggregateReceiver(receiverHist)

	infoResponse = model.InfoResponse{
		Coins:     user.Coins,
		Inventory: itemInfo,
		CoinHistory: model.History{
			Received: receiver,
			Sent:     sender,
		},
	}

	return infoResponse, nil
}
