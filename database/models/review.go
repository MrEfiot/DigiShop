package models

import (
	"DigiShop/tools"
	"gorm.io/gorm"
	"time"
)

type Review struct {
	ID        uint `gorm:"primaryKey"`
	ProductID uint `gorm:"index"`
	UserName  string
	Comment   string
	Rating    int
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func MakeReviewTable(db *gorm.DB) {
	err := db.AutoMigrate(&Review{})
	tools.CheckError(err, "failed to migrate review table!")
}
