package migrations

import (
	"DigiShop/database"
	"DigiShop/database/models"
)

func MakeDatabaseTables() {
	// create category table
	models.MakeCategoryTable(database.DB)

	// create subcategory table
	models.MakeSubcategoryTable(database.DB)
	//
	//// create product and review table
	models.MakeProductTable(database.DB)
	//
	//// create review table
	models.MakeReviewTable(database.DB)
}
