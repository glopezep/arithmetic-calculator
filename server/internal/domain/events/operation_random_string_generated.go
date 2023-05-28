package events

import (
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type OperationRandomStringGenerated struct {
	ID   uuid.UUID
	Cost money.Money
}

func (e *OperationRandomStringGenerated) String() string {
	return "OperationRandomStringGenerated"
}

func NewOperationRandomStringGenerated() *OperationRandomStringGenerated {
	return &OperationRandomStringGenerated{}
}
