package handlers

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/glopezep/arithmetic-calculator/internal/application"
	"github.com/glopezep/arithmetic-calculator/internal/application/commands"
	"github.com/glopezep/arithmetic-calculator/internal/interfaces/lambda/utils"
	"github.com/google/uuid"
)

type ExecuteOperationHandler struct {
	app *application.Application
}

type ExecuteOperationRequest struct {
	ID          string `json:"id"`
	FirstValue  int64  `json:"firstValue"`
	SecondValue int64  `json:"secondValue"`
}

type ExecuteOperationResponse struct{}

func (h *ExecuteOperationHandler) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var req ExecuteOperationRequest

	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return nil, err
	}

	err = h.app.Commands.ExecuteOperation.Execute(ctx, &commands.ExecuteOperationCommand{
		OperationID: uuid.MustParse(req.ID),
		FirstValue:  req.FirstValue,
		SecondValue: req.SecondValue,
	})

	if err != nil {
		return nil, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func StartExecuteOperationHandler(app *application.Application) {
	handler := ExecuteOperationHandler{app}

	lambda.Start(utils.HandleWithContext(handler.Handle))
}
