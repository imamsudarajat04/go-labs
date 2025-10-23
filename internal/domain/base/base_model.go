package base

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (bm *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	bm.ID = uuid.New().String()
	return
}
