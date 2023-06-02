package handlers

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/glopezep/arithmetic-calculator/internal/application"
	"github.com/glopezep/arithmetic-calculator/internal/application/queries"
	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
	"github.com/glopezep/arithmetic-calculator/internal/interfaces/lambda/helpers"
)

type ListOperationsHandler struct {
	app *application.Application
}

type ListOperationsRequest struct{}

type ListOperationsResponse struct {
	*repositories.PaginatedResult[entities.Operation]
}

func (h *ListOperationsHandler) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	offset, _ := strconv.Atoi(request.QueryStringParameters["offset"])
	limit, _ := strconv.Atoi(request.QueryStringParameters["limit"])
	sortBy := request.QueryStringParameters["sort_by"]
	orderBy := request.QueryStringParameters["order_by"]

	operations, err := h.app.Queries.ListOperations.Execute(ctx, &queries.ListOperationsQuery{
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

func StartListOperationsHandler(app *application.Application) {
	handler := ListOperationsHandler{app}

	lambda.Start(helpers.HandleWithContext(handler.Handle))
}
