package usecase

import (
	"testing"

	"github.com/njslxve/avito-shop/internal/mocks"
	"github.com/njslxve/avito-shop/internal/model"
	"github.com/njslxve/avito-shop/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInfo(t *testing.T) {
	mockTransRepo := new(mocks.MockTransactionRepository)
	mockCoinRepo := new(mocks.MockCoinRepository)

	testuser := model.User{
		Username:     "testuser",
		PasswordHash: "testpass",
		Coins:        100,
	}

	u := New(nil, nil, &repository.Repository{
		Transaction: mockTransRepo,
		Coin:        mockCoinRepo,
	})

	mockTransRepo.On("UserHistory", mock.Anything).Return([]string{"item1", "item2", "item3"}, nil)
	mockCoinRepo.On("SenderHistory", mock.Anything).Return([]model.Transaction{
		{
			Username: "testuser2",
			Amount:   10,
		},
	}, nil)

	mockCoinRepo.On("ReceiverHistory", mock.Anything).Return([]model.Transaction{
		{
			Username: "testuser2",
			Amount:   20,
		},
	}, nil)

	_, err := u.Info(testuser)

	assert.NoError(t, err)

}
