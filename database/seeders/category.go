package seeders

import (
	"DigiShop/database/models"
	"DigiShop/tools"
	"gorm.io/gorm"
)

func SeedCategoriesModel(db *gorm.DB) {
	modelData := []models.Category{
		{Name: "Electronic", Description: "Electronic Products"},
		{Name: "Clothing", Description: "Clothing Products"},
		{Name: "Sport", Description: "Sport Products"},
	}

	tx := db.Begin()

	for _, data := range modelData {
		err := tx.Create(&data).Error
		if err != nil {
			tx.Rollback()
			tools.CheckError(err, "failed seed categories model!")
		}
	}

	tx.Commit()
}
