package valueobjects

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidPassword = errors.New("invalid Password")
)

type Password string

func (p Password) Compare(pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(p.String()), []byte(pass))
}

func (p Password) String() string {
	return string(p)
}

func NewPassword(v string) (*Password, error) {
	if err := validatePassword(v); err != nil {
		return nil, ErrInvalidPassword
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(v), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	p := Password(string(hash))

	return &p, nil
}

func validatePassword(v string) error {
	if v == "" {
		return ErrInvalidPassword
	}

	return nil
}
