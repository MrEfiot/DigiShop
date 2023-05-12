package controller

import (
	"DigiShop/database"
	"DigiShop/database/models"
	"DigiShop/tools"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func MakeUserTable(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	tools.CheckError(err, "failed to migrate user table")
}

func GetUserPermissions(UserID uint) []string {
	var user models.User
	err := database.DB.Preload("RolePermissions").First(&user, UserID).Error
	tools.CheckError(err, "failed to get user permissions")

	var permissions []string

	for _, rolePermission := range user.RolePermissions {
		permissions = append(permissions, rolePermission.Permission)
	}

	return permissions
}

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	tools.CheckError(err, "failed to hash password")

	return string(hashedPassword)
}

func CompareHash(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
