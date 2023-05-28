package events

import (
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type OperationSubtracted struct {
	ID          uuid.UUID
	Cost        money.Money
	FirstValue  int64
	SecondValue int64
}

func (e *OperationSubtracted) String() string {
	return "OperationSubtracted"
}

func NewOperationSubtracted() *OperationSubtracted {
	return &OperationSubtracted{}
}
