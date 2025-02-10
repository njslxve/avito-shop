package main

import (
	"log/slog"

	"github.com/njslxve/avito-shop/internal/server"
	"github.com/njslxve/avito-shop/pkg/logger"
)

func main() {
	lg := logger.New()
	slog.SetDefault(lg)

	srv := server.New(lg)
	srv.Run()
}
