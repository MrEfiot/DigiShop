package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ViewDashboardHandler(c *gin.Context) {
	claims := GetClaims(c)

	userID := claims.UserID

	user := GetUser(userID)

	role := GetUserRole(user)
	if role == "owner" || role == "super_admin" {

		c.HTML(http.StatusOK, "admin_dashboard.gohtml", gin.H{
			"title": "Admin | Dashboard",
		})
		return
	}
	c.HTML(http.StatusOK, "dashboard.gohtml", gin.H{
		"title":  "Dashboard",
		"name":   user.Name,
		"family": user.Family,
		"email":  user.Email,
		"role":   user.RolePermissions[0].Role,
	})
}

func ViewDashboardTablesHandler(c *gin.Context) {
	claims := GetClaims(c)

	userID := claims.UserID

	user := GetUser(userID)

	users := GetAllUser(c)
	c.HTML(http.StatusOK, "tables_dashboard.gohtml", gin.H{
		"title":  "Tables | Dashboard",
		"userID": user.ID,
		"users":  users,
	})
	return
}
