package helpers

import (
	"context"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/services/token"
	"github.com/google/uuid"
)

type Handler func(ctx context.Context, event events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error)

type ContextKey string

type Context struct {
	UserID uuid.UUID
	Token  string
}

func HandleWithContext(handler Handler) Handler {
	tokenService := token.NewJwtTokenService()

	return func(
		ctx context.Context,
		event events.APIGatewayProxyRequest,
	) (*events.APIGatewayProxyResponse, error) {
		if event.Headers["Authorization"] == "" {
			return handler(ctx, event)
		}

		auth := event.Headers["Authorization"]
		tokenString := strings.Split(auth, " ")[1]
		claims, _ := tokenService.Verify(tokenString)
		newContext := context.WithValue(ctx, ContextKey("context"), Context{
			UserID: uuid.MustParse(claims.RegisteredClaims.Subject),
			Token:  tokenString,
		})

		return handler(newContext, event)
	}
}
