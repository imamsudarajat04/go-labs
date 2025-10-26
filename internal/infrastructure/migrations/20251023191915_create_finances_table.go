package migrations

import (
	"gorm.io/gorm"
)

func MigrateFinance(db *gorm.DB) error {
	type Finance struct {
		
	}

	return db.AutoMigrate(&Finance{})
}
