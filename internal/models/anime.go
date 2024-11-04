package models

import (
	"context"
	"time"

	pqdriver "github.com/lib/pq"
	"gorm.io/gorm"
)

type Anime struct {
	ID           int
	ThumbnailURL string
	Description  string
	Rating       string

	TitleKk string
	TitleJp string
	TitleEn string

	Status         string
	StartedAiring  time.Time
	FinishedAiring time.Time

	Genres pqdriver.StringArray `gorm:"type:text[]"`
	Themes pqdriver.StringArray `gorm:"type:text[]"`

	Seasons  int
	Episodes int
	Duration int

	Studios   pqdriver.StringArray `gorm:"type:text[]"`
	Producers pqdriver.StringArray `gorm:"type:text[]"`

	//Related []MediaEntry
}

func CreateAnime(storage *gorm.DB, ctx context.Context, anime Anime) error {
	return storage.WithContext(ctx).Create(&anime).Error
}

func GetAnimeById(storage *gorm.DB, ctx context.Context, id int) (Anime, error) {
	var anime Anime

	if err := storage.WithContext(ctx).First(&anime, id).Error; err != nil {
		return Anime{}, err
	}
	return anime, nil
}
