package events

import (
	"github.com/google/uuid"
)

type OperationRandomStringGenerated struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Cost        int64
	UserBalance int64
}

func (e *OperationRandomStringGenerated) String() string {
	return "OperationRandomStringGenerated"
}

func NewOperationRandomStringGenerated(id, userId uuid.UUID, cost, userBalance int64) *OperationRandomStringGenerated {
	return &OperationRandomStringGenerated{
		ID:          id,
		UserID:      userId,
		Cost:        cost,
		UserBalance: userBalance,
	}
}
