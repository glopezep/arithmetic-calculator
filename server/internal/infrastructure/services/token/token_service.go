package token

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type TokenService interface {
	Sign(userId uuid.UUID) (string, error)
	Verify(token string) (*Claims, error)
}
