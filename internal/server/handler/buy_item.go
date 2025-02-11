package handler

import (
	"log/slog"
	"net/http"

	"github.com/njslxve/avito-shop/internal/usecase"
)

func BuyItem(logger *slog.Logger, ucase *usecase.Usecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// валидируем item
		// либо 400 либо пускаем дальше

		// отдаем в логику и ждем 200
	}
}
