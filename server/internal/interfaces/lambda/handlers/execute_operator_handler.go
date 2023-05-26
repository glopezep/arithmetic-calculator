package handlers

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/glopezep/arithmetic-calculator/internal/application"
	"github.com/glopezep/arithmetic-calculator/internal/application/commands"
	"github.com/google/uuid"
	"github.com/stackus/errors"
)

type ExecuteOperatorEvent struct {
	ID uuid.UUID `json:"id"`
}

type ExecuteOperatorResponse struct {
	ID uuid.UUID `json:"id"`
}

func ExecuteOperatorHandler(ctx context.Context, event ExecuteOperatorEvent) (*ExecuteOperatorResponse, error) {
	app, err := application.NewApplication()
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize application")
	}

	cmd := &commands.DeleteOperationCommand{}

	if err = app.Commands.DeleteRecord.Execute(ctx, cmd); err != nil {
		return nil, errors.Wrap(err, "failed to delete the record")
	}

	return &ExecuteOperatorResponse{ID: event.ID}, nil
}

func StartExecuteOperatorHandler() {
	lambda.Start(ExecuteOperatorHandler)
}
