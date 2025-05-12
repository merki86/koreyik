package service

import (
	"context"
	"fmt"
	"math/rand/v2"

	"github.com/merki86/koreyik/internal/model"
	"github.com/merki86/koreyik/internal/repository"
)

type AnimeService struct {
	AnimeRepository *repository.AnimeRepository
}

func (s *AnimeService) GetAnimes(ctx context.Context) ([]model.Anime, error) {
	animes, err := s.AnimeRepository.GetAnimes(ctx)
	if err != nil {
		return nil, err
	}
	return animes, nil
}

func (s *AnimeService) GetAnimeById(id int, ctx context.Context) (*model.Anime, error) {
	anime, err := s.AnimeRepository.GetAnimeById(ctx, id)
	if err != nil {
		return nil, err
	}
	return anime, nil
}

func (s *AnimeService) GetRandomAnimeId(ctx context.Context) (int, error) {
	max, err := s.AnimeRepository.GetAnimeCount(ctx)
	if err != nil {
		return 0, err
	}

	// Get random number between 1 and max inclusive; [0; max) + 1 = [1; max]
	if max <= 0 {
		return 0, fmt.Errorf("invalid max value for generating random value (max must be greater than zero): %d", max)
	}
	id := rand.IntN(int(max)) + 1

	return id, nil
}

func (s *AnimeService) CreateAnime(anime model.Anime, ctx context.Context) error {
	return s.AnimeRepository.CreateAnime(ctx, anime)
}

func (s *AnimeService) UpdateAnime(anime model.Anime, ctx context.Context) error {
	return s.AnimeRepository.UpdateAnime(ctx, anime)
}

func (s *AnimeService) DeleteAnime(id int, ctx context.Context) error {
	return s.AnimeRepository.DeleteAnime(ctx, id)
}
