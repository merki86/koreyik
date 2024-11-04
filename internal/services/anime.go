package services

import (
	"context"

	"github.com/merki86/koreyik/internal/models"
	"gorm.io/gorm"
)

func GetAnimeById(id int, ctx context.Context, stg *gorm.DB) (models.Anime, error) {
	return models.GetAnimeById(stg, ctx, id)
}

func CreateAnime(anime models.Anime, ctx context.Context, stg *gorm.DB) error {
	return models.CreateAnime(stg, ctx, anime)
}
