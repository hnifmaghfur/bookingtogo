package domain

import (
	"time"

	"gorm.io/gorm"
)

type FamilyList struct {
	FlID       uint           `gorm:"primaryKey;column:fl_id" json:"fl_id"`
	CstID      uint           `gorm:"not null;column:cst_id" json:"cst_id"` // Foreign key to customers table
	FlName     string         `gorm:"size:50;not null;column:fl_name" json:"fl_name"`
	FlRelation string         `gorm:"size:50;not null;column:fl_relation" json:"fl_relation"` // e.g., "Wife", "Son", "Daughter"
	FlDob      string         `gorm:"size:50;not null;column:fl_dob" json:"fl_dob"`
	CreatedAt  time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index;column:deleted_at" json:"deleted_at,omitempty"`
}
