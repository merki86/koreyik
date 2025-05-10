package service

import (
	"context"
	"math/rand/v2"

	"github.com/merki86/koreyik/internal/model"
	"github.com/merki86/koreyik/internal/repository"
)

type AnimeService struct {
	AnimeRepository *repository.AnimeRepository
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
	id := rand.IntN(int(max))

	return id, nil
}

func (s *AnimeService) CreateAnime(anime model.Anime, ctx context.Context) error {
	return s.AnimeRepository.CreateAnime(ctx, anime)
}
