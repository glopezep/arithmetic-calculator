package mappers

import (
	"github.com/Rhymond/go-money"
	"github.com/glopezep/arithmetic-calculator/internal/common"
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
		AggregateBase: common.AggregateBase{
			ID: model.ID,
		},
		Type: *oType,
		Cost: *money.New(model.Cost, "USD"),
	}, nil
}

func NewOperationMapper() OperationMapper {
	return OperationMapper{}
}
