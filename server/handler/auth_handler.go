package handler

import (
	"DigiShop/database"
	"DigiShop/database/controller"
	"DigiShop/database/models"
	"DigiShop/tools"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func AuthHandler(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.String(http.StatusOK, "email not found!")
		return
	}

	success := controller.CompareHash(user.Password, password)
	if success {
		token, expireTime := createToken(user.ID)
		setCookieToken(token, expireTime, c)
		c.String(http.StatusOK, "Logged in successfully")
	} else {
		c.String(http.StatusOK, "password is wrong!")
		return
	}

}

func createToken(userID uint) (string, time.Time) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: strconv.Itoa(int(userID)),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("DigiShop_KEY"))
	tools.CheckError(err, "failed to generate jwt token")

	return tokenString, expirationTime
}

func setCookieToken(tokenString string, expireTime time.Time, c *gin.Context) {
	name := "token"
	token := tokenString
	expire := int(expireTime.Unix())
	path := "/"
	domain := ""

	c.SetCookie(name, token, expire, path, domain, true, true)
}
