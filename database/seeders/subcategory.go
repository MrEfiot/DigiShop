package seeders

import (
	"DigiShop/database/models"
	"DigiShop/tools"
	"gorm.io/gorm"
)

func SeedSubcategoriesModel(db *gorm.DB) {
	modelData := []models.Subcategory{
		{Name: "Laptop", Description: "Laptop Products", CategoryID: 1, Image: ""},
		{Name: "Mobile", Description: "Mobile Products", CategoryID: 1, Image: ""},
		{Name: "Manly", Description: "Manly Products", CategoryID: 2, Image: ""},
		{Name: "Feminine", Description: "Feminine Products", CategoryID: 2, Image: ""},
		{Name: "Childish", Description: "Childish Products", CategoryID: 2, Image: ""},
		{Name: "Manly and Feminine", Description: "Manly and Feminine Products", CategoryID: 2, Image: ""},
		{Name: "Sportswear", Description: "Sportswear Products", CategoryID: 3, Image: ""},
	}

	tx := db.Begin()

	for _, data := range modelData {
		err := tx.Create(&data).Error
		if err != nil {
			tx.Rollback()
			tools.CheckError(err, "failed seed subcategories model!")
		}
	}

	tx.Commit()
}
