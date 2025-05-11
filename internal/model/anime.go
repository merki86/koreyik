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
}
