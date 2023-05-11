package controller

import (
	"DigiShop/database/models"
	"DigiShop/tools"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func MakeUserTable(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	tools.CheckError(err, "failed to migrate user table")
}

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	tools.CheckError(err, "failed to hash password")

	return string(hashedPassword)
}
