package models

import (
	"DigiShop/tools"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID                uint   `gorm:"primaryKey"`
	SubcategoryID     uint   `gorm:"index"`
	Name              string `gorm:"unique;not null"`
	Image             string
	TechnicalSpec     string
	InstallationGuide string
	Reviews           []Review `gorm:"foreignKey:ProductID"`
	VideoLink         string
	CreatedAt         time.Time
	UpdatedAt         *time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

func MakeProductTable(db *gorm.DB) {
	err := db.AutoMigrate(&Product{})
	tools.CheckError(err, "failed to migrate product table!")
}
