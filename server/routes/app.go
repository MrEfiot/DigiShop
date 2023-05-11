package routes

import (
	"DigiShop/middleware"
	"DigiShop/server/handler"
	"DigiShop/tools"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

type Server struct {
	SC ServerConfig `json:"server"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func loadConfig(filename string) *Server {
	file, err := os.ReadFile(filename)
	tools.CheckError(err, "open database config!")

	var SC Server
	err = json.Unmarshal(file, &SC)
	tools.CheckError(err, "json unmarshal in load database config!")

	return &SC
}

func setupLogOutput() {
	f, err := os.Create("gin.log")
	tools.CheckError(err, "failed to create output logger")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func App() {
	setupLogOutput()

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// set static file
	router.Static("/static", "./static")

	// set middleware
	router.Use(gin.Recovery(), middleware.Logger())

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

	// user routes
	userRoutes(router)
}

func mainRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("views/*")
	router.GET("/", handler.ViewMainHandler)
	router.GET("product_upload", handler.ViewProductUploadHandler)
	router.POST("product_upload_handler", handler.ProductUploadHandler)
	router.GET("login", handler.ViewLoginHandler)
}

func databaseRoutes(router *gin.Engine) {
	router.GET("categories", handler.CategoryHandler)
	router.GET("category/:categoryID/subcategories", handler.SubcategoryHandler)
	router.GET("subcategory/:subcategoryID/products", handler.ProductHandler)
	router.GET("products/:productID/reviews", handler.ReviewHandler)
}

func databaseMakerRoutes(router *gin.Engine) {
	router.GET("maker/tables", handler.DatabaseMakerHandler)
	router.GET("maker/seeders", handler.DatabaseSeederHandler)
}

func userRoutes(router *gin.Engine) {
	router.GET("users", handler.UserHandler)
}
