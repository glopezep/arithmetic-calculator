package commands

import (
	"context"
	"strings"

	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
	eventdispatcher "github.com/glopezep/arithmetic-calculator/internal/infrastructure/event_dispatcher"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/services/token"
	"github.com/google/uuid"
)

type ExecuteOperationCommand struct {
	OperationID uuid.UUID
}

type ExecuteOperationCommandHandler struct {
	token           token.TokenService
	user            repositories.UserRepository
	operation       repositories.OperationRepository
	domainPublisher eventdispatcher.EventPublisher
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
		return err
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

	h.domainPublisher.Publish(ctx, u.GetEvents()...)

	return nil
}

func NewExecuteOperationCommandHandler(
	token token.TokenService,
	user repositories.UserRepository,
	operation repositories.OperationRepository,
	domainPublisher eventdispatcher.EventPublisher,
) *ExecuteOperationCommandHandler {
	return &ExecuteOperationCommandHandler{
		token,
		user,
		operation,
		domainPublisher,
	}
}
