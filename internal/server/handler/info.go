package handler

import (
	"log/slog"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/njslxve/avito-shop/internal/model"
)

type InfoInterface interface {
	User(string, string) (model.User, error)
	Info(model.User) (model.InfoResponse, error)
}

func Info(logger *slog.Logger, ucase InfoInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		const op = "handler.Info"

		token, ok := c.Get("token").(*jwt.Token)
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid token")
		}

		username, ok := claims["username"].(string)
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid token")
		}

		password, ok := claims["password"].(string)
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid token")
		}

		user, err := ucase.User(username, password)
		if err != nil {
			e := model.Error{
				Errors: ErrInternal,
			}

			logger.Error("failed to get user",
				slog.String("operation", op),
				slog.String("error", err.Error()),
			)

			return c.JSON(http.StatusInternalServerError, e)
		}

		userInfo, err := ucase.Info(user)
		if err != nil {
			e := model.Error{
				Errors: ErrInternal,
			}

			logger.Error("failed to get info",
				slog.String("operation", op),
				slog.String("error", err.Error()),
			)

			return c.JSON(http.StatusInternalServerError, e)
		}

		return c.JSON(http.StatusOK, userInfo)
	}
}
