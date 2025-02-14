package usecase

import (
	"testing"

	"github.com/njslxve/avito-shop/internal/mocks"
	"github.com/njslxve/avito-shop/internal/model"
	"github.com/njslxve/avito-shop/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSendCoin(t *testing.T) {
	mockCoinRepo := new(mocks.MockCoinRepository)
	mockUserRepo := new(mocks.MockUserRepository)

	testuser := model.User{
		Username:     "testuser",
		PasswordHash: "testpass",
		Coins:        1000,
	}

	testResiver := "testreceiver"
	var testAmount int64 = 200

	u := New(nil, nil, &repository.Repository{
		Coin: mockCoinRepo,
		User: mockUserRepo,
	})

	mockUserRepo.On("FindUserByName", mock.Anything).Return(model.User{
		Username:     testResiver,
		PasswordHash: "tstpass",
		Coins:        100,
	}, nil)

	mockUserRepo.On("UpdateUserCoins", mock.Anything, mock.Anything).Return(nil)

	mockCoinRepo.On("CreateTransfer", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	err := u.SendCoin(testuser, testResiver, testAmount)

	assert.NoError(t, err)
}
