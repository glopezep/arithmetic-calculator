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

type ListRecordsHandler struct {
	app *application.Application
}

type ListRecordsRequest struct {
	OperationID uuid.UUID
}

type ListRecordsResponse struct{}

func (h *ListRecordsHandler) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	ctx = context.WithValue(ctx, "authorization", request.Headers["authorization"])

	var req ListRecordsRequest

	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return nil, err
	}

	err = h.app.Queries.ListRecords.Execute(ctx, &queries.ListRecordsQuery{})

	if err != nil {
		return nil, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func StartListRecordsHandler(app *application.Application) {
	handler := ListRecordsHandler{app}

	lambda.Start(handler.Handle)
}
