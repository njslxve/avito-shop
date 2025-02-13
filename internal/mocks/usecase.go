package mocks

import (
	"github.com/njslxve/avito-shop/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockUsecase struct {
	mock.Mock
}

func (mu *MockUsecase) ValidateItem(itemname string) bool {
	args := mu.Called(itemname)

	return args.Bool(0)
}

func (mu *MockUsecase) User(username, password string) (model.User, error) {
	args := mu.Called(username, password)

	return args.Get(0).(model.User), args.Error(1)
}

func (mu *MockUsecase) Token(user model.User) (string, error) {
	args := mu.Called(user)

	return args.String(0), args.Error(1)
}

func (mu *MockUsecase) BuyItem(user model.User, item string) error {
	args := mu.Called(user, item)

	return args.Error(0)
}

func (mu *MockUsecase) SendCoin(sender model.User, to string, amount int64) error {
	args := mu.Called(sender, to, amount)

	return args.Error(0)
}
