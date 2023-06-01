package handlers

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/glopezep/arithmetic-calculator/internal/application"
	"github.com/glopezep/arithmetic-calculator/internal/application/queries"
	"github.com/glopezep/arithmetic-calculator/internal/interfaces/lambda/helpers"
)

type AuthHandler struct {
	app *application.Application
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

func (h *AuthHandler) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var req AuthRequest

	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return nil, err
	}

	token, err := h.app.Queries.AuthenticateUser.Execute(ctx, &queries.AuthenticateUserQuery{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		lambdaError := helpers.LambdaError{
			Code:        "403",
			Message:     "invalid credentials",
			Description: err,
		}

		return &events.APIGatewayProxyResponse{
			StatusCode: 403,
			Body:       string(lambdaError.MarshalJSON()),
		}, err
	}

	bytes, err := json.Marshal(AuthResponse{
		Token: token,
	})

	if err != nil {
		lambdaError := helpers.LambdaError{
			Code:        "500",
			Message:     "failed to marshal response",
			Description: err,
		}

		return &events.APIGatewayProxyResponse{
			StatusCode: 403,
			Body:       string(lambdaError.MarshalJSON()),
		}, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(bytes),
	}, nil
}

func StartAuthHandler(app *application.Application) {
	handler := AuthHandler{app}

	lambda.Start(handler.Handle)
}
