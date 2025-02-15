package main

import (
	"log/slog"
	"os"

	"github.com/njslxve/avito-shop/internal/client/postgres"
	"github.com/njslxve/avito-shop/internal/config"
	"github.com/njslxve/avito-shop/internal/repository"
	"github.com/njslxve/avito-shop/internal/server"
	"github.com/njslxve/avito-shop/internal/service"
	"github.com/njslxve/avito-shop/internal/service/auth"
	"github.com/njslxve/avito-shop/internal/service/shop"
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

	authService := auth.New(cfg, lg, repo)
	shopService := shop.New(lg, repo)

	ucase := service.New(authService, shopService)

	srv := server.New(cfg, lg, ucase)
	srv.Run()
}
