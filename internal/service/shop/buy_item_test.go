package shop

import (
	"testing"

	"github.com/njslxve/avito-shop/internal/mocks"
	"github.com/njslxve/avito-shop/internal/model"
	"github.com/njslxve/avito-shop/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBuyItem(t *testing.T) {
	mockTransRepo := new(mocks.MockTransactionRepository)
	mockCoinRepo := new(mocks.MockCoinRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	mockItemRepo := new(mocks.MockItemRepository)

	testuser := model.User{
		Username:     "testuser",
		PasswordHash: "testpassHash",
		Coins:        100,
	}

	testitem := "testitem"

	u := New(nil, &repository.Repository{
		Transaction: mockTransRepo,
		Coin:        mockCoinRepo,
		User:        mockUserRepo,
		Item:        mockItemRepo,
	})

	mockItemRepo.On("FindItem", mock.Anything).Return(model.Item{
		Type:  "testitem",
		Price: 10,
	}, nil)

	mockUserRepo.On("UpdateUserCoins", mock.Anything, mock.Anything).Return(nil)

	mockTransRepo.On("Create", mock.Anything, mock.Anything).Return(nil)

	err := u.BuyItem(testuser, testitem)

	assert.NoError(t, err)
}
