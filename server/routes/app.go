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
	setMiddleware(router)

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

func setMiddleware(router *gin.Engine) {
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())
	router.Use(middleware.AuthMiddleware())
	router.Use(middleware.AuthRequired())
}

func makeRoutes(router *gin.Engine) {
	// main route
	mainRoutes(router)

	// database routes
	apiRoutes(router)

	// database maker routes
	databaseMakerRoutes(router)

	// admin routes
	adminRoutes(router)
}

func mainRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("views/*")
	router.GET("/", handler.ViewMainHandler)
	router.GET("/dashboard", middleware.PageAccessMiddleware("owner", "super_admin", "admin", "user"), handler.ViewDashboardHandler)
	router.GET("/login", handler.ViewLoginHandler)
	router.POST("/auth", handler.AuthHandler)
}

func adminRoutes(router *gin.Engine) {
	r := router.Group("/admin/")
	r.GET("/product_upload", middleware.PageAccessMiddleware("owner", "super_admin", "admin"), handler.ViewProductUploadHandler)
	r.POST("/product_upload_handler", middleware.PageAccessMiddleware("owner", "super_admin", "admin"), handler.ProductUploadHandler)
}

func apiRoutes(router *gin.Engine) {
	r := router.Group("/api/")
	r.GET("/categories", handler.CategoryHandler)
	r.GET("/category/:categoryID/subcategories", handler.SubcategoryHandler)
	r.GET("/subcategory/:subcategoryID/products", handler.ProductHandler)
	r.GET("/product/:productID/reviews", handler.ReviewHandler)

	// user routes
	userRoutes(r)
}

func databaseMakerRoutes(router *gin.Engine) {
	r := router.Group("/db/")
	r.GET("/maker/tables", handler.DatabaseMakerHandler)
	r.GET("/maker/seeders", handler.DatabaseSeederHandler)
}

func userRoutes(router *gin.RouterGroup) {
	router.GET("/users", middleware.PageAccessMiddleware("owner", "super_admin"), handler.UserHandler)
}
