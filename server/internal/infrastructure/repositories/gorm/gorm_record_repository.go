package gorm

import (
	"context"

	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/db"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/db/models"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/mappers"
	"github.com/glopezep/arithmetic-calculator/internal/interfaces/lambda/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type gormRecordRepository struct {
	db     *gorm.DB
	mapper mappers.RecordMapper
}

func (r *gormRecordRepository) Save(ctx context.Context, re *entities.Record) error {
	record := models.Record{
		ID:                re.ID,
		OperationID:       re.OperationID,
		UserID:            re.UserID,
		Amount:            re.Amount,
		UserBalance:       re.UserBalance,
		OperationResponse: re.OperationResponse,
	}

	r.db.Create(&record)

	return nil
}

func (r *gormRecordRepository) Update(ctx context.Context, re *entities.Record) error {
	var record models.Record

	r.db.Save(&record)

	return nil
}

func (r *gormRecordRepository) Find(ctx context.Context, id uuid.UUID) (*entities.Record, error) {
	var record models.Record

	err := r.db.First(&record, "id = ?", id).Error

	return r.mapper.ToEntity(record), err
}

func (r *gormRecordRepository) FindAll(
	ctx context.Context,
	pageNumber, pageSize int,
	sortBy, orderBy string,
) (*repositories.PaginatedResult[entities.Record], error) {
	helperContext := ctx.Value(helpers.ContextKey("context")).(helpers.Context)

	var records []models.Record
	var result []*entities.Record
	var count int64

	r.db.
		Model(&models.Record{}).
		Count(&count)

	r.db.
		Scopes(db.Order(sortBy, orderBy)).
		Scopes(db.Paginate(pageNumber, pageSize)).
		Where("user_id = ?", helperContext.UserID).
		Find(&records)

	for _, v := range records {
		result = append(result, r.mapper.ToEntity(v))
	}

	return &repositories.PaginatedResult[entities.Record]{
		Items:       result,
		TotalCount:  count,
		Offset:      int64(pageNumber),
		Limit:       int64(pageSize),
		HasNextPage: pageNumber*pageSize < int(count),
	}, nil
}

func (r *gormRecordRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.Delete(&models.Record{}, id).Error
}

func NewGormRecordRepository(db *gorm.DB, mapper mappers.RecordMapper) repositories.RecordRepository {
	return &gormRecordRepository{
		db,
		mapper,
	}
}
