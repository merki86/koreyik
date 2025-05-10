package repository

import (
	"context"

	"github.com/merki86/koreyik/internal/model"
	"gorm.io/gorm"
)

type AnimeRepository struct {
	DB *gorm.DB
}

func (r *AnimeRepository) CreateAnime(ctx context.Context, anime model.Anime) error {
	return r.DB.WithContext(ctx).Create(&anime).Error
}

func (r *AnimeRepository) GetAnimeById(ctx context.Context, id int) (*model.Anime, error) {
	var anime model.Anime

	if err := r.DB.WithContext(ctx).First(&anime, id).Error; err != nil {
		return nil, err
	}
	return &anime, nil
}

func (r *AnimeRepository) GetAnimeCount(ctx context.Context) (int64, error) {
	var count int64

	if err := r.DB.WithContext(ctx).Model(&model.Anime{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
