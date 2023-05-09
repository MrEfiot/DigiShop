package models

import (
	"DigiShop/tools"
	"gorm.io/gorm"
)

type Category struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"unique;not null"`
	Description   string
	Subcategories []Subcategory `gorm:"foreignKey:CategoryID"`
}

func MakeCategoryTable(db *gorm.DB) {
	err := db.AutoMigrate(&Category{})
	tools.CheckError(err, "failed to migrate category table!")
}
