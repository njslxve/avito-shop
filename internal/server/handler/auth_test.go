package handler_test

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/njslxve/avito-shop/internal/mocks"
	"github.com/njslxve/avito-shop/internal/model"
	"github.com/njslxve/avito-shop/internal/server/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuth(t *testing.T) {
	logger := slog.Default()

	e := echo.New()

	mockucase := new(mocks.MockAuthService)

	h := handler.Auth(logger, mockucase)

	req := model.AuthRequest{
		Username: "username",
		Password: "password",
	}

	mockucase.On("User", mock.Anything, mock.Anything).Return(model.User{}, nil)
	mockucase.On("Token", mock.Anything).Return("token", nil)

	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest("POST", "/", bytes.NewBuffer(reqBody))

	rr := httptest.NewRecorder()

	c := e.NewContext(httpReq, rr)
	c.SetPath("/api/auth")

	err := h(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"token":"token"}`, rr.Body.String())

	mockucase.AssertExpectations(t)
}
