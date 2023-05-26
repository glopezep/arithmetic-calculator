package handlers

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/glopezep/arithmetic-calculator/internal/application"
	"github.com/glopezep/arithmetic-calculator/internal/application/queries"
	"github.com/stackus/errors"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

func AuthHandler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	app, err := application.NewApplication()
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize app")
	}

	var req AuthRequest

	err = json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal json")
	}

	token, err := app.Queries.AuthenticateUser.Execute(ctx, &queries.AuthenticateUserQuery{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return nil, errors.Wrap(err, "failed to authenticate user")
	}

	bytes, err := json.Marshal(AuthResponse{
		Token: token,
	})

	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal json")
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(bytes),
	}, nil
}

func StartAuthHandler() {
	lambda.Start(AuthHandler)
}
