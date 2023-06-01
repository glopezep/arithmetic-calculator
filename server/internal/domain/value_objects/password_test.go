package valueobjects

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPassword(t *testing.T) {
	pass := "1234"

	e, err := NewPassword(pass)

	require.NoError(t, err)
	require.NoError(t, e.Compare(pass))
}

func TestNewInvalidPassword(t *testing.T) {
	_, err := NewPassword("")

	require.Error(t, err)
}
