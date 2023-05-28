package entities

import (
	"errors"

	"github.com/Rhymond/go-money"
	"github.com/glopezep/arithmetic-calculator/internal/common"

	"github.com/glopezep/arithmetic-calculator/internal/domain/events"
	valueobjects "github.com/glopezep/arithmetic-calculator/internal/domain/value_objects"
	"github.com/google/uuid"
)

var (
	ErrOperationTypeIsBlank    = errors.New("the operation type cannot be blank")
	ErrOperationCostIsBlank    = errors.New("the operation cost cannot be blank")
	ErrOperationCostIsNegative = errors.New("the operation cost cannot be negative")
)

type Operation struct {
	common.AggregateBase
	Type valueobjects.OperationType
	Cost money.Money
}

func NewOperation(operationType string, cost money.Money) (*Operation, error) {
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

	o := &Operation{
		AggregateBase: common.AggregateBase{
			ID: uuid.New(),
		},
		Type: *oType,
		Cost: cost,
	}

	o.AddEvent(events.NewOperationCreated())

	return o, nil
}
