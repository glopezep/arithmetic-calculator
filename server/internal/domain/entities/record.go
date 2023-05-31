package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrOperationResponseIsBlank = errors.New("the operation response cannot be blank")
)

type Record struct {
	ID                uuid.UUID `json:"id"`
	OperationID       uuid.UUID `json:"operationId"`
	UserID            uuid.UUID `json:"userId"`
	Amount            int64     `json:"amount"`
	UserBalance       int64     `json:"userBalance"`
	OperationResponse string    `json:"operationResponse"`
	CreatedAt         time.Time `json:"createdAt"`
}

func NewRecord(
	operationId,
	userId uuid.UUID,
	amount,
	userBalance int64,
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
		UserBalance:       userBalance,
		OperationResponse: operationResponse,
		CreatedAt:         time.Now(),
	}, nil
}
