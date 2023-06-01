package entities

import (
	"testing"

	valueobjects "github.com/glopezep/arithmetic-calculator/internal/domain/value_objects"
	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	email := "test@test.com"
	pass := "1234"
	initialBalance := int64(100)

	u, err := NewUser(email, pass)

	require.NoError(t, err)
	require.Equal(t, u.Email, valueobjects.Email(email))
	require.Equal(t, u.Balance, initialBalance)
	require.Equal(t, u.Status, valueobjects.UserStatus("active"))
	require.NoError(t, u.Password.Compare(pass))
}

func TestNewInvalidUserEmail(t *testing.T) {
	email := ""
	pass := "1234"

	_, err := NewUser(email, pass)

	require.Error(t, err)
}

func TestNewInvalidUserPassword(t *testing.T) {
	email := "test@test.com"
	pass := ""

	_, err := NewUser(email, pass)

	require.Error(t, err)
}
