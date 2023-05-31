package utils

import (
	"context"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

type Handler func(ctx context.Context, event events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error)

type ContextKey string

func HandleWithContext(handler Handler) Handler {
	return func(
		ctx context.Context,
		event events.APIGatewayProxyRequest,
	) (*events.APIGatewayProxyResponse, error) {

		auth := event.Headers["Authorization"]
		token := strings.Split(auth, " ")[1]
		newContext := context.WithValue(ctx, ContextKey("token"), token)

		return handler(newContext, event)
	}
}
