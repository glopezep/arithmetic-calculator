package queries

import "context"

type ListOperationsQuery struct{}

type ListOperationsQueryHandler struct{}

func (h *ListOperationsQueryHandler) Execute(ctx context.Context, c *ListOperationsQuery) error {
	return nil
}
