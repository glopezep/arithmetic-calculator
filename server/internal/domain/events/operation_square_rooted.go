package events

import (
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type OperationSquareRooted struct {
	ID          uuid.UUID
	Cost        money.Money
	FirstValue  int64
	SecondValue int64
}

func (e *OperationSquareRooted) String() string {
	return "OperationSquareRooted"
}

func NewOperationSquareRooted() *OperationSquareRooted {
	return &OperationSquareRooted{}
}
