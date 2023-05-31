package mappers

import (
	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	valueobjects "github.com/glopezep/arithmetic-calculator/internal/domain/value_objects"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/db/models"
)

type OperationMapper struct{}

func (m *OperationMapper) ToEntity(model models.Operation) (*entities.Operation, error) {
	oType, err := valueobjects.NewOperationType(model.Type)
	if err != nil {
		return nil, err
	}

	return &entities.Operation{
		ID:   model.ID,
		Type: *oType,
		Cost: model.Cost,
	}, nil
}

func NewOperationMapper() OperationMapper {
	return OperationMapper{}
}
