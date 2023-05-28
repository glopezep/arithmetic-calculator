package common

import "github.com/google/uuid"

type Aggregate interface {
	Entity
	AddEvent(event Event)
	GetEvents() []Event
}

type AggregateBase struct {
	ID     uuid.UUID
	events []Event
}

func (a AggregateBase) GetID() uuid.UUID {
	return a.ID
}

func (a *AggregateBase) AddEvent(event Event) {
	a.events = append(a.events, event)
}

func (a AggregateBase) GetEvents() []Event {
	return a.events
}
