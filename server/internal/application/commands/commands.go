package commands

type Commands struct {
	CreateUser       CreateUserCommandHandler
	DeleteRecord     DeleteRecordCommandHandler
	ExecuteOperation ExecuteOperationCommandHandler
}
