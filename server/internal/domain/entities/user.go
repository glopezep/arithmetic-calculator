package entities

import (
	"errors"

	"github.com/glopezep/arithmetic-calculator/internal/common"
	"github.com/glopezep/arithmetic-calculator/internal/domain/events"
	valueobjects "github.com/glopezep/arithmetic-calculator/internal/domain/value_objects"
	"github.com/google/uuid"
)

var (
	ErrEmailIsBlank        = errors.New("the user email cannot be blank")
	ErrPasswordIsBlank     = errors.New("the user password cannot be blank")
	ErrBalanceInsufficient = errors.New("the user balance is insufficient to perform the operation")
)

type User struct {
	common.AggregateBase
	ID       uuid.UUID
	Email    valueobjects.Email
	Password valueobjects.Password
	Status   valueobjects.UserStatus
	Balance  int64
}

func (u *User) ExecuteOperation(operation *Operation, fistValue, secondValue int64) error {
	ok := u.Balance < operation.Cost

	if ok {
		return ErrBalanceInsufficient
	}

	newBalance := u.Balance - operation.Cost
	u.Balance = newBalance

	switch operation.Type {
	case valueobjects.OperationTypeDivision:
		u.AddEvent(events.NewOperationDivided(
			operation.ID,
			u.ID,
			operation.Cost,
			newBalance,
			fistValue,
			secondValue,
		))
	case valueobjects.OperationTypeMultiplication:
		u.AddEvent(events.NewOperationMultiplied(
			operation.ID,
			u.ID,
			operation.Cost,
			newBalance,
			fistValue,
			secondValue,
		))
	case valueobjects.OperationTypeRandomString:
		u.AddEvent(events.NewOperationRandomStringGenerated(
			operation.ID,
			u.ID,
			operation.Cost,
			newBalance,
		))
	case valueobjects.OperationTypeSquareRoot:
		u.AddEvent(events.NewOperationSquareRooted(
			operation.ID,
			u.ID,
			operation.Cost,
			newBalance,
			fistValue,
		))
	case valueobjects.OperationTypeSubtraction:
		u.AddEvent(events.NewOperationSubtracted(
			operation.ID,
			u.ID,
			operation.Cost,
			newBalance,
			fistValue,
			secondValue,
		))
	case valueobjects.OperationTypeAddition:
		u.AddEvent(events.NewOperationSummed(
			operation.ID,
			u.ID,
			operation.Cost,
			newBalance,
			fistValue,
			secondValue,
		))
	default:
	}

	return nil
}

func NewUser(email, password string) (*User, error) {
	if email == "" {
		return nil, ErrEmailIsBlank
	}

	if password == "" {
		return nil, ErrPasswordIsBlank
	}

	e, err := valueobjects.NewEmail(email)
	if err != nil {
		return nil, err
	}

	p, err := valueobjects.NewPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       uuid.New(),
		Email:    *e,
		Password: *p,
		Status:   valueobjects.UserStatusActive,
		Balance:  100,
	}, nil
}
