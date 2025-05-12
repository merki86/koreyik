package repository

import (
	"context"

	"github.com/merki86/koreyik/internal/model"
	"gorm.io/gorm"
)

type AnimeRepository struct {
	DB *gorm.DB
}

func (r *AnimeRepository) GetAnimes(ctx context.Context) ([]model.Anime, error) {
	var animes []model.Anime

	if err := r.DB.WithContext(ctx).Find(&animes).Error; err != nil {
		return nil, err
	}
	return animes, nil
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

func (r *AnimeRepository) CreateAnime(ctx context.Context, anime model.Anime) error {
	return r.DB.WithContext(ctx).Create(&anime).Error
}

func (r *AnimeRepository) UpdateAnime(ctx context.Context, anime model.Anime) error {
	return r.DB.WithContext(ctx).Save(&anime).Error
}

func (r *AnimeRepository) DeleteAnime(ctx context.Context, id int) error {
	return r.DB.WithContext(ctx).Delete(&model.Anime{}, id).Error
}
