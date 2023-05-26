package gorm

import (
	"context"

	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/db/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type gormRecordRepository struct {
	db gorm.DB
}

func (r *gormRecordRepository) Save(ctx context.Context, u *entities.Record) error {
	record := models.Record{}

	r.db.Create(&record)

	return nil
}

func (r *gormRecordRepository) Update(ctx context.Context, u *entities.Record) error {
	var record models.Record

	r.db.Save(&record)

	return nil
}

func (r *gormRecordRepository) Find(ctx context.Context, id uuid.UUID) (*entities.Record, error) {
	var record models.Record

	r.db.First(&record, "id = ?", id)

	return nil, nil
}

func (r *gormRecordRepository) FindAll(ctx context.Context) ([]*entities.Record, error) {
	var records []models.Record

	r.db.Find(&records)

	return nil, nil
}

func (r *gormRecordRepository) Delete(ctx context.Context, id uuid.UUID) error {
	r.db.Delete(&models.Record{}, "10")

	return nil
}

func NewGormRecordRepository() repositories.RecordRepository {
	return &gormRecordRepository{}
}
