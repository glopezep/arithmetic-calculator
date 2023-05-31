package commands

import (
	"context"

	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
	eventdispatcher "github.com/glopezep/arithmetic-calculator/internal/infrastructure/event_dispatcher"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/services/token"
	"github.com/glopezep/arithmetic-calculator/internal/interfaces/lambda/utils"
	"github.com/google/uuid"
)

type ExecuteOperationCommand struct {
	OperationID uuid.UUID
	FirstValue  int64
	SecondValue int64
}

type ExecuteOperationCommandHandler struct {
	token           token.TokenService
	user            repositories.UserRepository
	operation       repositories.OperationRepository
	record          repositories.RecordRepository
	domainPublisher eventdispatcher.EventPublisher
}

func (h *ExecuteOperationCommandHandler) Execute(ctx context.Context, c *ExecuteOperationCommand) error {
	token := ctx.Value(utils.ContextKey("token")).(string)

	claims, err := h.token.Verify(token)
	if err != nil {
		return err
	}

	userId := uuid.MustParse(claims.RegisteredClaims.Subject)

	u, err := h.user.Find(ctx, userId)
	if err != nil {
		return err
	}

	o, err := h.operation.Find(ctx, c.OperationID)
	if err != nil {
		return err
	}

	if err = u.ExecuteOperation(o, c.FirstValue, c.SecondValue); err != nil {
		return err
	}

	if err = h.user.Update(ctx, u); err != nil {
		return err
	}

	h.domainPublisher.Publish(ctx, u.GetEvents()...)

	return nil
}

func NewExecuteOperationCommandHandler(
	token token.TokenService,
	user repositories.UserRepository,
	operation repositories.OperationRepository,
	record repositories.RecordRepository,
	domainPublisher eventdispatcher.EventPublisher,
) *ExecuteOperationCommandHandler {
	return &ExecuteOperationCommandHandler{
		token,
		user,
		operation,
		record,
		domainPublisher,
	}
}
