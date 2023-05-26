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

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

func CreateUserHandler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	app, err := application.NewApplication()
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize app")
	}

	var req CreateUserRequest

	err = json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal json")
	}

	err = app.Commands.CreateUser.Execute(ctx, &commands.CreateUserCommand{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create user")
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func StartCreateUserHandler() {
	lambda.Start(CreateUserHandler)
}
