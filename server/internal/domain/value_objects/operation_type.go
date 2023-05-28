package valueobjects

import "errors"

var (
	ErrInvalidOperationType = errors.New("invalid operation type")
)

type OperationType string

const (
	OperationTypeAddition       OperationType = "addition"
	OperationTypeSubtraction    OperationType = "subtraction"
	OperationTypeMultiplication OperationType = "multiplication"
	OperationTypeDivision       OperationType = "division"
	OperationTypeSquareRoot     OperationType = "square_root"
	OperationTypeRandomString   OperationType = "random_string"
)

func (vo OperationType) String() string {
	return string(vo)
}

func NewOperationType(v string) (*OperationType, error) {
	if err := ValidateOperationType(OperationType(v)); err != nil {
		return nil, err
	}

	o := OperationType(v)

	return &o, nil
}

func ValidateOperationType(v OperationType) error {
	switch v {
	case
		OperationTypeAddition,
		OperationTypeSubtraction,
		OperationTypeMultiplication,
		OperationTypeDivision,
		OperationTypeSquareRoot,
		OperationTypeRandomString:
		return nil
	default:
		return ErrInvalidOperationType
	}
}
