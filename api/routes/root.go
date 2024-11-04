package routes

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(r *chi.Mux, stg *gorm.DB, log *slog.Logger) {
	// Register the API routes
	api := chi.NewRouter()

	registerAnime(api, stg, log)

	r.Mount("/api", api)
}
