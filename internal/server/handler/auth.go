package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/njslxve/avito-shop/internal/model"
	"github.com/njslxve/avito-shop/internal/usecase"
	"github.com/njslxve/avito-shop/internal/validation"
)

func Auth(logger *slog.Logger, ucase *usecase.Usecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.auth"

		var req model.AuthRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			e := model.Error{
				Errors: ErrBadRequest,
			}

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			if err := json.NewEncoder(w).Encode(e); err != nil {
				logger.Error("failed to encode error response",
					slog.String("operation", op),
					slog.String("error", err.Error()),
				)
			}

			return
		}

		err := validation.ValdateAuthRequest(req)
		if err != nil {
			e := model.Error{
				Errors: ErrBadRequest,
			}

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			if err := json.NewEncoder(w).Encode(e); err != nil {
				logger.Error("failed to encode error response",
					slog.String("operation", op),
					slog.String("error", err.Error()),
				)
			}

			return
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

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)

			if err := json.NewEncoder(w).Encode(e); err != nil {
				logger.Error("failed to encode error response",
					slog.String("operation", op),
					slog.String("error", err.Error()),
				)
			}

			return
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

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)

			if err := json.NewEncoder(w).Encode(e); err != nil {
				logger.Error("failed to encode error response",
					slog.String("operation", op),
					slog.String("error", err.Error()),
				)
			}

			return
		}

		res := model.AuthResponse{
			Token: token,
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(res); err != nil {
			logger.Error("failed to encode response",
				slog.String("operation", op),
				slog.String("error", err.Error()),
			)

			e := model.Error{
				Errors: ErrInternal,
			}

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)

			if err := json.NewEncoder(w).Encode(e); err != nil {
				logger.Error("failed to encode error response",
					slog.String("operation", op),
					slog.String("error", err.Error()),
				)
			}

			return
		}
	}
}
