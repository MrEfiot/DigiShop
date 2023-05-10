package seeders

import (
	"DigiShop/database/models"
	"DigiShop/tools"
	"gorm.io/gorm"
)

func SeedUsersRolesModel(db *gorm.DB) {
	modelData := []models.UserRole{
		{UserID: 1, RolePermissionID: 1},
		{UserID: 1, RolePermissionID: 2},
		{UserID: 1, RolePermissionID: 3},
		{UserID: 1, RolePermissionID: 4},
		{UserID: 1, RolePermissionID: 5},
		{UserID: 2, RolePermissionID: 6},
		{UserID: 2, RolePermissionID: 7},
		{UserID: 2, RolePermissionID: 8},
		{UserID: 2, RolePermissionID: 9},
		{UserID: 3, RolePermissionID: 10},
		{UserID: 3, RolePermissionID: 11},
		{UserID: 4, RolePermissionID: 12},
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
