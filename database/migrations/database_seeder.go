package migrations

import (
	"DigiShop/database"
	"DigiShop/database/seeders"
)

func ExecuteSeeders() {
	// seed categories model
	seeders.SeedCategoriesModel(database.DB)

	// seed subcategories model
	seeders.SeedSubcategoriesModel(database.DB)

	// seed products model
	seeders.SeedProductsModel(database.DB)

	// seed reviews model
	seeders.SeedReviewsModel(database.DB)

	// seed users model
	seeders.SeedUsersModel(database.DB)

	// seed roles permissions model
	seeders.SeedRolesPermissionsModel(database.DB)

	// seed users roles model
	seeders.SeedUsersRolesModel(database.DB)
}
