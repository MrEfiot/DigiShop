package routes

import (
	"DigiShop/config"
	"DigiShop/middleware"
	"DigiShop/server/handler"
	"DigiShop/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

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
	serverConfig := config.LoadServerConfig(tools.ConfigAddress).SC

	// show server information
	fmt.Println("App is running...")
	infoApp := fmt.Sprintf("Visit App: http://%s:%s\n", serverConfig.Host, serverConfig.Port)
	fmt.Println(infoApp)

	// start engine (server)
	address := fmt.Sprintf("%s:%s", serverConfig.Host, serverConfig.Port)
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
	accessLevelConfig := config.LoadAccessLevelConfig(tools.ConfigAddress).AL

	// main route
	mainRoutes(router)

	// database routes
	apiRoutes(router)

	// database maker routes
	databaseMakerRoutes(router)

	// admin routes
	adminRoutes(router)

	// admin routes
	AdminRoutes(router, accessLevelConfig)
}

func mainRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("views/*")
	router.GET("/", handler.ViewMainHandler)
	router.GET("/dashboard", middleware.AccessControlMiddleware("owner", "super_admin", "admin", "user"), handler.ViewDashboardHandler)
	router.GET("/login", handler.ViewLoginHandler)
	router.POST("/auth", handler.AuthHandler)
	router.GET("/logout", handler.LogoutHandler)
}

func adminRoutes(router *gin.Engine) {
	r := router.Group("/admin/")
	r.GET("/product_upload", middleware.AccessControlMiddleware("owner", "super_admin", "admin"), handler.ViewProductUploadHandler)
	r.POST("/product_upload_handler", middleware.AccessControlMiddleware("owner", "super_admin", "admin"), handler.ProductUploadHandler)
}

func apiRoutes(router *gin.Engine) {
	r := router.Group("/api/")
	r.GET("/categories", handler.CategoryHandler)
	r.GET("/category/:categoryID/subcategories", handler.SubcategoryHandler)
	r.GET("/subcategory/:subcategoryID/products", handler.ProductHandler)
	r.GET("/product/:productID/reviews", handler.ReviewHandler)

}

func databaseMakerRoutes(router *gin.Engine) {
	r := router.Group("/db/")
	r.GET("/maker/tables", handler.DatabaseMakerHandler)
	r.GET("/maker/seeders", handler.DatabaseSeederHandler)
}
