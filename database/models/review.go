package models

import (
	"DigiShop/tools"
	"gorm.io/gorm"
	"time"
)

type Review struct {
	gorm.Model
	ProductID uint
	UserName  string
	Comment   string
	Rating    int
	Date      time.Time
}

func MakeReviewTable(db *gorm.DB) {
	err := db.AutoMigrate(&Review{})
	tools.CheckError(err, "failed to migrate review table!")
}
