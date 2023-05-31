package handlers

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/glopezep/arithmetic-calculator/internal/application"
	"github.com/glopezep/arithmetic-calculator/internal/application/queries"
	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
)

type ListRecordsHandler struct {
	app *application.Application
}

type ListRecordsRequest struct{}

type ListRecordsResponse struct {
	Items []*entities.Record `json:"items"`
}

func (h *ListRecordsHandler) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	records, err := h.app.Queries.ListRecords.Execute(ctx, &queries.ListRecordsQuery{})
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(ListRecordsResponse{
		Items: records,
	})

	if err != nil {
		return nil, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(bytes),
	}, nil
}

func StartListRecordsHandler(app *application.Application) {
	handler := ListRecordsHandler{app}

	lambda.Start(handler.Handle)
}
