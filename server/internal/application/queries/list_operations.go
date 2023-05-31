package queries

import (
	"context"

	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
)

type ListOperationsQuery struct {
	Offset, Limit   int
	SortBy, OrderBy string
}

type ListOperationsQueryHandler struct {
	operation repositories.OperationRepository
}

func (h *ListOperationsQueryHandler) Execute(ctx context.Context, q *ListOperationsQuery) ([]*entities.Operation, error) {
	return h.operation.FindAll(ctx, q.Offset, q.Limit, q.SortBy, q.OrderBy)
}

func NewListOperationsQueryHandler(operation repositories.OperationRepository) *ListOperationsQueryHandler {
	return &ListOperationsQueryHandler{
		operation,
	}
}
