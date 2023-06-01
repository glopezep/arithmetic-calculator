package valueobjects

import "errors"

var (
	ErrInvalidEmail = errors.New("invalid email")
)

type Email string

func (vo Email) String() string {
	return string(vo)
}

func NewEmail(v string) (*Email, error) {
	if err := validateEmail(v); err != nil {
		return nil, err
	}

	vo := Email(v)

	return &vo, nil
}

func validateEmail(v string) error {
	if v == "" {
		return ErrInvalidEmail
	}
	return nil
}
