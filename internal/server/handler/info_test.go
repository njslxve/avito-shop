package handler_test

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/njslxve/avito-shop/internal/auth"
	"github.com/njslxve/avito-shop/internal/config"
	"github.com/njslxve/avito-shop/internal/mocks"
	"github.com/njslxve/avito-shop/internal/model"
	"github.com/njslxve/avito-shop/internal/server/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInfo(t *testing.T) {
	logger := slog.Default()

	cfg := &config.Config{
		JWTSecret: "secret",
	}

	a := auth.New(cfg)

	e := echo.New()
	e.Use(auth.JWTMiddleware(cfg))

	mockucase := new(mocks.MockUsecase)

	e.GET("api/info", handler.Info(logger, mockucase))

	testUser := model.User{
		ID: "testID-0000-test-test",
	}

	testToken, _ := a.GenerateToken(testUser.ID)

	mockucase.On("UserByID", mock.Anything).Return(testUser, nil)
	mockucase.On("Info", mock.Anything).Return(model.InfoResponse{
		Coins: 1000,
	}, nil)

	httpReq := httptest.NewRequest("GET", "/api/info", nil)
	httpReq.Header.Set(echo.HeaderAuthorization, "Bearer "+testToken)
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rr := httptest.NewRecorder()

	e.ServeHTTP(rr, httpReq)

	assert.Equal(t, http.StatusOK, rr.Code)

	mockucase.AssertExpectations(t)
}
