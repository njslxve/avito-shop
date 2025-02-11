package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/njslxve/avito-shop/internal/model"
	"github.com/njslxve/avito-shop/internal/usecase"
)

func BuyItem(logger *slog.Logger, ucase *usecase.Usecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		const op = "handler.BuyItem"

		var item = c.Param("item")

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

		err := ucase.BuyItem(item)
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
