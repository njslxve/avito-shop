package mocks

import (
	"github.com/njslxve/avito-shop/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockShopService struct {
	mock.Mock
}

func (ms *MockShopService) ValidateItem(itemname string) bool {
	args := ms.Called(itemname)

	return args.Bool(0)
}

func (ms *MockShopService) User(userID string) (model.User, error) {
	args := ms.Called(userID)

	return args.Get(0).(model.User), args.Error(1)
}

func (ms *MockShopService) BuyItem(user model.User, item string) error {
	args := ms.Called(user, item)

	return args.Error(0)
}

func (ms *MockShopService) SendCoin(sender model.User, to string, amount int64) error {
	args := ms.Called(sender, to, amount)

	return args.Error(0)
}

func (ms *MockShopService) Info(user model.User) (model.InfoResponse, error) {
	args := ms.Called(user)

	return args.Get(0).(model.InfoResponse), args.Error(1)
}

type MockAuthService struct {
	mock.Mock
}

func (ma *MockAuthService) User(username string, password string) (model.User, error) {
	args := ma.Called(username, password)

	return args.Get(0).(model.User), args.Error(1)
}

func (ma *MockAuthService) Token(userID string) (string, error) {
	args := ma.Called(userID)

	return args.String(0), args.Error(1)
}
