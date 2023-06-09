package middleware

import (
	"DigiShop/database/controller"
	"DigiShop/server/handler"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AccessControlMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")

		if err != nil {
			path := c.Request.URL.Path
			if path == "/dashboard" {
				c.Redirect(http.StatusSeeOther, "/login")
			}
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

		expirationUnix := claims.ExpiresAt
		expirationTime := time.Unix(expirationUnix, 0)

		if time.Now().After(expirationTime) {
			c.SetCookie("token", "", -1, "", "", false, true)

			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.Abort()
			return
		}

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid tokens",
			})
			c.Abort()
			return
		}

		role := controller.GetUserRole(claims.UserID)
		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				c.Next()
				return
			}
		}

		c.Redirect(http.StatusTemporaryRedirect, "/")
		c.Abort()
	}
}
