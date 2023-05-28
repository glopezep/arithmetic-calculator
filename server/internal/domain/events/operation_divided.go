package events

type OperationDivided struct{}

func (e *OperationDivided) String() string {
	return "OperationDivided"
}

func NewOperationDivided() *OperationDivided {
	return &OperationDivided{}
}
