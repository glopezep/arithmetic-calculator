package queries

import "context"

type ListRecordsQuery struct{}

type ListRecordsQueryHandler struct{}

func (h *ListRecordsQueryHandler) Execute(ctx context.Context, c *ListRecordsQuery) error {
	return nil
}

func NewListRecordsQueryHandler() *ListRecordsQueryHandler {
	return &ListRecordsQueryHandler{}
}
