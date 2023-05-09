package handler

import (
	"DigiShop/database/migrations"
	"github.com/gin-gonic/gin"
)

func MainHandler(c *gin.Context) {
	c.String(200, "Hello, World!\n")
	migrations.MakeDatabaseTables()
	c.String(200, "Create All Tables in Database!")
}
