package events

import (
	"github.com/google/uuid"
)

type OperationSquareRooted struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Cost        int64
	UserBalance int64
	FirstValue  int64
}

func (e *OperationSquareRooted) String() string {
	return "OperationSquareRooted"
}

func NewOperationSquareRooted(id, userId uuid.UUID, cost, userBalance, firstValue int64) *OperationSquareRooted {
	return &OperationSquareRooted{
		ID:          id,
		UserID:      userId,
		Cost:        cost,
		FirstValue:  firstValue,
		UserBalance: userBalance,
	}
}
