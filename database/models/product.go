package models

import (
	"gorm.io/gorm"
	"mime/multipart"
	"time"
)

type Product struct {
	ID                uint   `gorm:"primaryKey"`
	SubcategoryID     uint   `gorm:"index"`
	Name              string `gorm:"not null"`
	Image             string
	TechnicalSpec     string
	InstallationGuide string
	Reviews           []Review `gorm:"foreignKey:ProductID"`
	VideoLink         string
	CreatedAt         time.Time
	UpdatedAt         *time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

type ProductUpload struct {
	Name              string                `form:"product_name" binding:"required"`
	SubcategoryID     uint                  `form:"subcategory_id" binding:"required"`
	Image             *multipart.FileHeader `form:"image" binding:"required"`
	TechnicalSpec     *multipart.FileHeader `form:"technical_spec" binding:"required"`
	InstallationGuide *multipart.FileHeader `form:"installation_guide" binding:"required"`
	VideoLink         string                `form:"video_link" binding:"omitempty"`
}
