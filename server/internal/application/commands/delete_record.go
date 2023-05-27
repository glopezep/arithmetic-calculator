package commands

import (
	"context"

	"github.com/google/uuid"
)

type DeleteRecordCommand struct {
	ID uuid.UUID
}

type DeleteRecordCommandHandler struct{}

func (h *DeleteRecordCommandHandler) Execute(ctx context.Context, c *DeleteRecordCommand) error {
	return nil
}

func NewDeleteRecordCommandHandler() *DeleteRecordCommandHandler {
	return &DeleteRecordCommandHandler{}
}
