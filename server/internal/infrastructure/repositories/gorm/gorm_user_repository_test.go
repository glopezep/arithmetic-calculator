package gorm

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/db/models"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/mappers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type GormDBMock struct {
	mock.Mock
}

func TestGormUserRepositoryFind(t *testing.T) {
	mapper := mappers.NewUserMapper()
	repository := NewGormUserRepository(gormDB, mapper)

	u, _ := entities.NewUser("test@test.com", "1234")

	um := &models.User{
		ID:       u.ID,
		Email:    u.Email.String(),
		Password: u.Password.String(),
		Balance:  u.Balance,
	}

	sql := `SELECT * FROM "users" WHERE id = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`

	rows := sqlmock.
		NewRows([]string{"id", "email", "password", "balance", "created_at"}).
		AddRow(um.ID, um.Email, um.Password, u.Balance, um.CreatedAt)

	sqlMock.
		ExpectQuery(regexp.QuoteMeta(sql)).
		WithArgs(u.ID.String()).
		WillReturnRows(rows)

	uFromDB, err := repository.Find(context.Background(), u.ID)

	require.NoError(t, err)
	require.Equal(t, u, uFromDB)
}

func TestGormUserRepositoryFindByEmail(t *testing.T) {
	mapper := mappers.NewUserMapper()
	repository := NewGormUserRepository(gormDB, mapper)

	u, _ := entities.NewUser("test@test.com", "1234")

	um := &models.User{
		ID:       u.ID,
		Email:    u.Email.String(),
		Password: u.Password.String(),
		Balance:  u.Balance,
	}

	sql := `SELECT * FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`

	rows := sqlmock.
		NewRows([]string{"id", "email", "password", "balance", "created_at"}).
		AddRow(um.ID, um.Email, um.Password, u.Balance, um.CreatedAt)

	sqlMock.
		ExpectQuery(regexp.QuoteMeta(sql)).
		WithArgs(u.Email.String()).
		WillReturnRows(rows)

	uFromDB, err := repository.FindByEmail(context.Background(), u.Email.String())

	require.NoError(t, err)
	require.Equal(t, u, uFromDB)
}

func TestGormUserRepositoryFindNotFound(t *testing.T) {
	mapper := mappers.NewUserMapper()
	repository := NewGormUserRepository(gormDB, mapper)

	u, _ := entities.NewUser("test@test.com", "1234")

	um := &models.User{
		ID:       u.ID,
		Email:    u.Email.String(),
		Password: u.Password.String(),
		Balance:  u.Balance,
	}

	sql := `SELECT * FROM "users" WHERE id = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`

	rows := sqlmock.
		NewRows([]string{"id", "email", "password", "balance", "created_at"}).
		AddRow(um.ID, um.Email, um.Password, u.Balance, um.CreatedAt)

	sqlMock.
		ExpectQuery(regexp.QuoteMeta(sql)).
		WithArgs(u.ID.String()).
		WillReturnRows(rows)

	_, err := repository.Find(context.Background(), uuid.New())

	require.Error(t, err)
}
