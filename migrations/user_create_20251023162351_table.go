package migrations

import (
	"gorm.io/gorm"
)

func MigrateUser(db *gorm.DB) error {
	type User struct {
		
	}

	return db.AutoMigrate(&User{})
}
