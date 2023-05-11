package middleware

import (
	"github.com/gin-gonic/gin"
)

func CheckPermission(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

//func getUserPermission(userID uint) models.RolePermissions {
//	var user models.Users
//	var role models.RolePermissions
//
//
//}
