package pq

import (
	"fmt"

	"github.com/merki86/koreyik/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(storageConfig config.Storage) (*gorm.DB, error) {
	url := databaseUrlCreator(storageConfig)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func databaseUrlCreator(storage config.Storage) string {
	// URL should look like this -> "postgres://username:password@host:port/database_name"
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		storage.Username, storage.Password, storage.Server, storage.Port, storage.Database,
	)
}
