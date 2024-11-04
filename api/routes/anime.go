package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/merki86/koreyik/internal/models"
	"github.com/merki86/koreyik/internal/services"
	"gorm.io/gorm"
)

type animeImpl struct{}

func registerAnime(r chi.Router, stg *gorm.DB, log *slog.Logger) {
	impl := &animeImpl{}

	r.Route("/anime", func(r chi.Router) {
		r.Get("/{id}", impl.getAnime(stg, log))
		r.Post("/", impl.postAnime(stg))
	})
}

func (impl *animeImpl) getAnime(stg *gorm.DB, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		anime, err := services.GetAnimeById(id, r.Context(), stg)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, fmt.Sprintf("Anime not found. ID: %d", id), http.StatusNotFound)
			} else {
				log.Debug(err.Error())
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			return
		}

		serialized, err := json.Marshal(anime)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(serialized)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (impl *animeImpl) postAnime(stg *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newAnime models.Anime

		err := json.NewDecoder(r.Body).Decode(&newAnime)
		if err != nil {
			http.Error(w, "Json decode: "+err.Error(), http.StatusBadRequest)
			return
		}

		err = models.CreateAnime(stg, r.Context(), newAnime)
		if err != nil {
			http.Error(w, "Create Anime: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		http.Redirect(w, r, fmt.Sprintf("/anime/%s", strconv.Itoa(newAnime.ID)), http.StatusFound)
	}
}
