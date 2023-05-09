package migrations

import (
	"DigiShop/database"
	"DigiShop/database/models"
)

func MakeDatabaseTables() {
	// create category table in database
	models.MakeCategoryTable(database.DB)

	// create subcategory table in database
	models.MakeSubcategoryTable(database.DB)
}
