package models

import (
	"DigiShop/tools"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name              string `gorm:"unique;not null"`
	Image             string
	TechnicalSpec     string
	InstallationGuide string
	Reviews           []Review `gorm:"foreignKey:ProductID"`
	VideoLink         string
}

func MakeProductTable(db *gorm.DB) {
	err := db.AutoMigrate(&Product{})
	tools.CheckError(err, "failed to migrate product table!")
}
