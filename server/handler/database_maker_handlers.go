package handler

import (
	"DigiShop/database/migrations"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DatabaseMakerHandler(c *gin.Context) {
	migrations.MakeDatabaseTables()
	c.String(http.StatusOK, "Create All Tables in Database!")
}
