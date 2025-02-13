package usecase

import (
	"testing"

	"github.com/njslxve/avito-shop/internal/mocks"
	"github.com/njslxve/avito-shop/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)

	name := "testname"
	pass := "testpass"

	mockRepo.On("Create", mock.Anything).Return(nil)

	u := New(nil, nil, &repository.Repository{User: mockRepo})

	user, err := u.createUser(name, pass)

	assert.NoError(t, err)
	assert.Equal(t, name, user.Username)
}
