package entities

import (
	"testing"

	valueobjects "github.com/glopezep/arithmetic-calculator/internal/domain/value_objects"
	"github.com/stretchr/testify/require"
)

func TestNewOperation(t *testing.T) {
	cost := int64(10)
	o, err := NewOperation("addition", cost)

	require.NoError(t, err)
	require.Equal(t, o.Cost, cost)
	require.Equal(t, o.Type, valueobjects.OperationTypeAddition)
}

func TestNewInvalidOperation(t *testing.T) {
	cost := int64(10)
	_, err := NewOperation("", cost)

	require.Error(t, err)
}

func TestNewInvalidOperationCost(t *testing.T) {
	cost := int64(-10)
	_, err := NewOperation("addition", cost)

	require.Error(t, err)
}
