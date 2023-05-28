package events

import (
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type OperationMultiplied struct {
	ID          uuid.UUID
	Cost        money.Money
	FirstValue  int64
	SecondValue int64
}

func (e *OperationMultiplied) String() string {
	return "OperationMultiplied"
}

func NewOperationMultiplied() *OperationMultiplied {
	return &OperationMultiplied{}
}
