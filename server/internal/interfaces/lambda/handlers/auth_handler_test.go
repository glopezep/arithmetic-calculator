package handlers

import (
	"context"
	"encoding/json"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aws/aws-lambda-go/events"
	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/db/models"
	"github.com/stretchr/testify/require"
)

func TestAuthHandler(t *testing.T) {

	t.Run("Successful response", func(t *testing.T) {
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

		handler := AuthHandler{app}
		bytes, err := json.Marshal(AuthRequest{
			Email:    "test@test.com",
			Password: "1234",
		})

		require.NoError(t, err)

		res, err := handler.Handle(context.Background(), events.APIGatewayProxyRequest{
			Body: string(bytes),
		})

		require.NoError(t, err)

		var body AuthResponse

		err = json.Unmarshal([]byte(res.Body), &body)

		require.NoError(t, err)
		require.NoError(t, err)
		require.NotEmpty(t, body.Token)
	})

	t.Run("Successful", func(t *testing.T) {
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

		handler := AuthHandler{app}
		bytes, err := json.Marshal(AuthRequest{
			Email:    "test@test.com",
			Password: "1234",
		})

		require.NoError(t, err)

		res, err := handler.Handle(context.Background(), events.APIGatewayProxyRequest{
			Body: string(bytes),
		})

		require.NoError(t, err)

		var body AuthResponse

		err = json.Unmarshal([]byte(res.Body), &body)

		require.NoError(t, err)
		require.NoError(t, err)
		require.NotEmpty(t, body.Token)
	})

	t.Run("Non successful", func(t *testing.T) {
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

		handler := AuthHandler{app}
		bytes, err := json.Marshal(AuthRequest{
			Email:    "test@test.com",
			Password: "12345",
		})

		require.NoError(t, err)

		_, err = handler.Handle(context.Background(), events.APIGatewayProxyRequest{
			Body: string(bytes),
		})

		require.Error(t, err)
	})
}
