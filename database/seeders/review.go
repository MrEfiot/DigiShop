package seeders

import (
	"DigiShop/database/models"
	"DigiShop/tools"
	"gorm.io/gorm"
)

func SeedReviewModel(db *gorm.DB) {
	modelData := []models.Review{
		{ProductID: 1, UserName: "Tomas", Comment: "That's Perfect!", Rating: 35},
		{ProductID: 1, UserName: "Aida", Comment: "Very good:)", Rating: 120},
		{ProductID: 2, UserName: "Artur", Comment: "Great for gaming.", Rating: 7},
	}

	tx := db.Begin()

	for _, data := range modelData {
		err := db.Create(&data).Error
		if err != nil {
			tx.Rollback()
			tools.CheckError(err, "failed seed review models!")
		}
	}

	tx.Commit()
}
