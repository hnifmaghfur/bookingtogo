package domain

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	CstID         uint           `gorm:"primaryKey;column:cst_id" json:"cst_id"`
	NationalityID uint           `gorm:"not null;column:nationality_id" json:"nationality_id"`
	CstName       string         `gorm:"size:50;not null;column:cst_name" json:"cst_name"`
	CstDob        string         `gorm:"type:date;not null;column:cst_dob" json:"cst_dob"`
	CstPhoneNum   string         `gorm:"size:20;not null;column:cst_phonenum" json:"cst_phonenum"`
	CstEmail      string         `gorm:"size:50;unique;not null;column:cst_email" json:"cst_email"`
	FamilyList    []FamilyList   `gorm:"foreignKey:CstID" json:"family_list"`
	Nationality   *Nationality   `gorm:"foreignKey:NationalityID" json:"nationality"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index;column:deleted_at" json:"deleted_at,omitempty"`
}
