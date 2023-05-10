package controller

import (
	"DigiShop/database/models"
	"DigiShop/tools"
	"gorm.io/gorm"
)

func MakeUserRoleTable(db *gorm.DB) {
	err := db.AutoMigrate(&models.UserRole{})
	tools.CheckError(err, "failed to migrate user role table")
}
