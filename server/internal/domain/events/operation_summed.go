package events

type OperationSummed struct{}

func (e *OperationSummed) String() string {
	return "OperationSummed"
}

func NewOperationSummed() *OperationSummed {
	return &OperationSummed{}
}
