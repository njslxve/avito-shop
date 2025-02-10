package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/njslxve/avito-shop/internal/server/handler"
	"github.com/njslxve/avito-shop/internal/server/mw"
)

type Server struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *Server {
	return &Server{
		logger: logger,
	}
}

func (s *Server) Run() {
	r := chi.NewRouter()

	r.Use(middleware.Timeout(30 * time.Second))

	r.Route("/api", func(r chi.Router) {
		r.Post("/auth", handler.Auth)

		r.Group(func(r chi.Router) {
			r.Use(mw.AuthMiddleware)

			r.Get("/info", handler.Info)
		})
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	s.logger.Info("Starting server",
		slog.String("address", ":8080"),
	)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Debug("server error",
				slog.String("error", err.Error()),
			)
		}
	}()

	s.logger.Info("server started")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	<-done

	s.logger.Info("server shutting down")

	if err := srv.Shutdown(ctx); err != nil {
		s.logger.Error("failed to shutdown server")
	}

	s.logger.Info("server stopped")
}
