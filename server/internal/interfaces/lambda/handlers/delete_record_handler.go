package handlers

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/glopezep/arithmetic-calculator/internal/application"
	"github.com/glopezep/arithmetic-calculator/internal/application/commands"
	"github.com/glopezep/arithmetic-calculator/internal/interfaces/lambda/helpers"
	"github.com/google/uuid"
)

type DeleteRecordHandler struct {
	app *application.Application
}

type DeleteRecordRequest struct {
	ID string `json:"id"`
}

type DeleteRecordResponse struct{}

func (h *DeleteRecordHandler) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	err := h.app.Commands.DeleteRecord.Execute(ctx, &commands.DeleteRecordCommand{
		ID: uuid.MustParse(request.PathParameters["id"]),
	})

	if err != nil {
		return &events.APIGatewayProxyResponse{}, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func StartDeleteRecordHandler(app *application.Application) {
	handler := DeleteRecordHandler{app}

	lambda.Start(helpers.HandleWithContext(handler.Handle))
}
