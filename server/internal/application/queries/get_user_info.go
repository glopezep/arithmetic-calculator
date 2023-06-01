package queries

import (
	"context"

	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/services/token"
	"github.com/glopezep/arithmetic-calculator/internal/interfaces/lambda/helpers"
)

type GetUserInfoQuery struct {
}

type GetUserInfoQueryHandler struct {
	token token.TokenService
	user  repositories.UserRepository
}

func (h *GetUserInfoQueryHandler) Execute(ctx context.Context, q *GetUserInfoQuery) (*entities.User, error) {
	helperContext := ctx.Value(helpers.ContextKey("context")).(helpers.Context)

	u, err := h.user.Find(ctx, helperContext.UserID)
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
