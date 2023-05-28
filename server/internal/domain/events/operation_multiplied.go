package events

type OperationMultiplied struct{}

func (e *OperationMultiplied) String() string {
	return "OperationMultiplied"
}

func NewOperationMultiplied() *OperationMultiplied {
	return &OperationMultiplied{}
}
