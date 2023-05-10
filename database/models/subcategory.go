package models

import (
	"gorm.io/gorm"
	"time"
)

type Subcategory struct {
	ID          uint   `gorm:"primaryKey"`
	CategoryID  uint   `gorm:"index"`
	Name        string `gorm:"unique;not null"`
	Description string
	Image       string
	Products    []Product `gorm:"foreignKey:SubcategoryID"`
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
