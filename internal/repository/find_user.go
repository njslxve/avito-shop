package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/njslxve/avito-shop/internal/model"
)

func (s *Repository) FindUser(username string) (model.User, error) {
	const op = "storage.FindUser"

	var user model.User

	querry := qb.Select("username", "pass", "coins").
		From("users").
		Where(sq.Eq{"username": username})

	sql, args, err := querry.ToSql()
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}

	row := s.db.QueryRow(context.Background(), sql, args...)

	err = row.Scan(&user.Username, &user.Password, &user.Coins)
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
