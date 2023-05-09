package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MainHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!\n")
}
