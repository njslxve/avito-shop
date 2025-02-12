package main

import (
	"log/slog"
	"os"

	"github.com/njslxve/avito-shop/internal/auth"
	"github.com/njslxve/avito-shop/internal/client/postgres"
	"github.com/njslxve/avito-shop/internal/config"
	"github.com/njslxve/avito-shop/internal/repository"
	"github.com/njslxve/avito-shop/internal/server"
	"github.com/njslxve/avito-shop/internal/usecase"
	"github.com/njslxve/avito-shop/pkg/logger"
)

func main() {
	lg := logger.New()
	slog.SetDefault(lg)

	cfg, err := config.Load()
	if err != nil {
		lg.Error("failed to load config",
			slog.String("error", err.Error()))

		os.Exit(1)
	}

	client, err := postgres.NewClient(cfg)
	if err != nil {
		lg.Error("failed to connect to database",
			slog.String("error", err.Error()))

		os.Exit(1)
	}

	repo := repository.New(client)

	authService := auth.New(cfg)

	ucase := usecase.New(lg, authService, repo)

	srv := server.New(cfg, lg, ucase)
	srv.Run()
}
