package seeders

import (
	"DigiShop/database/controller"
	"DigiShop/database/models"
	"DigiShop/tools"
	"gorm.io/gorm"
)

func SeedUsersModel(db *gorm.DB) {
	modelData := []models.User{
		{Name: "Tomas", Family: "Shelby", Age: 27, NationalCode: "0", PhoneNumber: "0", Email: "Mr.Efiot@gmail.com", Password: "12345678"},
		{Name: "Artur", Family: "Shelby", Age: 32, NationalCode: "1", PhoneNumber: "1", Email: "Mr.Efiot2@gmail.com", Password: "12345678"},
		{Name: "Aida", Family: "Shelby", Age: 25, NationalCode: "2", PhoneNumber: "2", Email: "Mr.Efiot3@gmail.com", Password: "12345678"},
		{Name: "John", Family: "Shelby", Age: 23, NationalCode: "3", PhoneNumber: "3", Email: "Mr.Efiot4@gmail.com", Password: "12345678"},
	}

	tx := db.Begin()

	for _, data := range modelData {
		hashPassword := controller.HashPassword(data.Password)
		data.Password = hashPassword
		err := db.Create(&data).Error
		if err != nil {
			tx.Rollback()
			tools.CheckError(err, "failed seed user models")
		}
	}

	tx.Commit()
}
