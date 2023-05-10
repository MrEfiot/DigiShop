package controller

import (
	"DigiShop/database/models"
	"DigiShop/tools"
	"gorm.io/gorm"
)

func MakeSubcategoryTable(db *gorm.DB) {
	err := db.AutoMigrate(&models.Subcategory{})
	tools.CheckError(err, "failed to migrate subcategory table!")
}
