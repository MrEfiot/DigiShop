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

func ViewLoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.gohtml", gin.H{
		"title": "Login",
	})
}
