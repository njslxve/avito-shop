package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/njslxve/avito-shop/internal/model"
	"github.com/njslxve/avito-shop/internal/usecase"
	"github.com/njslxve/avito-shop/internal/validation"
)

func Auth(logger *slog.Logger, ucase *usecase.Usecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		const op = "handler.auth"

		var req model.AuthRequest

		if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
			e := model.Error{
				Errors: ErrBadRequest,
			}

			logger.Error("failed to decode request body",
				slog.String("operation", op),
				slog.String("error", err.Error()),
			)

			return c.JSON(http.StatusBadRequest, e)
		}

		err := validation.ValdateAuthRequest(req)
		if err != nil {
			e := model.Error{
				Errors: ErrBadRequest,
			}

			logger.Error("failed to validate request",
				slog.String("operation", op),
				slog.String("error", err.Error()),
			)

			return c.JSON(http.StatusBadRequest, e)
		}

		user, err := ucase.User(req.Username, req.Password)
		if err != nil {
			logger.Error("failed to get user",
				slog.String("operation", op),
				slog.String("error", err.Error()),
			)

			e := model.Error{
				Errors: ErrInternal,
			}

			return c.JSON(http.StatusInternalServerError, e)
		}

		token, err := ucase.Token(user)
		if err != nil {
			logger.Error("failed to generate token",
				slog.String("operation", op),
				slog.String("error", err.Error()),
			)

			e := model.Error{
				Errors: ErrInternal,
			}

			return c.JSON(http.StatusInternalServerError, e)
		}

		res := model.AuthResponse{
			Token: token,
		}

		return c.JSON(http.StatusOK, res)
	}
}
