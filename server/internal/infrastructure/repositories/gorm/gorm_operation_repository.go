package gorm

import (
	"context"

	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/db/models"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/mappers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type gormOperationRepository struct {
	db     *gorm.DB
	mapper mappers.OperationMapper
}

func (r *gormOperationRepository) Save(ctx context.Context, u *entities.Operation) error {
	operation := models.Operation{}

	r.db.Create(&operation)

	return nil
}

func (r *gormOperationRepository) Update(ctx context.Context, u *entities.Operation) error {
	var operation models.Operation

	r.db.Save(&operation)

	return nil
}

func (r *gormOperationRepository) Find(ctx context.Context, id uuid.UUID) (*entities.Operation, error) {
	var operation models.Operation

	r.db.First(&operation, "id = ?", id)

	return r.mapper.ToEntity(operation)

}

func (r *gormOperationRepository) FindAll(ctx context.Context) ([]*entities.Operation, error) {
	var operations []models.Operation
	var result []*entities.Operation

	r.db.Find(&operations)

	for _, v := range operations {
		e, err := r.mapper.ToEntity(v)
		if err != nil {
			return nil, err
		}

		result = append(result, e)
	}

	return result, nil
}

func (r *gormOperationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	r.db.Delete(&models.Operation{}, "10")

	return nil
}

func NewGormOperationRepository(db *gorm.DB, mapper mappers.OperationMapper) repositories.OperationRepository {
	return &gormOperationRepository{
		db,
		mapper,
	}
}
