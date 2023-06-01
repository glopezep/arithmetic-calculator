package valueobjects

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewEmail(t *testing.T) {
	email := "test@test.com"

	e, err := NewEmail(email)

	require.NoError(t, err)
	require.Equal(t, e.String(), email)
}

func TestNewInvalidEmail(t *testing.T) {
	_, err := NewEmail("")

	require.Error(t, err)
}
