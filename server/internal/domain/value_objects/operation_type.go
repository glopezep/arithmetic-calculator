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

func NewOperationType(v OperationType) (*OperationType, error) {
	if err := ValidateOperationType(v); err != nil {
		return nil, err
	}

	return &v, nil
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
