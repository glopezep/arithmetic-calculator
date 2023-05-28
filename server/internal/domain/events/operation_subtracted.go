package events

type OperationSubtracted struct{}

func (e *OperationSubtracted) String() string {
	return "OperationSubtracted"
}

func NewOperationSubtracted() *OperationSubtracted {
	return &OperationSubtracted{}
}
