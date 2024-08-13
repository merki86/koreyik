package app

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/serwennn/koreyik/internal/config"
	"github.com/serwennn/koreyik/internal/network/routes"
	"github.com/serwennn/koreyik/internal/server"
	"github.com/serwennn/koreyik/internal/storage"
	"gitlab.com/greyxor/slogor"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
)

func Run() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("Failed to load .env file: %s", err.Error()))
	}

	cfg := config.New()

	log := setupLogger(cfg.Env)

	stg, err := storage.New(cfg.StoragePath)
	if err != nil {
		log.Error("Failed to connect to the storage", "error", err.Error())
		os.Exit(1)
	} else {
		log.Info("Connected to the storage")
	}

	_ = stg // TODO: Use the storage

	// Router
	r := chi.NewRouter()

	r.Use(
		middleware.RequestID,
		middleware.Recoverer,
		middleware.Logger,
	)

	routes.RegisterRoutes(r)

	srv := server.NewServer(cfg, r)
	go func() {
		if err := srv.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("Failed to run the server", "error", err.Error())
			os.Exit(1)
		}
	}()

	log.Info(
		fmt.Sprintf("Server is running on http://%s", cfg.Address),
		"env", cfg.Env,
	)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	<-quit
	log.Info("Server is shutting down")

	if err := srv.Shutdown(); err != nil {
		log.Error("Failed to shut down the server", "error", err.Error())

		return
	}

	// TODO: Close the storage

	log.Info("Server has been shut down")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case EnvLocal:
		log = slog.New(
			//slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
			slogor.NewHandler(os.Stdout, slogor.Options{
				TimeFormat: "2006-01-02 15:04:05",
				Level:      slog.LevelDebug,
				ShowSource: false,
			}),
		)
	case EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default: // Default to production
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
