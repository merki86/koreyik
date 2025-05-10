package model

import (
	"time"

	pqdriver "github.com/lib/pq"
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

// func CreateAnime(storage *gorm.DB, ctx context.Context, anime Anime) error {
// 	return storage.WithContext(ctx).Create(&anime).Error
// }

// func GetAnimeById(storage *gorm.DB, ctx context.Context, id int) (Anime, error) {
// 	var anime Anime

// 	if err := storage.WithContext(ctx).First(&anime, id).Error; err != nil {
// 		return Anime{}, err
// 	}
// 	return anime, nil
// }

// func GetAnimeCount(storage *gorm.DB, ctx context.Context) (int64, error) {
// 	var count int64

// 	if err := storage.WithContext(ctx).Model(&Anime{}).Count(&count).Error; err != nil {
// 		return 0, err
// 	}
// 	return count, nil
// }
