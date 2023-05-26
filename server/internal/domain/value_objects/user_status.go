package valueobjects

import "errors"

var (
	ErrInvalidUserStatus = errors.New("invalid user status")
)

type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
)

func (vo UserStatus) String() string {
	return string(vo)
}

func NewUserStatus(v UserStatus) (*UserStatus, error) {
	if err := ValidateUserStatus(v); err != nil {
		return nil, err
	}

	return &v, nil
}

func ValidateUserStatus(v UserStatus) error {
	switch v {
	case
		UserStatusActive, UserStatusInactive:
		return nil
	default:
		return ErrInvalidUserStatus
	}
}
