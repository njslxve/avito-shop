package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepository struct {
	db *pgxpool.Pool
}

func newTransactionRepository(db *pgxpool.Pool) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (tr *TransactionRepository) Create(user, item string) error {
	const op = "repository.CreateTransaction"

	querry := qb.Insert("item_transactions").
		Columns("user_id", "item_id").
		Values(user, item)

	sql, args, err := querry.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = tr.db.Exec(context.Background(), sql, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (tr *TransactionRepository) UserHistory(user string) ([]string, error) {
	op := "repository.UserHistory"

	var items []string

	querry := qb.Select("i.type").
		From("item_transactions t").
		Join("items i ON t.item_id = i.id").
		Where(sq.Eq{"t.user_id": user})

	sql, args, err := querry.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := tr.db.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for rows.Next() {
		var item string

		err = rows.Scan(&item)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		items = append(items, item)
	}

	return items, nil
}
