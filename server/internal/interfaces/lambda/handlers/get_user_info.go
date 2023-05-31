package handlers

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/glopezep/arithmetic-calculator/internal/application"
	"github.com/glopezep/arithmetic-calculator/internal/application/queries"
	"github.com/glopezep/arithmetic-calculator/internal/interfaces/lambda/utils"
	"github.com/google/uuid"
)

type GetUserInfoHandler struct {
	app *application.Application
}

type GetUserInfoResponse struct {
	ID      uuid.UUID `json:"id"`
	Email   string    `json:"email"`
	Status  string    `json:"status"`
	Balance int64     `json:"balance"`
}

func (h *GetUserInfoHandler) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	u, err := h.app.Queries.GetUserInfo.Execute(ctx, &queries.GetUserInfoQuery{})

	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(GetUserInfoResponse{
		ID:      u.ID,
		Email:   u.Email.String(),
		Status:  u.Status.String(),
		Balance: u.Balance,
	})

	if err != nil {
		return nil, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(bytes),
	}, nil
}

func StartGetUserInfoHandler(app *application.Application) {
	handler := GetUserInfoHandler{app}

	lambda.Start(utils.HandleWithContext(handler.Handle))
}
