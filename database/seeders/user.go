package seeders

import (
	"DigiShop/database/models"
	"DigiShop/tools"
	"gorm.io/gorm"
)

func SeedUsersModel(db *gorm.DB) {
	modelData := []models.User{
		{Name: "Tomas", Family: "Shelby", Age: 27, NationalCode: "", PhoneNumber: "", Email: ""},
		{Name: "Artur", Family: "Shelby", Age: 32, NationalCode: "", PhoneNumber: "", Email: ""},
		{Name: "Aida", Family: "Shelby", Age: 25, NationalCode: "", PhoneNumber: "", Email: ""},
		{Name: "John", Family: "Shelby", Age: 23, NationalCode: "", PhoneNumber: "", Email: ""},
	}

	tx := db.Begin()

	for _, data := range modelData {
		err := db.Create(&data).Error
		if err != nil {
			tx.Rollback()
			tools.CheckError(err, "failed seed user models")
		}
	}

	tx.Commit()
}
