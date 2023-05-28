package eventhandlers

import (
	"context"

	"github.com/glopezep/arithmetic-calculator/internal/common"
	"github.com/glopezep/arithmetic-calculator/internal/domain/events"
	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
	eventdispatcher "github.com/glopezep/arithmetic-calculator/internal/infrastructure/event_dispatcher"
)

type OperationDomainEventHandlers interface {
	OnOperationDivided(ctx context.Context, event common.Event) error
	OnOperationMultiplied(ctx context.Context, event common.Event) error
	OnOperationRandomStringGenerated(ctx context.Context, event common.Event) error
	OnOperationSquareRooted(ctx context.Context, event common.Event) error
	OnOperationSubtracted(ctx context.Context, event common.Event) error
	OnOperationSummed(ctx context.Context, event common.Event) error
}

type OperationHandlers struct {
	operationRepository repositories.OperationRepository
}

var _ OperationDomainEventHandlers = (*OperationHandlers)(nil)

func NewOperationHandlers(operationRepository repositories.OperationRepository) *OperationHandlers {
	return &OperationHandlers{
		operationRepository,
	}
}

func (h OperationHandlers) OnOperationDivided(ctx context.Context, event common.Event) error {
	return nil
}

func (h OperationHandlers) OnOperationMultiplied(ctx context.Context, event common.Event) error {
	return nil
}
func (h OperationHandlers) OnOperationRandomStringGenerated(ctx context.Context, event common.Event) error {
	return nil
}
func (h OperationHandlers) OnOperationSquareRooted(ctx context.Context, event common.Event) error {
	return nil
}
func (h OperationHandlers) OnOperationSubtracted(ctx context.Context, event common.Event) error {
	return nil
}

func (h OperationHandlers) OnOperationSummed(ctx context.Context, event common.Event) error {
	// orderCreated := event.(*domain.OrderCreated)
	// return h.notifications.NotifyOrderCreated(ctx, orderCreated.Order.ID, orderCreated.Order.CustomerID)
	return nil
}

func RegisterOperationHandlers(operationHandlers OperationDomainEventHandlers, domainSubscriber eventdispatcher.EventSubscriber) {
	domainSubscriber.Subscribe(&events.OperationDivided{}, operationHandlers.OnOperationDivided)
	domainSubscriber.Subscribe(&events.OperationMultiplied{}, operationHandlers.OnOperationMultiplied)
	domainSubscriber.Subscribe(&events.OperationRandomStringGenerated{}, operationHandlers.OnOperationRandomStringGenerated)
	domainSubscriber.Subscribe(&events.OperationSquareRooted{}, operationHandlers.OnOperationSquareRooted)
	domainSubscriber.Subscribe(&events.OperationSubtracted{}, operationHandlers.OnOperationSubtracted)
	domainSubscriber.Subscribe(&events.OperationSummed{}, operationHandlers.OnOperationSummed)
}
