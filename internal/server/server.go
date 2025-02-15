package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/njslxve/avito-shop/internal/config"
	"github.com/njslxve/avito-shop/internal/server/handler"
	"github.com/njslxve/avito-shop/internal/service"
	"github.com/njslxve/avito-shop/internal/service/auth"
)

type Server struct {
	cfg     *config.Config
	logger  *slog.Logger
	service *service.Service
}

func New(cfg *config.Config, logger *slog.Logger, service *service.Service) *Server {
	return &Server{
		cfg:     cfg,
		logger:  logger,
		service: service,
	}
}

func (s *Server) Run() {
	e := echo.New()

	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	}))

	e.POST("/api/auth", handler.Auth(s.logger, s.service.Auth))

	g := e.Group("/api", auth.JWTMiddleware(s.cfg))
	g.GET("/info", handler.Info(s.logger, s.service.Shop))
	g.GET("/buy/:item", handler.BuyItem(s.logger, s.service.Shop))
	g.POST("/sendCoin", handler.SendCoin(s.logger, s.service.Shop))

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			s.logger.Error("failed to start server",
				slog.String("error", err.Error()),
			)
		}
	}()

	s.logger.Info("server started")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	<-done

	s.logger.Info("server shutting down")

	if err := e.Shutdown(ctx); err != nil {
		s.logger.Error("failed to shutdown server")
	}

	s.logger.Info("server stopped")
}
