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

	r.db.First(&record, "id = ?", id)

	return nil, nil
}

func (r *gormRecordRepository) FindAll(
	ctx context.Context,
	pageNumber, pageSize int,
	sortBy, orderBy string,
) ([]*entities.Record, error) {
	var records []models.Record
	var result []*entities.Record

	r.db.
		Scopes(db.Order(sortBy, orderBy)).
		Scopes(db.Paginate(pageNumber, pageSize)).
		Find(&records)

	for _, v := range records {
		e, err := r.mapper.ToEntity(v)
		if err != nil {
			return nil, err
		}

		result = append(result, e)
	}

	return result, nil
}

func (r *gormRecordRepository) Delete(ctx context.Context, id uuid.UUID) error {
	r.db.Delete(&models.Record{}, "10")

	return nil
}

func NewGormRecordRepository(db *gorm.DB, mapper mappers.RecordMapper) repositories.RecordRepository {
	return &gormRecordRepository{
		db,
		mapper,
	}
}
