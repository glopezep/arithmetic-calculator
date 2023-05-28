package events

import (
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type OperationSummed struct {
	ID          uuid.UUID
	Cost        money.Money
	FirstValue  int64
	SecondValue int64
}

func (e *OperationSummed) String() string {
	return "OperationSummed"
}

func NewOperationSummed() *OperationSummed {
	return &OperationSummed{}
}
