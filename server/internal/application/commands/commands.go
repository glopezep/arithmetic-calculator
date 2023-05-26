package commands

type Commands struct {
	CreateUser       CreateUserCommandHandler
	DeleteRecord     DeleteOperationCommandHandler
	ExecuteOperation ExecuteOperationCommandHandler
}
