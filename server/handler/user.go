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

func GetUserHandler(c *gin.Context) {
	var users []models.User
	err := database.DB.Preload("RolePermissions").Find(&users).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
	}

	c.JSON(http.StatusOK, users)
}

func AddUserHandler(c *gin.Context) {
	var userForm models.UserForm
	err := c.ShouldBind(&userForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := database.DB.Begin()

	checkExistingEmailAndPhone(tx, userForm, c)
	rolePermission := getRolePermission(tx, userForm, c)

	user := models.User{
		Name:            userForm.Name,
		Family:          userForm.Family,
		Age:             userForm.Age,
		NationalCode:    userForm.NationalCode,
		PhoneNumber:     userForm.PhoneNumber,
		Email:           userForm.Email,
		Password:        controller.HashPassword(userForm.Password),
		RolePermissions: []models.RolePermission{rolePermission},
	}

	err = tx.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user"})
		tx.Rollback()
		return
	}

	err = tx.Commit().Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User added successfully"})
}

func EditUserHandler(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	err := database.DB.Where("id = ?", userID).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var userForm models.UserForm
	err = c.ShouldBind(&userForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := database.DB.Begin()

	checkExistingEmailAndPhone(tx, userForm, c)

	user.Name = userForm.Name
	user.Family = userForm.Family
	user.Age = userForm.Age
	user.NationalCode = userForm.NationalCode
	user.PhoneNumber = userForm.PhoneNumber
	user.Email = userForm.Email

	err = database.DB.Save(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUserHandler(c *gin.Context) {
	userID := c.Param("id")

	err := database.DB.Where("id = ?", userID).Delete(&models.User{}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func checkExistingEmailAndPhone(tx *gorm.DB, userForm models.UserForm, c *gin.Context) {
	var existingUser models.User
	err := tx.Where("phone_number = ?", userForm.PhoneNumber).First(&existingUser).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Phone number already exists"})
		tx.Rollback()
		return
	}

	err = tx.Where("email = ?", userForm.Email).First(&existingUser).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Email already exists"})
		tx.Rollback()
		return
	}

}

func getRolePermission(tx *gorm.DB, userForm models.UserForm, c *gin.Context) models.RolePermission {
	var rolePermission models.RolePermission
	err := tx.Where("id = ?", userForm.RolePermissionID).First(&rolePermission).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "role permission not found"})
		tx.Rollback()
		return models.RolePermission{}
	}

	return rolePermission
}
