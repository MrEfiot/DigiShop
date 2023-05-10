package controller

import (
	"DigiShop/database/models"
	"DigiShop/tools"
	"gorm.io/gorm"
)

func MakeCategoryTable(db *gorm.DB) {
	err := db.AutoMigrate(&models.Category{})
	tools.CheckError(err, "failed to migrate category table")
}
