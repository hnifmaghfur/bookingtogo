package domain

import (
	"time"

	"gorm.io/gorm"
)

type Nationality struct {
	gorm.Model
	ID        uint           `gorm:"primaryKey;column:id" json:"id"`
	Code      string         `gorm:"size:20;not null;unique;column:code" json:"code"` // Contoh: ID, US, JP
	Country   string         `gorm:"size:50;not null;column:country" json:"country"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at" json:"deleted_at,omitempty"`
}
