package handlers

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/glopezep/arithmetic-calculator/internal/application"
	"github.com/glopezep/arithmetic-calculator/internal/application/queries"
	"github.com/google/uuid"
)

type ListOperationsHandler struct {
	app *application.Application
}

type ListOperationsRequest struct {
	OperationID uuid.UUID
}

type ListOperationsResponse struct{}

func (h *ListOperationsHandler) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	ctx = context.WithValue(ctx, "authorization", request.Headers["authorization"])

	var req ListOperationsRequest

	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return nil, err
	}

	err = h.app.Queries.ListOperations.Execute(ctx, &queries.ListOperationsQuery{})

	if err != nil {
		return nil, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func StartListOperationsHandler(app *application.Application) {
	handler := ListOperationsHandler{app}

	lambda.Start(handler.Handle)
}
