package events

import (
	"github.com/google/uuid"
)

type OperationMultiplied struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Cost        int64
	UserBalance int64
	FirstValue  int64
	SecondValue int64
}

func (e *OperationMultiplied) String() string {
	return "OperationMultiplied"
}

func NewOperationMultiplied(id, userId uuid.UUID, cost, userBalance, firstValue, secondValue int64) *OperationMultiplied {
	return &OperationMultiplied{
		ID:          id,
		UserID:      userId,
		Cost:        cost,
		UserBalance: userBalance,
		FirstValue:  firstValue,
		SecondValue: secondValue,
	}
}
