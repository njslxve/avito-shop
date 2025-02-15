package auth

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
	muckRepo := new(mocks.MockUserRepository)

	name := "testname"
	password := "testpass"

	hash := sha256.New()
	hash.Write([]byte(password))

	testHash := fmt.Sprintf("%x", hash.Sum(nil))

	muckRepo.On("FindUserByName", mock.Anything).Return(model.User{
		Username:     name,
		PasswordHash: testHash,
	}, nil)

	u := New(nil, nil, &repository.Repository{User: muckRepo})

	user, err := u.User(name, password)

	assert.NoError(t, err)
	assert.Equal(t, name, user.Username)
	assert.Equal(t, testHash, user.PasswordHash)
}
