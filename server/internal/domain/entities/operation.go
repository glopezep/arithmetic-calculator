package entities

import (
	"errors"

	valueobjects "github.com/glopezep/arithmetic-calculator/internal/domain/value_objects"
	"github.com/google/uuid"
)

var (
	ErrOperationTypeIsBlank    = errors.New("the operation type cannot be blank")
	ErrOperationCostIsBlank    = errors.New("the operation cost cannot be blank")
	ErrOperationCostIsNegative = errors.New("the operation cost cannot be negative")
)

type Operation struct {
	ID   uuid.UUID                  `json:"id"`
	Type valueobjects.OperationType `json:"type"`
	Cost int64                      `json:"cost"`
}

func NewOperation(operationType string, cost int64) (*Operation, error) {
	if operationType == "" {
		return nil, ErrOperationTypeIsBlank
	}

	if cost < 0 {
		return nil, ErrOperationCostIsNegative
	}

	oType, err := valueobjects.NewOperationType(operationType)
	if err != nil {
		return nil, err
	}

	o := &Operation{
		ID:   uuid.New(),
		Type: *oType,
		Cost: cost,
	}

	return o, nil
}
