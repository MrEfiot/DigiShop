package controller

import (
	"DigiShop/database/models"
	"DigiShop/tools"
	"gorm.io/gorm"
)

func MakeReviewTable(db *gorm.DB) {
	err := db.AutoMigrate(&models.Review{})
	tools.CheckError(err, "failed to migrate review table!")
}
