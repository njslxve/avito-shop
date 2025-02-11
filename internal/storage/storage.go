package storage

import (
	"log/slog"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/njslxve/avito-shop/internal/client/warehouse"
)

type Storage struct {
	logger *slog.Logger
	db     *pgx.Conn
	wh     *warehouse.Warehouse
}

func New(logger *slog.Logger, db *pgx.Conn, wh *warehouse.Warehouse) *Storage {
	return &Storage{
		logger: logger,
		db:     db,
		wh:     wh,
	}
}

var (
	qb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)
