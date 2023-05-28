package entities

import (
	"errors"

	"github.com/Rhymond/go-money"
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
	Balance  money.Money
}

func (u *User) ExecuteOperation(operation Operation) error {
	ok, err := u.Balance.LessThan(&operation.Cost)

	if err != nil {
		return err
	}

	if ok {
		return ErrBalanceInsufficient
	}

	newBalance, err := u.Balance.Subtract(&operation.Cost)
	if err != nil {
		return err
	}

	u.Balance = *newBalance

	switch operation.Type {
	case valueobjects.OperationTypeDivision:
		u.AddEvent(events.NewOperationDivided())
	case valueobjects.OperationTypeMultiplication:
		u.AddEvent(events.NewOperationMultiplied())
	case valueobjects.OperationTypeRandomString:
		u.AddEvent(events.NewOperationRandomStringGenerated())
	case valueobjects.OperationTypeSquareRoot:
		u.AddEvent(events.NewOperationSquareRooted())
	case valueobjects.OperationTypeSubtraction:
		u.AddEvent(events.NewOperationSubtracted())
	case valueobjects.OperationTypeAddition:
		u.AddEvent(events.NewOperationSummed())
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
		Balance:  *money.New(100, "USD"),
	}, nil
}
