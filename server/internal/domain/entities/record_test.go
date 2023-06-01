package entities

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewRecord(t *testing.T) {
	cost := int64(10)
	o, _ := NewOperation("addition", cost)

	userId := uuid.New()
	userBalance := int64(100)

	r, err := NewRecord(o.ID, userId, cost, userBalance, "10")

	require.NoError(t, err)
	require.Equal(t, r.Amount, o.Cost)
	require.Equal(t, r.OperationID, o.ID)
	require.Equal(t, r.OperationResponse, "10")
	require.Equal(t, r.UserBalance, userBalance)
	require.Equal(t, r.UserID, userId)
}

func TestNewInvalidRecord(t *testing.T) {
	cost := int64(10)
	o, _ := NewOperation("addition", cost)

	userId := uuid.New()
	userBalance := int64(100)

	_, err := NewRecord(o.ID, userId, cost, userBalance, "")

	require.Error(t, err)
}
