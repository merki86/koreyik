package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/merki86/koreyik/internal/model"
	"github.com/merki86/koreyik/internal/repository"
	"github.com/merki86/koreyik/internal/service"
	"gorm.io/gorm"
)

type animeImpl struct {
	animeService *service.AnimeService
}

func registerAnime(r chi.Router, stg *gorm.DB, log *slog.Logger) {
	impl := &animeImpl{
		animeService: &service.AnimeService{
			AnimeRepository: &repository.AnimeRepository{
				DB: stg,
			},
		},
	}

	r.Route("/anime", func(r chi.Router) {
		r.Get("/{id}", impl.getAnime(stg, log))
		r.Get("/random", impl.getRandomAnime(stg, log))
		r.Post("/", impl.postAnime(stg, log))
	})
}

func (impl *animeImpl) getAnime(stg *gorm.DB, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Debug("Convert ID: " + err.Error())
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		anime, err := impl.animeService.GetAnimeById(id, r.Context(), stg)
		if err != nil {
			log.Debug("Get anime by ID: " + err.Error())
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, fmt.Sprintf("Anime not found. ID: %d", id), http.StatusNotFound)
			} else {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			return
		}

		serialized, err := json.Marshal(anime)
		if err != nil {
			log.Debug("Marshal JSON: " + err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(serialized)
		if err != nil {
			log.Debug("Write http response: " + err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}
}

func (impl *animeImpl) getRandomAnime(stg *gorm.DB, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := impl.animeService.GetRandomAnimeId(r.Context(), stg)
		if err != nil {
			log.Debug("Get random anime: " + err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("%d", id), http.StatusSeeOther)
	}
}

func (impl *animeImpl) postAnime(stg *gorm.DB, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var anime model.Anime

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Debug("Read body: " + err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(body, &anime)
		if err != nil {
			log.Debug("Unmarshal JSON: " + err.Error())
			if isValid := json.Valid(body); !isValid {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
			} else {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			return
		}

		err = impl.animeService.CreateAnime(anime, r.Context(), stg)
		if err != nil {
			log.Debug("Create Anime: " + err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		http.Redirect(w, r, fmt.Sprintf("/anime/%s", strconv.Itoa(anime.ID)), http.StatusFound)
	}
}
