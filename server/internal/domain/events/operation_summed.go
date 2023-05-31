package events

import (
	"github.com/google/uuid"
)

type OperationSummed struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Cost        int64
	UserBalance int64
	FirstValue  int64
	SecondValue int64
}

func (e *OperationSummed) String() string {
	return "OperationSummed"
}

func NewOperationSummed(id, userId uuid.UUID, cost, userBalance, firstValue, secondValue int64) *OperationSummed {
	return &OperationSummed{
		ID:          id,
		UserID:      userId,
		Cost:        cost,
		UserBalance: userBalance,
		FirstValue:  firstValue,
		SecondValue: secondValue,
	}
}
