package handlers

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/glopezep/arithmetic-calculator/internal/application"
	"github.com/glopezep/arithmetic-calculator/internal/application/commands"
)

type CreateUserHandler struct {
	app *application.Application
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

func (h *CreateUserHandler) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var req CreateUserRequest

	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return &events.APIGatewayProxyResponse{}, err
	}

	err = h.app.Commands.CreateUser.Execute(ctx, &commands.CreateUserCommand{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return &events.APIGatewayProxyResponse{}, err

	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func StartCreateUserHandler(app *application.Application) {
	handler := CreateUserHandler{app}

	lambda.Start(handler.Handle)
}
