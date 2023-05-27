package token

import "github.com/google/uuid"

type TokenService interface {
	Sign(userId uuid.UUID) (string, error)
	Verify(token string) (*Claims, error)
}
