package models

import (
	"DigiShop/tools"
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"unique;not null"`
	Description   string
	Subcategories []Subcategory `gorm:"foreignKey:CategoryID"`
	CreatedAt     time.Time
	UpdatedAt     *time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func MakeCategoryTable(db *gorm.DB) {
	err := db.AutoMigrate(&Category{})
	tools.CheckError(err, "failed to migrate category table!")
}
