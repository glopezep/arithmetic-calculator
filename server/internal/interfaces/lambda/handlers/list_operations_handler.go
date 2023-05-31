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
)

type ListOperationsHandler struct {
	app *application.Application
}

type ListOperationsRequest struct{}

type ListOperationsResponse struct {
	Items []*entities.Operation `json:"items"`
}

func (h *ListOperationsHandler) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	offset, _ := strconv.Atoi(request.QueryStringParameters["offset"])
	limit, _ := strconv.Atoi(request.QueryStringParameters["limit"])

	operations, err := h.app.Queries.ListOperations.Execute(ctx, &queries.ListOperationsQuery{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(ListOperationsResponse{
		Items: operations,
	})

	if err != nil {
		return nil, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(bytes),
	}, nil
}

func StartListOperationsHandler(app *application.Application) {
	handler := ListOperationsHandler{app}

	lambda.Start(handler.Handle)
}
