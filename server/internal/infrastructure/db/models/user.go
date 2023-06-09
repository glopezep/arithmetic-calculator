package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID
	Email    string
	Password string
	Balance  int64
	Records  []Record
}
