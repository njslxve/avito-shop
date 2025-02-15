package handler

import (
	"log/slog"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/njslxve/avito-shop/internal/apperror"
	"github.com/njslxve/avito-shop/internal/model"
)

type InfoInterface interface {
	User(string) (model.User, error)
	Info(model.User) (model.InfoResponse, error)
}

func Info(logger *slog.Logger, ucase InfoInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		const op = "handler.Info"

		token, ok := c.Get("token").(*jwt.Token)
		if !ok {
			logger.Error("invalid token",
				slog.String("operation", op),
			)

			return echo.NewHTTPError(http.StatusBadRequest, apperror.ErrBadRequestToken)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			logger.Error("invalid token",
				slog.String("operation", op),
			)

			return echo.NewHTTPError(http.StatusBadRequest, apperror.ErrBadRequestToken)
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			logger.Error("invalid token",
				slog.String("operation", op),
			)

			return echo.NewHTTPError(http.StatusBadRequest, apperror.ErrBadRequestToken)
		}

		user, err := ucase.User(userID)
		if err != nil {
			e := model.Error{
				Errors: apperror.ErrInternal,
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
				Errors: apperror.ErrInternal,
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
