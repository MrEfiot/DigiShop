package main

import (
	"DigiShop/database"
	"DigiShop/server/routes"
)

func main() {
	database.SetupDB()
	routes.App()
}
