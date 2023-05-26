package commands

import (
	"context"

	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
	"github.com/stackus/errors"
)

type CreateUserCommand struct {
	Email,
	Password string
}

type CreateUserCommandHandler struct {
	user repositories.UserRepository
}

func (h *CreateUserCommandHandler) Execute(ctx context.Context, c *CreateUserCommand) error {
	u, err := entities.NewUser(c.Email, c.Password)
	if err != nil {
		return errors.Wrap(err, "failed to create user")
	}

	h.user.Save(ctx, u)

	return nil

}

func NewCreateUserCommandHandler(
	userRepository repositories.UserRepository,
) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{
		user: userRepository,
	}
}
