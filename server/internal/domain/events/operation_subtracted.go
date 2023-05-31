package events

import (
	"github.com/google/uuid"
)

type OperationSubtracted struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Cost        int64
	UserBalance int64
	FirstValue  int64
	SecondValue int64
}

func (e *OperationSubtracted) String() string {
	return "OperationSubtracted"
}

func NewOperationSubtracted(id, userId uuid.UUID, cost, userBalance, firstValue, secondValue int64) *OperationSubtracted {
	return &OperationSubtracted{
		ID:          id,
		UserID:      userId,
		Cost:        cost,
		UserBalance: userBalance,
		FirstValue:  firstValue,
		SecondValue: secondValue,
	}
}
