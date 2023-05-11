package handler

import (
	"DigiShop/database"
	"DigiShop/database/controller"
	"DigiShop/database/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func AuthHandler(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.String(http.StatusOK, "email not found!")
		return
	}

	success := controller.CompareHash(user.Password, password)
	if success {
		c.String(http.StatusOK, "your logged in :)")
	} else {
		c.String(http.StatusOK, "password is wrong!")
		return
	}
}
