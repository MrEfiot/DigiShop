package seeders

import (
	"DigiShop/database/models"
	"DigiShop/tools"
	"gorm.io/gorm"
)

func SeedRolePermissionsModel(db *gorm.DB) {
	modelData := []models.RolePermission{
		{Role: "owner", Permission: "read"},
		{Role: "owner", Permission: "write"},
		{Role: "owner", Permission: "create"},
		{Role: "owner", Permission: "update"},
		{Role: "owner", Permission: "delete"},
		{Role: "super_admin", Permission: "read"},
		{Role: "super_admin", Permission: "write"},
		{Role: "super_admin", Permission: "create"},
		{Role: "super_admin", Permission: "update"},
		{Role: "admin", Permission: "read"},
		{Role: "admin", Permission: "write"},
		{Role: "user", Permission: "read"},
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
