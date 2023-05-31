package queries

type Queries struct {
	AuthenticateUser AuthenticateUserQueryHandler
	GetUserInfo      GetUserInfoQueryHandler
	ListOperations   ListOperationsQueryHandler
	ListRecords      ListRecordsQueryHandler
}
