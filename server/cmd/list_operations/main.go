package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/glopezep/arithmetic-calculator/internal/application"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/config"
	"github.com/glopezep/arithmetic-calculator/internal/interfaces/lambda/handlers"
	"github.com/glopezep/arithmetic-calculator/internal/interfaces/lambda/helpers"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var app *application.Application

func init() {
	conf := config.NewConfig()

	db, err := gorm.Open(postgres.Open(conf.DBSource), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	a, err := application.NewApplication(db)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot initialize app")
	}

	app = a
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return handlers.ListOperationsHandler(ctx, request, app)
}

func main() {
	lambda.Start(helpers.HandleWithContext(handler))
}
