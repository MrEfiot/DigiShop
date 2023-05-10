package models

import (
	"gorm.io/gorm"
	"time"
)

type Review struct {
	ID        uint `gorm:"primaryKey"`
	ProductID uint `gorm:"index"`
	UserName  string
	Comment   string
	Rating    int
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
