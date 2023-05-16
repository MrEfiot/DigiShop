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
	PhoneNumber     string `gorm:"unique"`
	Email           string `gorm:"unique"`
	Password        string
	RolePermissions []RolePermission `gorm:"many2many:user_roles;"`
	CreatedAt       time.Time
	UpdatedAt       *time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

type UserForm struct {
	Name             string `form:"name" binding:"required"`
	Family           string `form:"family" binding:"required"`
	Age              int    `form:"age"`
	NationalCode     string `form:"national_code"`
	PhoneNumber      string `form:"phone_number" binding:"required_without=Email"`
	Email            string `form:"email" binding:"required_without=PhoneNumber"`
	Password         string `form:"password"`
	RolePermissionID uint   `form:"role_permission_id"`
}
