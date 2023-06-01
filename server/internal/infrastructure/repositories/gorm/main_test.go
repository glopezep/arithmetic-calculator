package gorm

import (
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDB *gorm.DB
var sqlMock sqlmock.Sqlmock

func TestMain(m *testing.M) {
	db, mock, _ := sqlmock.New()

	gormDB, _ = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))

	sqlMock = mock

	os.Exit(m.Run())
}
