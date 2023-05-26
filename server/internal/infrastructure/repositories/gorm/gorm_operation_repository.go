package gorm

import (
	"context"

	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/db/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type gormOperationRepository struct {
	db gorm.DB
}

func (r *gormOperationRepository) Save(ctx context.Context, u *entities.Operation) error {
	record := models.Operation{}

	r.db.Create(&record)

	return nil
}

func (r *gormOperationRepository) Update(ctx context.Context, u *entities.Operation) error {
	var record models.Operation

	r.db.Save(&record)

	return nil
}

func (r *gormOperationRepository) Find(ctx context.Context, id uuid.UUID) (*entities.Operation, error) {
	var record models.Operation

	r.db.First(&record, "id = ?", id)

	return nil, nil
}

func (r *gormOperationRepository) FindAll(ctx context.Context) ([]*entities.Operation, error) {
	var records []models.Operation

	r.db.Find(&records)

	return nil, nil
}

func (r *gormOperationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	r.db.Delete(&models.Operation{}, "10")

	return nil
}

func NewGormOperationRepository() repositories.OperationRepository {
	return &gormOperationRepository{}
}
