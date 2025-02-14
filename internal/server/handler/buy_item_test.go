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

func TestBuyItem(t *testing.T) {
	logger := slog.Default()

	cfg := &config.Config{
		JWTSecret: "secret",
	}

	a := auth.New(cfg)

	e := echo.New()
	e.Use(auth.JWTMiddleware(cfg))

	mockucase := new(mocks.MockUsecase)

	e.GET("api/buy/:item", handler.BuyItem(logger, mockucase))

	testUser := model.User{
		ID: "testID-0000-test-test",
	}

	testToken, _ := a.GenerateToken(testUser.ID)

	mockucase.On("ValidateItem", mock.Anything).Return(true)
	mockucase.On("BuyItem", mock.Anything, mock.Anything).Return(nil)
	mockucase.On("UserByID", mock.Anything).Return(model.User{}, nil)

	httpReq := httptest.NewRequest("GET", "/api/buy/powerbank", nil)
	httpReq.Header.Set(echo.HeaderAuthorization, "Bearer "+testToken)

	rr := httptest.NewRecorder()

	e.ServeHTTP(rr, httpReq)

	assert.Equal(t, http.StatusOK, rr.Code)

	mockucase.AssertExpectations(t)
}
