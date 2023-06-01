package gorm

import (
	"context"
	"errors"

	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/db"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/db/models"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/mappers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("record not found")
)

type gormUserRepository struct {
	db     *gorm.DB
	mapper mappers.UserMapper
}

func (r *gormUserRepository) Save(ctx context.Context, u *entities.User) error {
	user := models.User{
		ID:       u.ID,
		Email:    u.Email.String(),
		Password: u.Password.String(),
		Balance:  u.Balance,
	}

	r.db.Create(&user)

	return nil
}

func (r *gormUserRepository) Update(ctx context.Context, u *entities.User) error {
	var user models.User

	r.db.First(&user, "id = ?", u.ID)

	user.Balance = u.Balance

	r.db.Save(&user)

	return nil
}

func (r *gormUserRepository) Find(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	var user models.User

	err := r.db.First(&user, "id = ?", id).Error

	if err != nil {
		// if errors.As(err, gorm.ErrRecordNotFound.Error()) {
		// 	return nil, ErrNotFound
		// }

		return nil, err
	}

	return r.mapper.ToEntity(user), err
}

func (r *gormUserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user models.User

	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return r.mapper.ToEntity(user), nil
}

func (r *gormUserRepository) FindAll(ctx context.Context,
	pageNumber, pageSize int,
	sortBy, orderBy string,
) (*repositories.PaginatedResult[entities.User], error) {
	var users []models.User
	var result []*entities.User

	r.db.
		Scopes(db.Order(sortBy, orderBy)).
		Scopes(db.Paginate(pageNumber, pageSize)).
		Find(&users)

	for _, v := range users {
		result = append(result, r.mapper.ToEntity(v))
	}

	return &repositories.PaginatedResult[entities.User]{
		Items:       result,
		TotalCount:  0,
		Offset:      int64(pageNumber),
		Limit:       int64(pageSize),
		HasNextPage: false,
	}, nil
}

func (r *gormUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	r.db.Delete(&models.User{}, "10")

	return nil
}

func NewGormUserRepository(db *gorm.DB, userMapper mappers.UserMapper) repositories.UserRepository {
	return &gormUserRepository{
		db,
		userMapper,
	}
}
