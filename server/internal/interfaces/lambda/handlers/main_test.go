package handlers

import (
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/glopezep/arithmetic-calculator/internal/application"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDB *gorm.DB
var sqlMock sqlmock.Sqlmock
var app *application.Application

func TestMain(m *testing.M) {
	db, mock, _ := sqlmock.New()

	gormDB, _ = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))

	sqlMock = mock

	a, err := application.NewApplication(gormDB)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize app")
	}

	app = a

	os.Exit(m.Run())
}
