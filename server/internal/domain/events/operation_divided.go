package events

import (
	"github.com/google/uuid"
)

type OperationDivided struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Cost        int64
	UserBalance int64
	FirstValue  int64
	SecondValue int64
}

func (e *OperationDivided) String() string {
	return "OperationDivided"
}

func NewOperationDivided(id, userId uuid.UUID, cost, userBalance, firstValue, secondValue int64) *OperationDivided {
	return &OperationDivided{
		ID:          id,
		UserID:      userId,
		Cost:        cost,
		UserBalance: userBalance,
		FirstValue:  firstValue,
		SecondValue: secondValue,
	}
}
