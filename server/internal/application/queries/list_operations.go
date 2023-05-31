package queries

import (
	"context"

	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
)

type ListOperationsQuery struct{}

type ListOperationsQueryHandler struct {
	operation repositories.OperationRepository
}

func (h *ListOperationsQueryHandler) Execute(ctx context.Context, c *ListOperationsQuery) ([]*entities.Operation, error) {
	return h.operation.FindAll(ctx)
}

func NewListOperationsQueryHandler(operation repositories.OperationRepository) *ListOperationsQueryHandler {
	return &ListOperationsQueryHandler{
		operation,
	}
}
