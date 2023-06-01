package db

import (
	"fmt"

	"gorm.io/gorm"
)

func Order(sortBy, orderBy string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if sortBy == "" {
			sortBy = "created_at"
		}

		if orderBy == "" || orderBy != "asc" && orderBy != "desc" {
			orderBy = "desc"
		}

		return db.Order(fmt.Sprintf("%v %v", sortBy, orderBy))
	}
}
