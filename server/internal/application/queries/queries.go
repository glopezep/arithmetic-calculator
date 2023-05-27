package queries

type Queries struct {
	AuthenticateUser AuthenticateUserQueryHandler
	ListOperations   ListOperationsQueryHandler
	ListRecords      ListRecordsQueryHandler
}
