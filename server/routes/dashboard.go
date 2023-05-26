package routes

import (
	"DigiShop/config"
	"DigiShop/middleware"
	"DigiShop/server/handler"
	"github.com/gin-gonic/gin"
)

func DashboardRoutes(router *gin.Engine, al config.AccessLevel) {
	dashboard := router.Group("/dashboard/")
	dashboard.GET("/", middleware.AccessControlMiddleware(al.LevelOne, al.LevelTwo, al.LevelThree, al.LevelFour), handler.ViewDashboardHandler)
	dashboard.GET("/tables", middleware.AccessControlMiddleware(al.LevelOne, al.LevelTwo), handler.ViewDashboardTablesHandler)
}
