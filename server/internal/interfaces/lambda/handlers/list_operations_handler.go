package handlers

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/glopezep/arithmetic-calculator/internal/application"
	"github.com/glopezep/arithmetic-calculator/internal/application/queries"
	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
)

type ListOperationsRequest struct{}

type ListOperationsResponse struct {
	*repositories.PaginatedResult[entities.Operation]
}

func ListOperationsHandler(ctx context.Context, request events.APIGatewayProxyRequest, app *application.Application) (*events.APIGatewayProxyResponse, error) {
	offset, _ := strconv.Atoi(request.QueryStringParameters["offset"])
	limit, _ := strconv.Atoi(request.QueryStringParameters["limit"])
	sortBy := request.QueryStringParameters["sort_by"]
	orderBy := request.QueryStringParameters["order_by"]

	operations, err := app.Queries.ListOperations.Execute(ctx, &queries.ListOperationsQuery{
		Offset:  offset,
		Limit:   limit,
		SortBy:  sortBy,
		OrderBy: orderBy,
	})
	if err != nil {
		return &events.APIGatewayProxyResponse{}, err
	}

	bytes, err := json.Marshal(ListOperationsResponse{
		PaginatedResult: operations,
	})

	if err != nil {
		return &events.APIGatewayProxyResponse{}, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(bytes),
	}, nil
}
