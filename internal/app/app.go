package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	middlewareLogger "github.com/serwennn/koreyik/api/middleware/logger"
	"github.com/serwennn/koreyik/api/routes"
	"github.com/serwennn/koreyik/internal/config"
	"github.com/serwennn/koreyik/internal/server"
	"github.com/serwennn/koreyik/internal/storage/pq"
	"github.com/serwennn/koreyik/internal/storage/red"
	"gitlab.com/greyxor/slogor"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
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
		log.Error("Failed to connect to the storage", "error", err.Error())
		os.Exit(1)
	} else {
		log.Info(
			"Connected to the storage",
			slog.String("server", cfg.Storage.Server),
			slog.Int("port", cfg.Storage.Port),
		)

		log.Debug("Storage info",
			slog.String("server", cfg.Storage.Server),
			slog.Int("port", cfg.Storage.Port),
			slog.String("database", cfg.Storage.Database),
			slog.String("username", cfg.Storage.Username),
		)
	}

	// Loading cache server (Redis)
	cacheClient, err := red.New(cfg.CacheServer)
	if err != nil {
		log.Error("Failed to connect to the cache server", "error", err.Error())
		os.Exit(1)
	} else {
		log.Info(
			"Connected to the cache server",
			slog.String("address", cfg.CacheServer.Address),
			slog.Int("database", cfg.CacheServer.Database),
		)

		log.Debug("Cache server info",
			slog.String("address", cfg.CacheServer.Address),
			slog.Int("database", cfg.CacheServer.Database),
		)
	}

	// Router
	r := chi.NewRouter()

	r.Use(
		middlewareLogger.New(log),
		middleware.RequestID,
		middleware.Recoverer,
	)

	routes.RegisterRoutes(r, stg, cacheClient, log)

	// Load static files
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "./web/static"))
	FileServer(r, "/web/static", filesDir)

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
		fmt.Sprintf("Server is running on http://%s/", cfg.Server.Address),
		slog.String("time_taken", timeTaken.String()),
	)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	<-quit
	log.Info("Server is shutting down")

	// TODO: move timeout to config
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error(
			"Failed to shut down the server",
			slog.String("error", err.Error()),
		)
		return
	}

	// Gracefully shutdown the storage
	stg.Shutdown()
	log.Info("Storage has been shut down")

	// Gracefully shutdown the cache server
	err = cacheClient.Shutdown()
	if err != nil {
		log.Error(
			"Failed to shut down the cache",
			slog.String("error", err.Error()),
		)
	} else {
		log.Info("Cache server has been shutdown")
	}

	log.Info("Server has been shut down")
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
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
