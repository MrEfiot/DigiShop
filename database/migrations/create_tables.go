package migrations

import (
	"DigiShop/database"
	"DigiShop/database/models"
)

func MakeDatabaseTables() {
	// create database and setup operation
	db := database.SetupDB()

	// create category table in database
	models.MakeCategoryTable(db)
}
