package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Operation struct {
	gorm.Model
	ID      uuid.UUID
	Type    string
	Cost    int64
	Records []Record
}
