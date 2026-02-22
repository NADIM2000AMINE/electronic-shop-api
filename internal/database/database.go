package database

import (
	"electronic-shop-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	// Création automatique des tables à partir de nos modèles
	return db.AutoMigrate(
		&models.Shop{},
		&models.User{},
		&models.Product{},
		&models.Transaction{},
	)
}
