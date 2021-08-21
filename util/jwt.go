package util

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cbuelvasc/cinema-backend/config"
	"github.com/cbuelvasc/cinema-backend/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func GenerateJwtToken(user *model.User) (string, error) {
	expTimeMs, _ := strconv.Atoi(config.JWTExpirationMs)
	exp := time.Now().Add(time.Millisecond * time.Duration(expTimeMs)).Unix()

	// Set custom claims
	claims := &model.JwtCustomClaims{
		user.ID.Hex(),
		user.Email,
		user.Name,
		user.Lastname,
		fmt.Sprintf("%s", user.BirthDate),
		user.Biography,
		user.Location,
		user.WebSite,
		jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response
	jwt, err := token.SignedString([]byte(config.JWTSecret))
	return jwt, err
}

func GetUserIdFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtCustomClaims)
	return claims.ID
}
