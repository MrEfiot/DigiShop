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
}
