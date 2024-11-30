package services

import (
	"context"
	"math/rand/v2"

	"github.com/merki86/koreyik/internal/models"
	"gorm.io/gorm"
)

func GetAnimeById(id int, ctx context.Context, stg *gorm.DB) (models.Anime, error) {
	return models.GetAnimeById(stg, ctx, id)
}

func GetRandomAnimeId(ctx context.Context, stg *gorm.DB) (int, error) {
	max, err := models.GetAnimeCount(stg, ctx)
	if err != nil {
		return 0, err
	}
	id := rand.IntN(int(max))

	return id, nil
}

func CreateAnime(anime models.Anime, ctx context.Context, stg *gorm.DB) error {
	return models.CreateAnime(stg, ctx, anime)
}
