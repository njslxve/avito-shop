package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/njslxve/avito-shop/internal/model"
)

type ItemRepository struct {
	db *pgx.Conn
}

func newItemRepository(db *pgx.Conn) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}

func (ir *ItemRepository) FindItem(itemname string) (model.Item, error) {
	const op = "storage.FindItem"

	var item model.Item

	querry := qb.Select("id", "type", "price").
		From("items").
		Where(sq.Eq{"type": itemname})

	sql, args, err := querry.ToSql()
	if err != nil {
		return model.Item{}, fmt.Errorf("%s: %w", op, err)
	}

	row := ir.db.QueryRow(context.Background(), sql, args...)

	err = row.Scan(&item.ID, &item.Type, &item.Price)
	if err != nil {
		return model.Item{}, fmt.Errorf("%s: %w", op, err)
	}

	return item, nil
}
