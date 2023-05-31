package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	ID                uuid.UUID
	OperationID       uuid.UUID `gorm:"primaryKey"`
	UserID            uuid.UUID `gorm:"primaryKey"`
	Amount            int64
	UserBalance       int64
	OperationResponse string
}
