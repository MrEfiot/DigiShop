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

func DatabaseSeederHandler(c *gin.Context) {
	migrations.ExecuteSeeders()
	c.String(http.StatusOK, "The Primary Data were Created!")
}
