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
	err := database.DB.Preload("Subcategories").First(&category, categoryID).Error
	tools.CheckError(err, "failed preload subcategories")

	subcategories := category.Subcategories
	c.JSON(http.StatusOK, subcategories)
}

func ProductHandler(c *gin.Context) {
	subcategoryID := c.Param("subcategoryID")

	var subcategory models.Subcategory
	err := database.DB.Preload("Products.Reviews").First(&subcategory, subcategoryID).Error
	tools.CheckError(err, "failed preload products")

	products := subcategory.Products
	c.JSON(http.StatusOK, products)
}

func ReviewHandler(c *gin.Context) {
	productID := c.Param("productID")

	var product models.Product
	err := database.DB.Preload("Reviews").First(&product, productID).Error
	tools.CheckError(err, "failed preload reviews")

	reviews := product.Reviews
	c.JSON(http.StatusOK, reviews)
}
