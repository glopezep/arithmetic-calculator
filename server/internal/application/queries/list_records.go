package queries

import (
	"context"

	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
)

type ListRecordsQuery struct{}

type ListRecordsQueryHandler struct {
	record repositories.RecordRepository
}

func (h *ListRecordsQueryHandler) Execute(ctx context.Context, c *ListRecordsQuery) ([]*entities.Record, error) {
	return h.record.FindAll(ctx)
}

func NewListRecordsQueryHandler(record repositories.RecordRepository) *ListRecordsQueryHandler {
	return &ListRecordsQueryHandler{
		record,
	}
}
