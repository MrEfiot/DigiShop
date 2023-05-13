package handler

import (
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
	c.HTML(http.StatusOK, "dashboard.gohtml", gin.H{
		"title": "Dashboard",
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
	c.HTML(http.StatusOK, "login.gohtml", gin.H{
		"title": "Login",
	})
}
