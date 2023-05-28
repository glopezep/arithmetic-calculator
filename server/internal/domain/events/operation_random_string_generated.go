package events

type OperationRandomStringGenerated struct{}

func (e *OperationRandomStringGenerated) String() string {
	return "OperationRandomStringGenerated"
}

func NewOperationRandomStringGenerated() *OperationRandomStringGenerated {
	return &OperationRandomStringGenerated{}
}
