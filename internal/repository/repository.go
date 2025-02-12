package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/njslxve/avito-shop/internal/model"
)

type Repository struct {
	User        UserRepositoryIterface
	Item        ItemRepositoryInterface
	Coin        CoinRepositoryInterface
	Transaction TransactionRepositoryInterface
}

type UserRepositoryIterface interface {
	Create(user model.User) error
	FindUser(username string) (model.User, error)
	UpdateUserCoins(user model.User, amount int64) error
}

type ItemRepositoryInterface interface {
	FindItem(itemname string) (model.Item, error)
}

type CoinRepositoryInterface interface {
	CreateTransfer(string, string, int) error
	SenderHistory(string) ([]model.Transaction, error)
	ReceiverHistory(string) ([]model.Transaction, error)
}

type TransactionRepositoryInterface interface {
	Create(string, string) error
	UserHistory(string) ([]string, error)
}

func New(db *pgx.Conn) *Repository {
	return &Repository{
		User:        newUserRepository(db),
		Item:        newItemRepository(db),
		Coin:        newCoinRepository(db),
		Transaction: newTransactionRepository(db),
	}
}

var (
	qb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)
