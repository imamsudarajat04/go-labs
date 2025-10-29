package migrations

import (
	"gorm.io/gorm"
)

func MigrateTestModel(db *gorm.DB) error {
	type TestModel struct {
		
	}

	return db.AutoMigrate(&TestModel{})
}
