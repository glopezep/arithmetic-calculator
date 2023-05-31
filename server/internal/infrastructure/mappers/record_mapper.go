package mappers

import (
	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/db/models"
)

type RecordMapper struct{}

func (m *RecordMapper) ToEntity(model models.Record) (*entities.Record, error) {

	return &entities.Record{
		ID:                model.ID,
		OperationID:       model.OperationID,
		UserID:            model.UserID,
		Amount:            model.Amount,
		UserBalance:       model.UserBalance,
		OperationResponse: model.OperationResponse,
		CreatedAt:         model.CreatedAt,
	}, nil
}

func NewRecordMapper() RecordMapper {
	return RecordMapper{}
}
