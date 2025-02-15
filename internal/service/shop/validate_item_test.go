package shop

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/njslxve/avito-shop/internal/mocks"
	"github.com/njslxve/avito-shop/internal/model"
	"github.com/njslxve/avito-shop/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestValidateItemHappy(t *testing.T) {
	mockItemRepo := new(mocks.MockItemRepository)

	itemname := "testitem"

	mockItemRepo.On("FindItem", mock.Anything).Return(model.Item{
		Type:  "testitem",
		Price: 10,
	}, nil)

	u := New(nil, &repository.Repository{
		Item: mockItemRepo,
	})

	valid := u.ValidateItem(itemname)

	assert.True(t, valid)
}

func TestValidateItemBad(t *testing.T) {
	logger := slog.Default()

	mockItemRepo := new(mocks.MockItemRepository)

	itemname := "testitem"

	e := fmt.Errorf("item not found")

	mockItemRepo.On("FindItem", mock.Anything).Return(model.Item{}, e)

	u := New(logger, &repository.Repository{
		Item: mockItemRepo,
	})

	valid := u.ValidateItem(itemname)

	assert.False(t, valid)
}
