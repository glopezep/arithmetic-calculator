package commands

import (
	"context"
	"errors"

	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
	"github.com/glopezep/arithmetic-calculator/internal/interfaces/lambda/helpers"
	"github.com/google/uuid"
)

var (
	ErrInvalidResourceOwner = errors.New("this user cannot delete this resource")
)

type DeleteRecordCommand struct {
	ID uuid.UUID
}

type DeleteRecordCommandHandler struct {
	record repositories.RecordRepository
}

func (h *DeleteRecordCommandHandler) Execute(ctx context.Context, c *DeleteRecordCommand) error {
	helperContext := ctx.Value(helpers.ContextKey("context")).(helpers.Context)

	r, err := h.record.Find(ctx, c.ID)
	if err != nil {
		return err
	}

	if r.UserID != helperContext.UserID {
		return ErrInvalidResourceOwner
	}

	err = h.record.Delete(ctx, c.ID)
	if err != nil {
		return err
	}

	return nil
}

func NewDeleteRecordCommandHandler(record repositories.RecordRepository) *DeleteRecordCommandHandler {
	return &DeleteRecordCommandHandler{
		record,
	}
}
