package usecase

import (
	"testing"

	"github.com/njslxve/avito-shop/internal/mocks"
	"github.com/njslxve/avito-shop/internal/model"
	"github.com/njslxve/avito-shop/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUser(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)

	name := "testname"
	pass := "testpass"

	mockRepo.On("FindUser", mock.Anything).Return(model.User{Username: name, Password: pass}, nil)

	u := New(nil, nil, &repository.Repository{User: mockRepo})

	user, err := u.User(name, pass)

	assert.NoError(t, err)
	assert.Equal(t, name, user.Username)
}
