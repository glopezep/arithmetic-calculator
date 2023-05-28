package events

type OperationSquareRooted struct{}

func (e *OperationSquareRooted) String() string {
	return "OperationSquareRooted"
}

func NewOperationSquareRooted() *OperationSquareRooted {
	return &OperationSquareRooted{}
}
