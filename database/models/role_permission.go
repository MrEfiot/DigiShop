package models

import (
	"gorm.io/gorm"
	"time"
)

type RolePermission struct {
	ID         uint   `gorm:"primaryKey"`
	Role       string `gorm:"uniqueIndex:idx_role_permission"`
	Permission string `gorm:"uniqueIndex:idx_role_permission"`
	Users      []User `gorm:"many2many:user_roles;"`
	CreatedAt  time.Time
	UpdatedAt  *time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
