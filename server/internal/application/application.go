package application

import (
	"github.com/glopezep/arithmetic-calculator/internal/application/commands"
	eventhandlers "github.com/glopezep/arithmetic-calculator/internal/application/event_handlers"
	"github.com/glopezep/arithmetic-calculator/internal/application/queries"
	eventdispatcher "github.com/glopezep/arithmetic-calculator/internal/infrastructure/event_dispatcher"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/mappers"
	gormRepositories "github.com/glopezep/arithmetic-calculator/internal/infrastructure/repositories/gorm"
	randomstring "github.com/glopezep/arithmetic-calculator/internal/infrastructure/services/random_string"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/services/token"
	"gorm.io/gorm"
)

type Application struct {
	Commands commands.Commands
	Queries  queries.Queries
}

func NewApplication(conn *gorm.DB) (*Application, error) {
	// conf := config.NewConfig()

	// if conf.Environment == "development" {
	// 	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	// }

	domainDispatcher := eventdispatcher.NewEventDispatcher()
	userMapper := mappers.NewUserMapper()
	operationMapper := mappers.NewOperationMapper()
	recordMapper := mappers.NewRecordMapper()
	userRepository := gormRepositories.NewGormUserRepository(conn, userMapper)
	operationRepository := gormRepositories.NewGormOperationRepository(conn, operationMapper)
	recordRepository := gormRepositories.NewGormRecordRepository(conn, recordMapper)
	tokenService := token.NewJwtTokenService()
	randomStringService := randomstring.NewRandomStringService()
	operationHandlers := eventhandlers.NewOperationHandlers(operationRepository, recordRepository, randomStringService)

	eventhandlers.RegisterOperationHandlers(*operationHandlers, domainDispatcher)

	return &Application{
		Commands: commands.Commands{
			CreateUser:   *commands.NewCreateUserCommandHandler(userRepository),
			DeleteRecord: *commands.NewDeleteRecordCommandHandler(),
			ExecuteOperation: *commands.NewExecuteOperationCommandHandler(
				tokenService,
				userRepository,
				operationRepository,
				recordRepository,
				domainDispatcher,
			),
		},
		Queries: queries.Queries{
			AuthenticateUser: *queries.NewAuthenticateUserQueryHandler(userRepository, tokenService),
			GetUserInfo:      *queries.NewGetUserInfoQueryHandler(tokenService, userRepository),
			ListOperations:   *queries.NewListOperationsQueryHandler(operationRepository),
			ListRecords:      *queries.NewListRecordsQueryHandler(recordRepository),
		},
	}, nil
}
