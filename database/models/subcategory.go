package models

import (
	"DigiShop/tools"
	"gorm.io/gorm"
)

type Subcategory struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"unique;not null"`
	Description string
	CategoryID  uint `gorm:"index"`
	Category    Category
}

func MakeSubcategoryTable(db *gorm.DB) {
	err := db.AutoMigrate(&Subcategory{})
	tools.CheckError(err, "failed to migrate subcategory table!")
}
