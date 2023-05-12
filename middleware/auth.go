package middleware

import (
	"DigiShop/database/controller"
	"DigiShop/server/handler"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if shouldSkipAuth(c.Request.URL.Path) {
			c.Next()
			return
		}

		tokenString, err := c.Cookie("token")
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected singin mothod")
			}

			return []byte("DigiShop_KEY"), nil
		})

		if err != nil || !token.Valid {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		c.Next()
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if shouldSkipAuth(c.Request.URL.Path) {
			c.Next()
			return
		}

		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}
		claims := &handler.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("DigiShop_KEY"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}

		userID := claims.UserID

		permissions := controller.GetUserPermissions(userID)
		accessPermission := false

		for i := range permissions {
			if permissions[i] == "write" {
				accessPermission = true
			}
		}

		if !accessPermission {
			c.Redirect(http.StatusTemporaryRedirect, "/")
		}

		c.Next()
	}
}

func shouldSkipAuth(path string) bool {
	paths := []string{"/product_upload", "/categories"}
	active := true

	for i := range paths {
		if paths[i] == path {
			active = false
		}
	}

	return active
}
