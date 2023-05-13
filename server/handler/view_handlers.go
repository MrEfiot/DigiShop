package handler

import (
	"DigiShop/database"
	"DigiShop/database/models"
	"DigiShop/tools"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ViewMainHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.gohtml", gin.H{
		"title":   "DigiShop",
		"welcome": "Welcome to DigiShop",
	})
}

func ViewProductUploadHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "product_upload.gohtml", gin.H{
		"title": "Upload Product",
	})
}

func ViewDashboardHandler(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	tools.CheckError(err, "failed to get jwt token in view dashboard handler")

	claims := &Claims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("DigiShop_KEY"), nil
	})
	tools.CheckError(err, "failed to jwt parse with claims in view dashboard handler")

	userID := claims.UserID

	var user models.User

	err = database.DB.Preload("RolePermissions").First(&user, userID).Error
	tools.CheckError(err, "failed to get user info in view dashboard handler")

	c.HTML(http.StatusOK, "dashboard.gohtml", gin.H{
		"title":  "Dashboard",
		"name":   user.Name,
		"family": user.Family,
		"email":  user.Email,
		"role":   user.RolePermissions[0].Role,
	})
}

func ViewLoginHandler(c *gin.Context) {
	errorMessage := c.Query("error")
	if errorMessage == "wrong-password" {
		c.HTML(http.StatusOK, "login.gohtml", gin.H{
			"error": "your password is wrong!",
		})
		return
	}

	token, err := c.Cookie("token")
	tools.CheckError(err, "failed to get jwt token in view login handler")

	if token != "" {
		c.Redirect(http.StatusSeeOther, "/dashboard")
		return
	}

	c.HTML(http.StatusOK, "login.gohtml", gin.H{
		"title": "Login",
	})
}
