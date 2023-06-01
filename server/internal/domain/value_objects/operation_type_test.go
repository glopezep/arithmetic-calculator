package valueobjects

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewOperationType(t *testing.T) {
	o, err := NewOperationType("addition")

	require.NoError(t, err)
	require.Equal(t, *o, OperationTypeAddition)

	o, err = NewOperationType("subtraction")

	require.NoError(t, err)
	require.Equal(t, *o, OperationTypeSubtraction)

	o, err = NewOperationType("multiplication")

	require.NoError(t, err)
	require.Equal(t, *o, OperationTypeMultiplication)

	o, err = NewOperationType("division")

	require.NoError(t, err)
	require.Equal(t, *o, OperationTypeDivision)

	o, err = NewOperationType("square_root")

	require.NoError(t, err)
	require.Equal(t, *o, OperationTypeSquareRoot)

	o, err = NewOperationType("random_string")

	require.NoError(t, err)
	require.Equal(t, *o, OperationTypeRandomString)
}

func TestNewInvalidOperationType(t *testing.T) {
	_, err := NewOperationType("")

	require.Error(t, err)
}
