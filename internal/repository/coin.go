package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/njslxve/avito-shop/internal/model"
)

type CoinRepository struct {
	db *pgx.Conn
}

func newCoinRepository(db *pgx.Conn) *CoinRepository {
	return &CoinRepository{
		db: db,
	}
}

func (cr *CoinRepository) CreateTransfer(from string, to string, amount int) error {
	const op = "repository.TransferCoins"

	querry := qb.Insert("user_transactions").
		Columns("from_user_id", "to_user_id", "amount").
		Values(from, to, amount)

	sql, args, err := querry.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = cr.db.Exec(context.Background(), sql, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (cr *CoinRepository) SenderHistory(sender string) ([]model.Transaction, error) {
	const op = "repository.HistoryBySender"

	var transactions []model.Transaction

	querry := qb.Select("u.username", "t.amount").
		From("user_transactions t").
		Join("users u ON t.to_user_id = u.id").
		Where(sq.Eq{"t.from_user_id": sender}).
		OrderBy("t.created_at DESC")

	sql, args, err := querry.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := cr.db.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for rows.Next() {
		var t model.Transaction

		err = rows.Scan(&t.Username, &t.Amount)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		transactions = append(transactions, t)
	}

	return transactions, nil
}

func (cr *CoinRepository) ReceiverHistory(receiver string) ([]model.Transaction, error) {
	const op = "repository.HistoryByReceiver"

	var transactions []model.Transaction

	querry := qb.Select("u.username", "t.amount").
		From("user_transactions t").
		Join("users u ON t.from_user_id = u.id").
		Where(sq.Eq{"t.to_user_id": receiver}).
		OrderBy("t.created_at DESC")

	sql, args, err := querry.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := cr.db.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for rows.Next() {
		var t model.Transaction

		err = rows.Scan(&t.Username, &t.Amount)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		transactions = append(transactions, t)
	}

	return transactions, nil
}
