package routes

import (
	"DigiShop/config"
	"DigiShop/middleware"
	"DigiShop/server/handler"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine, al config.AccessLevel) {
	admin := router.Group("/admin/")

	// products routes
	productRoutes(admin, al)

	// userRoutes
	userRoutes(admin, al)
}

func userRoutes(router *gin.RouterGroup, al config.AccessLevel) {
	router.GET("users", middleware.AccessControlMiddleware(al.LevelOne, al.LevelTwo), handler.GetUserHandler)
	router.POST("users/add", middleware.AccessControlMiddleware(al.LevelOne, al.LevelTwo), handler.AddUserHandler)
	router.PUT("users/edit/:id", middleware.AccessControlMiddleware(al.LevelOne), handler.EditUserHandler)
	router.DELETE("users/delete/:id", middleware.AccessControlMiddleware(al.LevelOne), handler.DeleteUserHandler)
}

func productRoutes(router *gin.RouterGroup, al config.AccessLevel) {
	router.GET("products", handler.GetProductHandler)
	router.POST("products/add", middleware.AccessControlMiddleware(al.LevelOne, al.LevelTwo), handler.AddProductHandler)
	router.PUT("products/edit/:id", middleware.AccessControlMiddleware(al.LevelOne, al.LevelTwo), handler.EditProductHandler)
	router.DELETE("products/delete/:id", middleware.AccessControlMiddleware(al.LevelOne, al.LevelTwo), handler.DeleteProductHandler)
}
