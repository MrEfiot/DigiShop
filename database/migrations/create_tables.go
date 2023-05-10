package migrations

import (
	"DigiShop/database"
	"DigiShop/database/controller"
)

func MakeDatabaseTables() {
	// create category table
	controller.MakeCategoryTable(database.DB)

	// create subcategory table
	controller.MakeSubcategoryTable(database.DB)

	// create product and review table
	controller.MakeProductTable(database.DB)

	// create review table
	controller.MakeReviewTable(database.DB)

	// create user table
	controller.MakeUserTable(database.DB)

	// create role_permission table
	controller.MakeRolePermissionTable(database.DB)

	// create user_role table
	controller.MakeUserRoleTable(database.DB)
}
