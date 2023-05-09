package handler

import (
	"DigiShop/database"
	"DigiShop/database/models"
	"DigiShop/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CategoryHandler(c *gin.Context) {
	var categories []models.Category
	err := database.DB.Find(&categories).Error
	tools.CheckError(err, "failed to retrieve categories")
	c.JSON(http.StatusOK, categories)
}

func SubcategoryHandler(c *gin.Context) {
	categoryID := c.Param("categoryID")

	var category models.Category
	err := database.DB.Preload("Subcategories.Category").First(&category, categoryID).Error
	tools.CheckError(err, "failed preload subcategories")

	subcategories := category.Subcategories
	c.JSON(http.StatusOK, subcategories)
}
