package mappers

import (
	"github.com/glopezep/arithmetic-calculator/internal/common"
	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	valueobjects "github.com/glopezep/arithmetic-calculator/internal/domain/value_objects"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/db/models"
)

type UserMapper struct{}

func (m *UserMapper) ToEntity(u models.User) *entities.User {
	return &entities.User{
		AggregateBase: common.AggregateBase{
			ID: u.ID,
		},
		Password: valueobjects.Password(u.Password),
		Email:    valueobjects.Email(u.Email),
		Status:   valueobjects.UserStatusActive,
		Balance:  u.Balance,
	}
}

func NewUserMapper() UserMapper {
	return UserMapper{}
}
