package controller

import (
	"DigiShop/database/models"
	"DigiShop/tools"
	"gorm.io/gorm"
)

func MakeUserTable(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	tools.CheckError(err, "failed to migrate user table")
}
