package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/serwennn/koreyik/internal/storage/pq"
	"github.com/serwennn/koreyik/internal/storage/red"
	"log/slog"
)

func RegisterRoutes(r *chi.Mux, stg *pq.Storage, cacheServer *red.CacheServer, log *slog.Logger) {
	// Register the API routes
	api := chi.NewRouter()

	registerAnime(api, stg, cacheServer, log)

	r.Mount("/api", api)
}
