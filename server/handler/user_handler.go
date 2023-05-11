package handler

import (
	"DigiShop/database"
	"DigiShop/database/models"
	"DigiShop/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserHandler(c *gin.Context) {
	var users []models.User
	err := database.DB.Preload("RolePermissions").Find(&users).Error
	tools.CheckError(err, "failed to preloading users")

	c.JSON(http.StatusOK, users)
}
