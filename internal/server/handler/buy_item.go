package handler

import (
	"log/slog"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/njslxve/avito-shop/internal/model"
)

type ByuItemInterface interface {
	User(username, password string) (model.User, error)
	ValidateItem(item string) bool
	BuyItem(user model.User, item string) error
}

func BuyItem(logger *slog.Logger, ucase ByuItemInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		const op = "handler.BuyItem"

		var item = c.Param("item")
		token, ok := c.Get("token").(*jwt.Token)
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid token1")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid token2")
		}

		username, ok := claims["username"].(string)
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid token3")
		}

		password, ok := claims["password"].(string)
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid token4")
		}

		user, err := ucase.User(username, password)
		if err != nil {
			e := model.Error{
				Errors: ErrInternal,
			}

			logger.Error("failed to find user",
				slog.String("operation", op),
				slog.String("error", err.Error()),
			)

			return c.JSON(http.StatusInternalServerError, e)
		}

		if !ucase.ValidateItem(item) {
			e := model.Error{
				Errors: ErrBadRequest,
			}

			logger.Error("failed to validate item",
				slog.String("operation", op),
				slog.String("item", c.FormValue("item")),
			)

			return c.JSON(http.StatusBadRequest, e)
		}

		err = ucase.BuyItem(user, item)
		if err != nil {
			e := model.Error{
				Errors: ErrInternal,
			}

			logger.Error("failed to buy item",
				slog.String("operation", op),
				slog.String("item", item),
				slog.String("error", err.Error()),
			)

			return c.JSON(http.StatusInternalServerError, e)
		}

		return c.JSON(http.StatusOK, nil)
	}
}
