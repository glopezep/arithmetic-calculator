package events

import (
	"github.com/google/uuid"
)

type OperationCreated struct {
	ID uuid.UUID
}

func (e *OperationCreated) String() string {
	return "OperationCreated"
}

func NewOperationCreated(id uuid.UUID) *OperationCreated {
	return &OperationCreated{
		ID: id,
	}
}
