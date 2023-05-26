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
	FindAll(ctx context.Context) ([]*T, error)
}
