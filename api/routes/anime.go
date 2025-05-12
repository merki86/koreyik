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

type animeHandler struct {
	animeService *service.AnimeService
}

func registerAnime(r chi.Router, stg *gorm.DB, log *slog.Logger) {
	handler := &animeHandler{
		animeService: &service.AnimeService{
			AnimeRepository: &repository.AnimeRepository{
				DB: stg,
			},
		},
	}

	r.Route("/anime", func(r chi.Router) {
		r.Get("/{id}", handler.getAnime(log))
		r.Get("/random", handler.getRandomAnime(log))
		r.Post("/", handler.postAnime(log))
		r.Put("/{id}", handler.updateAnime(log))
	})
}

func (h *animeHandler) getAnime(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get ID from URL and convert to int
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Debug("Convert ID: " + err.Error())
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		// Get anime by ID
		anime, err := h.animeService.GetAnimeById(id, r.Context())
		if err != nil {
			log.Debug("Get anime by ID: " + err.Error())
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, fmt.Sprintf("Anime not found. ID: %d", id), http.StatusNotFound)
			} else {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			return
		}

		// Struct to JSON
		serialized, err := json.Marshal(anime)
		if err != nil {
			log.Debug("Marshal JSON: " + err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Write response
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(serialized)
		if err != nil {
			log.Debug("Write http response: " + err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}
}

func (h *animeHandler) getRandomAnime(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := h.animeService.GetRandomAnimeId(r.Context())
		if err != nil {
			log.Debug("Get random anime: " + err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("%d", id), http.StatusSeeOther)
	}
}

func (h *animeHandler) postAnime(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var anime model.Anime

		// Read body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Debug("Read body: " + err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// JSON to struct
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

		// Create anime
		err = h.animeService.CreateAnime(anime, r.Context())
		if err != nil {
			log.Debug("Create Anime: " + err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		http.Redirect(w, r, fmt.Sprintf("/anime/%s", strconv.Itoa(anime.ID)), http.StatusFound)
	}
}

func (h *animeHandler) updateAnime(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get ID from URL and convert to int
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Debug("Convert ID: " + err.Error())
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var anime model.Anime

		// Read body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Debug("Read body: " + err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// JSON to struct
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

		// Update anime
		anime.ID = id
		err = h.animeService.UpdateAnime(anime, r.Context())
		if err != nil {
			log.Debug("Update Anime: " + err.Error())
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, fmt.Sprintf("Anime not found. ID: %d", id), http.StatusNotFound)
			} else {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			return
		}

		// Write response
		w.WriteHeader(http.StatusOK)
		http.Redirect(w, r, fmt.Sprintf("/anime/%s", strconv.Itoa(anime.ID)), http.StatusFound)
	}
}
