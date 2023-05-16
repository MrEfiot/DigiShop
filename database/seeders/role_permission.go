package seeders

import (
	"DigiShop/database/models"
	"DigiShop/tools"
	"gorm.io/gorm"
)

func SeedRolePermissionsModel(db *gorm.DB) {
	modelData := []models.RolePermission{
		{Role: "owner", Permission: "owner"},
		{Role: "super_admin", Permission: "super admin"},
		{Role: "admin", Permission: "admin"},
		{Role: "user", Permission: "user"},
	}

	tx := db.Begin()

	for _, data := range modelData {
		err := tx.Create(&data).Error
		if err != nil {
			tx.Rollback()
			tools.CheckError(err, "failed seed role_permission model")
		}
	}

	tx.Commit()
}
