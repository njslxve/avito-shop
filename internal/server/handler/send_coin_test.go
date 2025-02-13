package handler_test

import (
	"bytes"
	"encoding/json"
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

func TestSendCoin(t *testing.T) {
	logger := slog.Default()

	cfg := &config.Config{
		JWTSecret: "secret",
	}

	a := auth.New(cfg)

	e := echo.New()
	e.Use(auth.JWTMiddleware(cfg))

	mockucase := new(mocks.MockUsecase)

	e.POST("api/sendCoin", handler.SendCoin(logger, mockucase))

	testUser := model.User{
		Username: "testuser",
		Password: "testpass",
		Coins:    1000,
	}

	req := model.SendCoinRequest{
		ToUser: "testreceiver",
		Amount: 200,
	}

	reqBody, _ := json.Marshal(req)

	testToken, _ := a.GenerateToken(testUser.Username, testUser.Password)

	mockucase.On("User", mock.Anything, mock.Anything).Return(testUser, nil)
	mockucase.On("SendCoin", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	httpReq := httptest.NewRequest("POST", "/api/sendCoin", bytes.NewBuffer(reqBody))
	httpReq.Header.Set(echo.HeaderAuthorization, "Bearer "+testToken)
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	t.Logf("AAAAAAAAAA: %s", httpReq.Body)

	rr := httptest.NewRecorder()

	e.ServeHTTP(rr, httpReq)

	assert.Equal(t, http.StatusOK, rr.Code)

	mockucase.AssertExpectations(t)
}
