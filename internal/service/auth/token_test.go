package auth

import (
	"log/slog"
	"testing"

	"github.com/njslxve/avito-shop/internal/config"
	"github.com/njslxve/avito-shop/internal/mocks"
	"github.com/njslxve/avito-shop/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	cfg := &config.Config{
		JWTSecret: "secret",
	}

	logger := slog.Default()
	mockRepo := new(mocks.MockUserRepository)
	testUserID := "testID-0000-test-test"

	u := New(cfg, logger, &repository.Repository{
		User: mockRepo,
	})

	_, err := u.Token(testUserID)

	assert.NoError(t, err)
}
