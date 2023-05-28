package application

import (
	"os"

	"github.com/glopezep/arithmetic-calculator/internal/application/commands"
	"github.com/glopezep/arithmetic-calculator/internal/application/queries"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/db"
	eventdispatcher "github.com/glopezep/arithmetic-calculator/internal/infrastructure/event_dispatcher"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/mappers"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/repositories/gorm"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/services/token"
	"github.com/stackus/errors"
)

type Application struct {
	Commands commands.Commands
	Queries  queries.Queries
}

func NewApplication() (*Application, error) {
	domainDispatcher := eventdispatcher.NewEventDispatcher()

	database := db.NewDatabase()

	gormDB, err := database.Open(os.Getenv("DB_SOURCE"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}

	err = db.AutoMigrate(gormDB)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}

	userMapper := mappers.NewUserMapper()
	userRepository := gorm.NewGormUserRepository(gormDB, userMapper)
	operationRepository := gorm.NewGormOperationRepository()
	tokenService := token.NewJwtTokenService()

	return &Application{
		Commands: commands.Commands{
			CreateUser:   *commands.NewCreateUserCommandHandler(userRepository),
			DeleteRecord: *commands.NewDeleteRecordCommandHandler(),
			ExecuteOperation: *commands.NewExecuteOperationCommandHandler(
				tokenService,
				userRepository,
				operationRepository,
				domainDispatcher,
			),
		},
		Queries: queries.Queries{
			AuthenticateUser: *queries.NewAuthenticateUserQueryHandler(userRepository, tokenService),
			ListOperations:   *queries.NewListOperationsQueryHandler(),
			ListRecords:      *queries.NewListRecordsQueryHandler(),
		},
	}, nil
}
