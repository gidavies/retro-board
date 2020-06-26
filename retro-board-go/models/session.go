package models

import (
	"github.com/jinzhu/gorm"
)

type Session struct {
	gorm.Model
	Url       string `gorm:"not null"`
	Name      string `gorm:"not null"`
	CreatedBy User   `gorm:"not null"`
	// createdAt *time.Time
	// updatedAt *time.Time
	// Birthday     *time.Time
	// Email        string  `gorm:"type:varchar(100);unique_index"`
	// Role         string  `gorm:"size:255"` // set field size to 255
	// MemberNumber *string `gorm:"unique;not null"` // set member number to unique and not null
	// Num          int     `gorm:"AUTO_INCREMENT"` // set num to auto incrementable
	// Address      string  `gorm:"index:addr"` // create index with name `addr` for address
	// IgnoreMe     int     `gorm:"-"` // ignore this field
}
