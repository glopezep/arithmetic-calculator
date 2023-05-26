package db

import (
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct{}

func (d *Database) Open(dsn string) (*gorm.DB, error) {

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Operation{},
		&models.Record{},
	)
}

func NewDatabase() Database {
	return Database{}
}
