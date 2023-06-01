package gorm

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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

	var (
		id    = uuid.New()
		email = "test@test.com"
	)

	sqlMock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE (id = $1) LIMIT 1`)).
		WithArgs(id.String()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email"}).
			AddRow(id.String(), email))

	_, err := repository.Find(context.Background(), id)

	require.NoError(t, err)
}

func TestGormUserRepositoryFindAll(t *testing.T) {
	mapper := mappers.NewUserMapper()
	repository := NewGormUserRepository(gormDB, mapper)

	const sqlSelectAll = `SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL ORDER BY created_at desc LIMIT 10"`

	sqlMock.ExpectQuery(sqlSelectAll).WillReturnRows(sqlmock.NewRows(nil))

	_, err := repository.FindAll(context.Background(), 1, 10, "created_at", "desc")

	require.NoError(t, err)
}
