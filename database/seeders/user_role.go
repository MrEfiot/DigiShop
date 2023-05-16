package seeders

import (
	"DigiShop/database/models"
	"DigiShop/tools"
	"gorm.io/gorm"
)

func SeedUsersRolesModel(db *gorm.DB) {
	modelData := []models.UserRole{
		{UserID: 1, RolePermissionID: 1},
		{UserID: 2, RolePermissionID: 2},
		{UserID: 3, RolePermissionID: 3},
		{UserID: 4, RolePermissionID: 4},
	}

	tx := db.Begin()

	for _, data := range modelData {
		err := tx.Create(&data).Error
		if err != nil {
			tx.Rollback()
			tools.CheckError(err, "failed seed user_role model")
		}
	}

	tx.Commit()
}
