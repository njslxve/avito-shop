package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/njslxve/avito-shop/internal/model"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func newUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(user model.User) error {
	const op = "repository.CreateUser"

	querry := qb.Insert("users").
		Columns("username", "pass").
		Values(user.Username, user.Password)

	sql, args, err := querry.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = ur.db.Exec(context.Background(), sql, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (ur *UserRepository) FindUser(username string) (model.User, error) {
	const op = "repository.FindUser"

	var user model.User

	querry := qb.Select("id", "username", "pass", "coins").
		From("users").
		Where(sq.Eq{"username": username})

	sql, args, err := querry.ToSql()
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}

	row := ur.db.QueryRow(context.Background(), sql, args...)

	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Coins)
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (ur *UserRepository) UpdateUserCoins(user model.User, amount int64) error {
	const op = "repository.UpdateUser"

	querry := qb.Update("users").
		Set("coins", user.Coins+amount).
		Where(sq.Eq{"username": user.Username})

	sql, args, err := querry.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = ur.db.Exec(context.Background(), sql, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
