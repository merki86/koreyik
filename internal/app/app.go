package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	middlewareLogger "github.com/merki86/koreyik/api/middleware/logger"
	"github.com/merki86/koreyik/api/routes"
	"github.com/merki86/koreyik/internal/config"
	"github.com/merki86/koreyik/internal/models"
	"github.com/merki86/koreyik/internal/server"
	"github.com/merki86/koreyik/internal/storage/pq"
	"gitlab.com/greyxor/slogor"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
)

func Run() {

	// Counting server starting time
	startTime := time.Now()

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load .env file: %s", err.Error())
		os.Exit(1)
	}

	cfg := config.New()

	log := setupLogger(cfg.Env)

	log.Info(
		"Starting Köreyik!",
		slog.String("env", cfg.Env),
		slog.String("version", cfg.Version),
	)

	// Loading database (PostgreSQL)
	stg, err := pq.New(cfg.Storage)
	if err != nil {
		log.Error(
			"Failed to connect to the storage",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	} else {
		log.Info(
			"Storage connected",
			slog.String("server", cfg.Storage.Server),
			slog.Int("port", cfg.Storage.Port),
		)

		log.Debug(
			"Storage info",
			slog.String("server", cfg.Storage.Server),
			slog.Int("port", cfg.Storage.Port),
			slog.String("database", cfg.Storage.Database),
			slog.String("username", cfg.Storage.Username),
		)
	}

	stg.AutoMigrate(&models.Anime{})

	// Router
	r := chi.NewRouter()

	r.Use(
		middlewareLogger.New(log),
		middleware.RequestID,
		middleware.Recoverer,
	)

	routes.RegisterRoutes(r, stg, log)

	srv := server.New(cfg, r)
	go func() {
		if err := srv.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error(
				"Failed to run the server",
				slog.String("error", err.Error()),
			)
			os.Exit(1)
		}
	}()

	timeTaken := time.Since(startTime)
	log.Info(
		fmt.Sprintf("Server started: http://%s", cfg.Server.Address),
		slog.String("time_taken", timeTaken.String()),
	)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	<-quit
	log.Info("Server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error(
			"Failed to shut down the server",
			slog.String("error", err.Error()),
		)
		return
	}

	log.Info("Server shut down")
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
