package handlers

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/glopezep/arithmetic-calculator/internal/application"
	"github.com/glopezep/arithmetic-calculator/internal/application/commands"
	"github.com/google/uuid"
)

type ExecuteOperationHandler struct {
	app *application.Application
}

type ExecuteOperationRequest struct {
	OperationID uuid.UUID
}

type ExecuteOperationResponse struct{}

func (h *ExecuteOperationHandler) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	ctx = context.WithValue(ctx, "authorization", request.Headers["authorization"])

	var req ExecuteOperationRequest

	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return nil, err
	}

	err = h.app.Commands.ExecuteOperation.Execute(ctx, &commands.ExecuteOperationCommand{
		OperationID: req.OperationID,
	})

	if err != nil {
		return nil, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func StartExecuteOperationHandler(app *application.Application) {
	handler := AuthHandler{app}
	lambda.Start(handler.Handle)
}
