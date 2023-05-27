package commands

import (
	"context"
	"strings"

	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/services/token"
	"github.com/google/uuid"
	"github.com/stackus/errors"
)

type ExecuteOperationCommand struct {
	OperationID uuid.UUID
}

type ExecuteOperationCommandHandler struct {
	token     token.TokenService
	user      repositories.UserRepository
	operation repositories.OperationRepository
	record    repositories.RecordRepository
}

func (h *ExecuteOperationCommandHandler) Execute(ctx context.Context, c *ExecuteOperationCommand) error {
	auth := ctx.Value("authorization").(string)
	tokenString := strings.Split(auth, " ")[1]

	claims, err := h.token.Verify(tokenString)
	if err != nil {
		return err
	}

	userId, err := uuid.FromBytes([]byte(claims.Issuer))
	if err != nil {
		return err
	}

	u, err := h.user.Find(ctx, userId)
	if err != nil {
		return errors.Wrap(err, "failed to find user")
	}

	o, err := h.operation.Find(ctx, c.OperationID)
	if err != nil {
		return err
	}

	if err = u.ExecuteOperation(*o); err != nil {
		return err
	}

	if err = h.user.Update(ctx, u); err != nil {
		return err
	}

	r, err := entities.NewRecord(c.OperationID, userId, o.Cost, "")
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
