package controller

import (
	"DigiShop/database"
	"DigiShop/database/models"
	"DigiShop/tools"
	"gorm.io/gorm"
)

func MakeProductTable(db *gorm.DB) {
	err := db.AutoMigrate(&models.Product{})
	tools.CheckError(err, "failed to migrate product table!")
}

func InsertProductTable(product *models.Product) {
	err := database.DB.Create(product).Error
	tools.CheckError(err, "failed to insert product")
}
