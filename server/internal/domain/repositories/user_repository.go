package repositories

import (
	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"golang.org/x/net/context"
)

type UserRepository interface {
	Repository[entities.User]
	FindByEmail(context.Context, string) (*entities.User, error)
}
