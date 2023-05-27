package entities

import (
	"errors"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

var (
	ErrOperationResponseIsBlank = errors.New("the operation response cannot be blank")
)

type Record struct {
	ID                uuid.UUID
	OperationID       uuid.UUID
	UserID            uuid.UUID
	Amount            money.Money
	UserBalance       money.Money
	OperationResponse string
	Date              time.Time
}

func NewRecord(
	operationId,
	userId uuid.UUID,
	amount money.Money,
	operationResponse string,
) (*Record, error) {
	if operationResponse == "" {
		return nil, ErrOperationResponseIsBlank
	}

	return &Record{
		ID:                uuid.New(),
		OperationID:       operationId,
		UserID:            userId,
		Amount:            amount,
		OperationResponse: operationResponse,
		Date:              time.Now(),
	}, nil
}
