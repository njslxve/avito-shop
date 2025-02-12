package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/njslxve/avito-shop/internal/model"
	"github.com/njslxve/avito-shop/internal/usecase"
)

func Info(logger *slog.Logger, ucase *usecase.Usecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		const op = "handler.Info"

		userInfo, err := ucase.Info()
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
