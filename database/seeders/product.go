package seeders

import (
	"DigiShop/database/models"
	"DigiShop/tools"
	"gorm.io/gorm"
)

func SeedProductModel(db *gorm.DB) {
	modelData := []models.Product{
		{SubcategoryID: 1, Name: "Macbook Air M1", Image: "", TechnicalSpec: "", InstallationGuide: "", VideoLink: ""},
		{SubcategoryID: 1, Name: "ASUS ROG Zephyrus G14", Image: "", TechnicalSpec: "", InstallationGuide: "", VideoLink: ""},
		{SubcategoryID: 2, Name: "iPhone 14 Pro Max", Image: "", TechnicalSpec: "", InstallationGuide: "", VideoLink: ""},
	}

	tx := db.Begin()

	for _, data := range modelData {
		err := db.Create(&data).Error
		if err != nil {
			tx.Rollback()
			tools.CheckError(err, "failed seed product models!")
		}
	}

	tx.Commit()
}
