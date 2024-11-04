package services

import (
	"context"

	"github.com/merki86/koreyik/internal/models"
	"gorm.io/gorm"
)

func GetAnimeById(id int, ctx context.Context, stg *gorm.DB) (models.Anime, error) {
	return models.GetAnimeById(stg, ctx, id)
}
