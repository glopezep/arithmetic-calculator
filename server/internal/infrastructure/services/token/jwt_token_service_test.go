package token

import (
	"testing"

	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"github.com/stretchr/testify/require"
)

func TestJWTTokenServiceSign(t *testing.T) {
	token := NewJwtTokenService()

	u, err := entities.NewUser("test@test.com", "1234")
	require.NoError(t, err)

	tokenString, err := token.Sign(u.ID)

	require.NoError(t, err)
	require.NotEmpty(t, tokenString)
}

func TestJWTTokenServiceVerify(t *testing.T) {
	token := NewJwtTokenService()

	u, err := entities.NewUser("test@test.com", "1234")
	require.NoError(t, err)

	tokenString, err := token.Sign(u.ID)

	require.NoError(t, err)
	require.NotEmpty(t, tokenString)

	claims, err := token.Verify(tokenString)
	require.NoError(t, err)
	require.NotEmpty(t, claims)
	require.Equal(t, claims.RegisteredClaims.Subject, u.ID.String())
}
