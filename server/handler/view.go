package handler

import (
	"DigiShop/tools"
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

func ViewLoginHandler(c *gin.Context) {
	email := c.Query("email")
	errorMessage := c.Query("error")
	if errorMessage == "wrong-password" {
		c.HTML(http.StatusOK, "login.gohtml", gin.H{
			"error": "Your password is wrong!",
			"email": email,
		})
		return
	}

	if errorMessage == "email-not-found" {
		c.HTML(http.StatusOK, "login.gohtml", gin.H{
			"error": "Email not found!",
			"email": email,
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
