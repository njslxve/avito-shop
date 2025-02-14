package mocks

import (
	"github.com/njslxve/avito-shop/internal/model"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

type MockUserRepository struct {
	mock.Mock
}

func (mu *MockUserRepository) Create(user model.User) (string, error) {
	args := mu.Called(user)

	return args.String(0), args.Error(1)
}

func (mu *MockUserRepository) FindUserByName(username string) (model.User, error) {
	args := mu.Called(username)

	return args.Get(0).(model.User), args.Error(1)
}

func (mu *MockUserRepository) FindUserByID(userID string) (model.User, error) {
	args := mu.Called(userID)

	return args.Get(0).(model.User), args.Error(1)
}

func (mu *MockUserRepository) UpdateUserCoins(user model.User, amount int64) error {
	args := mu.Called(user, amount)

	return args.Error(0)
}

type MockItemRepository struct {
	mock.Mock
}

func (mi *MockItemRepository) FindItem(itemname string) (model.Item, error) {
	args := mi.Called(itemname)

	return args.Get(0).(model.Item), args.Error(1)
}

type MockTransactionRepository struct {
	mock.Mock
}

func (mt *MockTransactionRepository) Create(user, item string) error {
	args := mt.Called(user, item)

	return args.Error(0)
}

func (mt *MockTransactionRepository) UserHistory(user string) ([]string, error) {
	args := mt.Called(user)

	return args.Get(0).([]string), args.Error(1)
}

type MockCoinRepository struct {
	mock.Mock
}

func (mc *MockCoinRepository) CreateTransfer(from string, to string, amount int64) error {
	args := mc.Called(from, to, amount)

	return args.Error(0)
}

func (mc *MockCoinRepository) SenderHistory(sender string) ([]model.Transaction, error) {
	args := mc.Called(sender)

	return args.Get(0).([]model.Transaction), args.Error(1)
}

func (mc *MockCoinRepository) ReceiverHistory(receiver string) ([]model.Transaction, error) {
	args := mc.Called(receiver)

	return args.Get(0).([]model.Transaction), args.Error(1)
}
