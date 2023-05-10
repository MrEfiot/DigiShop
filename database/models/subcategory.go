package models

import (
	"DigiShop/tools"
	"gorm.io/gorm"
	"time"
)

type Subcategory struct {
	ID          uint   `gorm:"primaryKey"`
	CategoryID  uint   `gorm:"index"`
	Name        string `gorm:"unique;not null"`
	Description string
	Image       string
	Products    []Product `gorm:"foreignKey:SubcategoryID"`
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func MakeSubcategoryTable(db *gorm.DB) {
	err := db.AutoMigrate(&Subcategory{})
	tools.CheckError(err, "failed to migrate subcategory table!")
}
