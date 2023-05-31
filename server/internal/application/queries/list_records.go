package queries

import (
	"context"

	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
)

type ListRecordsQuery struct {
	Offset, Limit   int
	SortBy, OrderBy string
}

type ListRecordsQueryHandler struct {
	record repositories.RecordRepository
}

func (h *ListRecordsQueryHandler) Execute(ctx context.Context, q *ListRecordsQuery) (*repositories.PaginatedResult[entities.Record], error) {
	return h.record.FindAll(ctx, q.Offset, q.Limit, q.SortBy, q.OrderBy)
}

func NewListRecordsQueryHandler(record repositories.RecordRepository) *ListRecordsQueryHandler {
	return &ListRecordsQueryHandler{
		record,
	}
}
