package handler

import (
	"log/slog"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/njslxve/avito-shop/internal/apperror"
	"github.com/njslxve/avito-shop/internal/model"
	"github.com/njslxve/avito-shop/internal/validation"
)

type SendCoinInterface interface {
	User(userID string) (model.User, error)
	SendCoin(model.User, string, int64) error
}

func SendCoin(logger *slog.Logger, ucase SendCoinInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		const op = "handler.SendCoin"

		token, ok := c.Get("token").(*jwt.Token)
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid token1") //TODO
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid token2") //TODO
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid token3") //TODO
		}

		sender, err := ucase.User(userID)
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

		var req model.SendCoinRequest

		if err := c.Bind(&req); err != nil {
			e := model.Error{
				Errors: apperror.ErrBadRequest,
			}

			logger.Error("failed to bind request body",
				slog.String("operation", op),
				slog.String("error", err.Error()),
			)

			return c.JSON(http.StatusBadRequest, e)
		}

		err = validation.ValidateSendCoinRequest(req)
		if err != nil {
			e := model.Error{
				Errors: apperror.ErrBadRequest,
			}

			logger.Error("failed to validate request",
				slog.String("operation", op),
				slog.String("error", err.Error()),
			)

			return c.JSON(http.StatusBadRequest, e)
		}

		err = ucase.SendCoin(sender, req.ToUser, req.Amount)
		if err != nil {
			e := model.Error{
				Errors: apperror.ErrInternal,
			}

			logger.Error("failed to send coin",
				slog.String("operation", op),
				slog.String("error", err.Error()),
			)

			return c.JSON(http.StatusInternalServerError, e)
		}

		return c.JSON(http.StatusOK, nil)
	}
}
