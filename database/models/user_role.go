package models

import (
	"gorm.io/gorm"
	"time"
)

type UserRole struct {
	UserID           uint
	RolePermissionID uint
	CreatedAt        time.Time
	UpdatedAt        *time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
