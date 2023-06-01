package db

import (
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
)

func RunMigrations(conf config.Config) {
	m, err := migrate.New("file://internal/db/migrations", conf.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("failed connect to db")
	}
	if err := m.Up(); err != nil {
		log.Fatal().Err(err).Msg("failed run migrations")

	}
}
