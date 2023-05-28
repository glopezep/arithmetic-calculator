package events

import (
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type OperationDivided struct {
	ID          uuid.UUID
	Cost        money.Money
	FirstValue  int64
	SecondValue int64
}

func (e *OperationDivided) String() string {
	return "OperationDivided"
}

func NewOperationDivided() *OperationDivided {
	return &OperationDivided{}
}
