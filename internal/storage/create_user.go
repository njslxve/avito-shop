package storage

import (
	"context"
	"fmt"

	"github.com/njslxve/avito-shop/internal/model"
)

func (s *Storage) CreateUser(username, password string) error {
	const op = "storage.CreateUser"

	user := model.User{
		Username: username,
		Password: password,
		Coins:    1000,
	}

	querry := qb.Insert("users").
		Columns("username", "pass", "coins").
		Values(user.Username, user.Password, user.Coins)

	sql, args, err := querry.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = s.db.Exec(context.Background(), sql, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
