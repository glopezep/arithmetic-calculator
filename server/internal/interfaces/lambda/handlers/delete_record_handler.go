package handlers

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/glopezep/arithmetic-calculator/internal/application"
	"github.com/glopezep/arithmetic-calculator/internal/application/commands"
	"github.com/stackus/errors"
)

type DeleteRecordHandler struct {
	app *application.Application
}

type DeleteRecordRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DeleteRecordResponse struct {
	Token string `json:"token"`
}

func (h *DeleteRecordHandler) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var req DeleteRecordRequest

	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal json")
	}

	err = h.app.Commands.DeleteRecord.Execute(ctx, &commands.DeleteRecordCommand{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create user")
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func StartDeleteRecordHandler(app *application.Application) {
	handler := DeleteRecordHandler{app}

	lambda.Start(handler.Handle)
}
