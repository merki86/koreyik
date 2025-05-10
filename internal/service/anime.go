package service

import (
	"context"
	"math/rand/v2"

	"github.com/merki86/koreyik/internal/model"
	"github.com/merki86/koreyik/internal/repository"
	"gorm.io/gorm"
)

type AnimeService struct {
	animeRepository *repository.AnimeRepository
}

func (s *AnimeService) GetAnimeById(id int, ctx context.Context, stg *gorm.DB) (model.Anime, error) {
	return s.animeRepository.GetAnimeById(stg, ctx, id)
}

func (s *AnimeService) GetRandomAnimeId(ctx context.Context, stg *gorm.DB) (int, error) {
	max, err := s.animeRepository.GetAnimeCount(stg, ctx)
	if err != nil {
		return 0, err
	}
	id := rand.IntN(int(max))

	return id, nil
}

func (s *AnimeService) CreateAnime(anime model.Anime, ctx context.Context, stg *gorm.DB) error {
	return s.animeRepository.CreateAnime(stg, ctx, anime)
}
