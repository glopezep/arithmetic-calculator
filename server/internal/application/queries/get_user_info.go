package queries

import (
	"context"

	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/services/token"
	"github.com/glopezep/arithmetic-calculator/internal/interfaces/lambda/utils"
	"github.com/google/uuid"
)

type GetUserInfoQuery struct {
}

type GetUserInfoQueryHandler struct {
	token token.TokenService
	user  repositories.UserRepository
}

func (h *GetUserInfoQueryHandler) Execute(ctx context.Context, q *GetUserInfoQuery) (*entities.User, error) {
	token := ctx.Value(utils.ContextKey("token")).(string)

	claims, err := h.token.Verify(token)
	if err != nil {
		return nil, err
	}

	userId := uuid.MustParse(claims.RegisteredClaims.Subject)

	u, err := h.user.Find(ctx, userId)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func NewGetUserInfoQueryHandler(token token.TokenService, user repositories.UserRepository) *GetUserInfoQueryHandler {
	return &GetUserInfoQueryHandler{
		token,
		user,
	}
}
