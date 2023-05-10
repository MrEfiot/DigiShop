package routes

import (
	"DigiShop/server/handler"
	"DigiShop/tools"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

type Server struct {
	SC ServerConfig `json:"server"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func App() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// start all routes in application
	makeRoutes(router)

	// load server config
	config := loadConfig(tools.ConfigAddress)

	// show server information
	fmt.Println("App is running...")
	infoApp := fmt.Sprintf("Visit App: http://%s:%s\n", config.SC.Host, config.SC.Port)
	fmt.Println(infoApp)

	// start engine (server)
	address := fmt.Sprintf("%s:%s", config.SC.Host, config.SC.Port)
	err := router.Run(address)
	tools.CheckError(err, "failed to start the server!")
}

func makeRoutes(router *gin.Engine) {
	// main route
	mainRoutes(router)

	// database routes
	databaseRoutes(router)

	// database maker routes
	databaseMakerRoutes(router)
}

func mainRoutes(router *gin.Engine) {
	router.GET("/", handler.MainHandler)
}

func databaseRoutes(router *gin.Engine) {
	router.GET("categories", handler.CategoryHandler)
	router.GET("category/:categoryID/subcategories", handler.SubcategoryHandler)
	router.GET("subcategory/:subcategoryID/products", handler.ProductHandler)
	router.GET("products/:productID/reviews", handler.ReviewHandler)
}

func databaseMakerRoutes(router *gin.Engine) {
	router.GET("/maker/tables", handler.DatabaseMakerHandler)
	router.GET("/maker/seeders", handler.DatabaseSeederHandler)
}

func loadConfig(filename string) *Server {
	file, err := os.ReadFile(filename)
	tools.CheckError(err, "open database config!")

	var SC Server
	err = json.Unmarshal(file, &SC)
	tools.CheckError(err, "json unmarshal in load database config!")

	return &SC
}
