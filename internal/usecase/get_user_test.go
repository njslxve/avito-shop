package usecase

import (
	"crypto/sha256"
	"fmt"
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

	hash := sha256.New()
	hash.Write([]byte(pass))

	passwd := fmt.Sprintf("%x", hash.Sum(nil))

	mockRepo.On("FindUserByName", mock.Anything).Return(model.User{Username: name, PasswordHash: passwd}, nil)

	u := New(nil, nil, &repository.Repository{User: mockRepo})

	user, err := u.UserByName(name, pass)

	assert.NoError(t, err)
	assert.Equal(t, name, user.Username)
}
