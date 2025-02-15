package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/njslxve/avito-shop/internal/client/postgres"
	"github.com/njslxve/avito-shop/internal/config"
	"github.com/njslxve/avito-shop/internal/model"
	"github.com/njslxve/avito-shop/internal/repository"
	"github.com/njslxve/avito-shop/internal/server"
	"github.com/njslxve/avito-shop/internal/service"
	"github.com/njslxve/avito-shop/internal/service/auth"
	"github.com/njslxve/avito-shop/internal/service/shop"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	testCongig := &config.Config{
		Port:      "8080",
		JWTSecret: "secret",
		DB: config.Database{
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
			Password: "password",
			Name:     "postgres",
		},
	}

	logger := slog.Default()

	client, err := postgres.NewClient(testCongig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	repo := repository.New(client)

	authService := auth.New(testCongig, logger, repo)
	shopService := shop.New(logger, repo)

	ucase := service.New(authService, shopService)

	srv := server.New(testCongig, logger, ucase)

	go func() {
		srv.Run()
	}()

	time.Sleep(100 * time.Millisecond)

	code := m.Run()

	os.Exit(code)
}

func TestBuyItem(t *testing.T) {
	token, err := token("byu_testuser", "buy_testpass")
	assert.NoError(t, err)

	httpReq, err := http.NewRequest("GET", "http://localhost:8080/api/buy/powerbank", nil)
	httpReq.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	assert.NoError(t, err)

	resp, err := http.DefaultClient.Do(httpReq)
	defer resp.Body.Close()

	assert.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestSendCoin(t *testing.T) {
	firstToken, err := token("sender_testuser", "sender_testpass")
	assert.NoError(t, err)

	_, err = token("receiver_testuser", "receiver_testpass")
	assert.NoError(t, err)

	req := model.SendCoinRequest{
		ToUser: "receiver_testuser",
		Amount: 100,
	}

	reqBody, _ := json.Marshal(req)

	httpReq, err := http.NewRequest("POST", "http://localhost:8080/api/sendCoin", bytes.NewBuffer(reqBody))
	httpReq.Header.Set(echo.HeaderAuthorization, "Bearer "+firstToken)
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	assert.NoError(t, err)

	resp, err := http.DefaultClient.Do(httpReq)
	defer resp.Body.Close()

	assert.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestInfo(t *testing.T) {
	token, err := token("info_testuser", "info_testpass")
	assert.NoError(t, err)

	httpReq, err := http.NewRequest("GET", "http://localhost:8080/api/info", nil)
	httpReq.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	assert.NoError(t, err)

	resp, err := http.DefaultClient.Do(httpReq)
	defer resp.Body.Close()

	assert.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode)
}

func token(username, pass string) (string, error) {
	req := model.AuthRequest{
		Username: username,
		Password: pass,
	}

	reqBody, _ := json.Marshal(req)

	httpReq, err := http.Post("http://localhost:8080/api/auth", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	defer httpReq.Body.Close()

	if httpReq.StatusCode != 200 {
		return "", fmt.Errorf("failed to get token")
	}

	var resBody model.AuthResponse

	err = json.NewDecoder(httpReq.Body).Decode(&resBody)
	if err != nil {
		return "", err
	}

	return resBody.Token, nil
}
