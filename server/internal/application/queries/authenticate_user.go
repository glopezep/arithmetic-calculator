package queries

import (
	"context"
	"fmt"

	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/services/token"
	"github.com/stackus/errors"
)

type AuthenticateUserQuery struct {
	Email,
	Password string
}

type AuthenticateUserQueryHandler struct {
	user  repositories.UserRepository
	token token.TokenService
}

func (h *AuthenticateUserQueryHandler) Execute(ctx context.Context, c *AuthenticateUserQuery) (string, error) {
	u, err := h.user.FindByEmail(ctx, c.Email)
	if err != nil {
		return "", errors.Wrap(err, "failed to authenticate user")
	}

	fmt.Println(c.Password)
	fmt.Println(u.Password.String())

	if err = u.Password.Compare(c.Password); err != nil {
		return "", errors.Wrap(err, "failed to authenticate user")
	}

	return h.token.Sign(u.ID)
}

func NewAuthenticateUserQueryHandler(
	userRepository repositories.UserRepository,
	tokenService token.TokenService,
) *AuthenticateUserQueryHandler {
	return &AuthenticateUserQueryHandler{
		user:  userRepository,
		token: tokenService,
	}
}
