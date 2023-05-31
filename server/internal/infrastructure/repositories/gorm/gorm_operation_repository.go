package gorm

import (
	"context"

	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/db"
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

func (r *gormOperationRepository) FindAll(ctx context.Context,
	pageNumber, pageSize int,
	sortBy, orderBy string,
) (*repositories.PaginatedResult[entities.Operation], error) {
	var operations []models.Operation
	var result []*entities.Operation
	var count int64

	r.db.
		Model(&models.Operation{}).
		Count(&count)

	r.db.
		Scopes(db.Order(sortBy, orderBy)).
		Scopes(db.Paginate(pageNumber, pageSize)).
		Find(&operations)

	for _, v := range operations {
		e, err := r.mapper.ToEntity(v)
		if err != nil {
			return nil, err
		}

		result = append(result, e)
	}

	return &repositories.PaginatedResult[entities.Operation]{
		Items:       result,
		TotalCount:  count,
		Offset:      int64(pageNumber),
		Limit:       int64(pageSize),
		HasNextPage: pageNumber*pageSize < int(count),
	}, nil
}

func (r *gormOperationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}

func NewGormOperationRepository(db *gorm.DB, mapper mappers.OperationMapper) repositories.OperationRepository {
	return &gormOperationRepository{
		db,
		mapper,
	}
}
