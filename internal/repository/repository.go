package repository

import (
	"log/slog"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/njslxve/avito-shop/internal/client/warehouse"
)

type Repository struct {
	logger *slog.Logger
	db     *pgx.Conn
	wh     *warehouse.Warehouse
}

func New(logger *slog.Logger, db *pgx.Conn, wh *warehouse.Warehouse) *Repository {
	return &Repository{
		logger: logger,
		db:     db,
		wh:     wh,
	}
}

var (
	qb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)
