package pq

import (
	"context"
	"fmt"

	"github.com/merki86/koreyik/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

var ctx = context.Background()

func New(storageConfig config.Storage) (*Storage, error) {
	url := databaseUrlCreator(storageConfig)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Storage{DB: db}, nil
}

func databaseUrlCreator(storage config.Storage) string {
	// URL should look like this -> "postgres://username:password@host:port/database_name"
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		storage.Username, storage.Password, storage.Server, storage.Port, storage.Database,
	)
}
