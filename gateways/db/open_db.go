package db

import (
	"fmt"

	"github.com/irenicaa/go-blog-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// OpenDB ...
func OpenDB(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("unable to connect to the DB: %w", err)
	}

	if err := db.AutoMigrate(&models.Post{}); err != nil {
		return nil, fmt.Errorf("unable to automatically migrate the DB: %w", err)
	}

	return db, nil
}
