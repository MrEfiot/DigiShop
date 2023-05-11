package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ViewMainHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.gohtml", gin.H{
		"welcome": "Welcome to DigiShop",
	})
}

func ViewProductUploadHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "upload_product.gohtml", gin.H{})
}
