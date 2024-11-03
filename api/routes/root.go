package routes

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/merki86/koreyik/internal/storage/pq"
)

func RegisterRoutes(r *chi.Mux, stg *pq.Storage, log *slog.Logger) {
	// Register the API routes
	api := chi.NewRouter()

	registerAnime(api, stg, log)

	r.Mount("/api", api)
}
