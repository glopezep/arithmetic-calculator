package commands

import (
	"context"

	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/stackus/errors"
)

type ExecuteOperationCommand struct{}

type ExecuteOperationCommandHandler struct {
	user      repositories.UserRepository
	operation repositories.OperationRepository
	record    repositories.RecordRepository
}

func (h *ExecuteOperationCommandHandler) Execute(ctx context.Context, c *ExecuteOperationCommand) error {
	userId := uuid.New()
	operationId := uuid.New()

	u, err := h.user.Find(ctx, userId)
	if err != nil {
		return errors.Wrap(err, "failed to find user")
	}

	o, err := h.operation.Find(ctx, operationId)
	if err != nil {
		return errors.Wrap(err, "failed to find operation")
	}

	if err = u.ExecuteOperation(*o); err != nil {
		return errors.Wrap(err, "failed to execute operation")
	}

	if err = h.user.Update(ctx, u); err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	r, err := entities.NewRecord(operationId, userId, o.Cost, "")
	if err != nil {
		return errors.Wrap(err, "failed to create record")
	}

	if err = h.record.Save(ctx, r); err != nil {
		return errors.Wrap(err, "failed to save record")
	}

	return nil
}

func NewExecuteOperationCommandHandler() *ExecuteOperationCommandHandler {
	return &ExecuteOperationCommandHandler{}
}
