package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/merki86/koreyik/internal/models"
	"github.com/merki86/koreyik/internal/storage/pq"
)

type animeImpl struct{}

func registerAnime(r chi.Router, stg *pq.Storage, log *slog.Logger) {
	impl := &animeImpl{}

	r.Route("/anime", func(r chi.Router) {
		r.Get("/{id}", impl.getAnime(stg, log))
		r.Post("/", impl.postAnime(stg))
	})
}

func (impl *animeImpl) getAnime(stg *pq.Storage, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		// Get an entry from the main storage
		anime, err := models.GetAnime(stg, r.Context(), id)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				http.Error(w, fmt.Sprintf("Anime not found. ID: %d", id), http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		log.Debug("Got an entry from the storage")

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

func (impl *animeImpl) postAnime(stg *pq.Storage) http.HandlerFunc {
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
