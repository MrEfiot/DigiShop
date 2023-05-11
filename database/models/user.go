package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID              uint   `gorm:"primaryKey"`
	Name            string `gorm:"not null"`
	Family          string `gorm:"not null"`
	Age             int
	NationalCode    string
	PhoneNumber     string
	Email           string
	RolePermissions []RolePermission `gorm:"many2many:user_roles;"`
	CreatedAt       time.Time
	UpdatedAt       *time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
