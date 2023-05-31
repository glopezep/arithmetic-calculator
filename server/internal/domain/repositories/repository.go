package repositories

import (
	"context"

	"github.com/google/uuid"
)

type Repository[T any] interface {
	Save(ctx context.Context, t *T) error
	Update(ctx context.Context, t *T) error
	Delete(ctx context.Context, id uuid.UUID) error
	Find(ctx context.Context, id uuid.UUID) (*T, error)
	FindAll(ctx context.Context, pageNumber, pageSize int, sortBy, orderBy string) (*PaginatedResult[T], error)
}

type PaginatedResult[T any] struct {
	Items       []*T  `json:"items"`
	TotalCount  int64 `json:"totalCount"`
	Offset      int64 `json:"offset"`
	Limit       int64 `json:"limit"`
	HasNextPage bool  `json:"hasNextPage"`
}
