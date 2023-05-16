package handler

import (
	"DigiShop/database"
	"DigiShop/database/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProductHandler(c *gin.Context) {
	var products []models.Product
	err := database.DB.Preload("Reviews").Find(&products).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func AddProductHandler(c *gin.Context) {
	var product models.Product
	err := c.ShouldBindJSON(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = database.DB.Create(&product).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added successfully"})
}

func EditProductHandler(c *gin.Context) {
	productID := c.Param("id")

	var product models.Product

	err := database.DB.Where("id = ?", productID).First(&product).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var updateProduct models.Product
	err = c.ShouldBindJSON(&updateProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.Name = updateProduct.Name
	product.Image = updateProduct.Image
	product.TechnicalSpec = updateProduct.TechnicalSpec
	product.InstallationGuide = updateProduct.InstallationGuide
	product.VideoLink = updateProduct.VideoLink

	err = database.DB.Save(&product).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func DeleteProductHandler(c *gin.Context) {
	productID := c.Param("id")

	err := database.DB.Where("id = ?", productID).Delete(&models.Product{}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
