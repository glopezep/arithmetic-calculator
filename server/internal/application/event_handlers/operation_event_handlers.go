package eventhandlers

import (
	"context"
	"math"
	"strconv"

	"github.com/glopezep/arithmetic-calculator/internal/common"
	"github.com/glopezep/arithmetic-calculator/internal/domain/entities"
	"github.com/glopezep/arithmetic-calculator/internal/domain/events"
	"github.com/glopezep/arithmetic-calculator/internal/domain/repositories"
	eventdispatcher "github.com/glopezep/arithmetic-calculator/internal/infrastructure/event_dispatcher"
	randomstring "github.com/glopezep/arithmetic-calculator/internal/infrastructure/services/random_string"
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
	operation    repositories.OperationRepository
	record       repositories.RecordRepository
	randomString randomstring.RandomStringService
}

var _ OperationDomainEventHandlers = (*OperationHandlers)(nil)

func NewOperationHandlers(
	operation repositories.OperationRepository,
	record repositories.RecordRepository,
	randomString randomstring.RandomStringService,
) *OperationHandlers {
	return &OperationHandlers{
		operation,
		record,
		randomString,
	}
}

func (h OperationHandlers) OnOperationDivided(ctx context.Context, event common.Event) error {
	e := event.(*events.OperationDivided)
	str := strconv.Itoa(int(e.FirstValue) / int(e.SecondValue))

	r, err := entities.NewRecord(e.ID, e.UserID, e.Cost, e.UserBalance, str)
	if err != nil {
		return err
	}

	err = h.record.Save(ctx, r)
	if err != nil {
		return err
	}

	return nil
}

func (h OperationHandlers) OnOperationMultiplied(ctx context.Context, event common.Event) error {
	e := event.(*events.OperationMultiplied)
	str := strconv.Itoa(int(e.FirstValue) * int(e.SecondValue))

	r, err := entities.NewRecord(e.ID, e.UserID, e.Cost, e.UserBalance, str)
	if err != nil {
		return err
	}

	err = h.record.Save(ctx, r)
	if err != nil {
		return err
	}

	return nil
}

func (h OperationHandlers) OnOperationRandomStringGenerated(ctx context.Context, event common.Event) error {
	e := event.(*events.OperationRandomStringGenerated)

	str, err := h.randomString.Generate()
	if err != nil {
		return err
	}

	r, err := entities.NewRecord(e.ID, e.UserID, e.Cost, e.UserBalance, str)
	if err != nil {
		return err
	}

	h.record.Save(ctx, r)

	return nil
}
func (h OperationHandlers) OnOperationSquareRooted(ctx context.Context, event common.Event) error {
	e := event.(*events.OperationSquareRooted)
	str := strconv.Itoa(int(math.Sqrt(float64(e.FirstValue))))

	r, err := entities.NewRecord(e.ID, e.UserID, e.Cost, e.UserBalance, str)
	if err != nil {
		return err
	}

	err = h.record.Save(ctx, r)
	if err != nil {
		return err
	}

	return nil
}

func (h OperationHandlers) OnOperationSubtracted(ctx context.Context, event common.Event) error {
	e := event.(*events.OperationSubtracted)
	str := strconv.Itoa(int(e.FirstValue) - int(e.SecondValue))

	r, err := entities.NewRecord(e.ID, e.UserID, e.Cost, e.UserBalance, str)
	if err != nil {
		return err
	}

	err = h.record.Save(ctx, r)
	if err != nil {
		return err
	}

	return nil
}

func (h OperationHandlers) OnOperationSummed(ctx context.Context, event common.Event) error {
	e := event.(*events.OperationSummed)
	str := strconv.Itoa(int(e.FirstValue) + int(e.SecondValue))

	r, err := entities.NewRecord(e.ID, e.UserID, e.Cost, e.UserBalance, str)
	if err != nil {
		return err
	}

	err = h.record.Save(ctx, r)
	if err != nil {
		return err
	}

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
