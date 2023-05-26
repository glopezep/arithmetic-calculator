package commands

import (
	"context"

	"github.com/google/uuid"
)

type DeleteOperationCommand struct {
	ID uuid.UUID
}

type DeleteOperationCommandHandler struct{}

func (h *DeleteOperationCommandHandler) Execute(ctx context.Context, c *DeleteOperationCommand) error {
	return nil
}

func NewDeleteOperationCommandHandler() *DeleteOperationCommandHandler {
	return &DeleteOperationCommandHandler{}
}
