package controller

import (
	"DigiShop/database/models"
	"DigiShop/tools"
	"gorm.io/gorm"
)

func MakeRolePermissionTable(db *gorm.DB) {
	err := db.AutoMigrate(&models.RolePermission{})
	tools.CheckError(err, "failed to migrate role permission table")
}
