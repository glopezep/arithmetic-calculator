package common

import (
	"context"
)

type EventHandler func(ctx context.Context, event Event) error

type Event interface {
	String() string
}
