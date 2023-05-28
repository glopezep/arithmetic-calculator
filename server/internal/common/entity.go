package common

import "github.com/google/uuid"

type Entity interface {
	GetID() string
}

type EntityBase struct {
	ID uuid.UUID
}

func (e EntityBase) GetID() uuid.UUID {
	return e.ID
}
