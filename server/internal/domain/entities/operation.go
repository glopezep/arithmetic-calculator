package entities

import (
	"errors"

	"github.com/Rhymond/go-money"
	valueobjects "github.com/glopezep/arithmetic-calculator/internal/domain/value_objects"
	"github.com/google/uuid"
)

var (
	ErrOperationTypeIsBlank    = errors.New("the operation type cannot be blank")
	ErrOperationCostIsBlank    = errors.New("the operation cost cannot be blank")
	ErrOperationCostIsNegative = errors.New("the operation cost cannot be negative")
)

type Operation struct {
	ID   uuid.UUID
	Type valueobjects.OperationType
	Cost money.Money
}

func NewOperation(operationType valueobjects.OperationType, cost money.Money) (*Operation, error) {
	if operationType == "" {
		return nil, ErrOperationTypeIsBlank
	}

	if cost.Amount() < 0 {
		return nil, ErrOperationCostIsNegative
	}

	oType, err := valueobjects.NewOperationType(operationType)
	if err != nil {
		return nil, err
	}

	return &Operation{
		ID:   uuid.New(),
		Type: *oType,
		Cost: cost,
	}, nil
}
