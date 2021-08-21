package security

import (
	"github.com/cbuelvasc/cinema-backend/config"
	"github.com/cbuelvasc/cinema-backend/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var whiteListPaths = []string{
	"/favicon.ico",
	"/api",
	"/api/*",
	"/api/cinema/v1/signin",
	"/api/cinema/v1/signup",
}

// change default error message
func init() {
	middleware.ErrJWTMissing.Code = 401
	middleware.ErrJWTMissing.Message = "Unauthorized"
}

func WebSecurityConfig(e *echo.Echo) {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte(config.JWTSecret),
		Skipper:    skipAuth,
	}
	e.Use(middleware.JWTWithConfig(config))
}

func skipAuth(e echo.Context) bool {
	// Skip authentication for and signup login requests
	for _, path := range whiteListPaths {
		if path == e.Path() {
			return true
		}
	}
	return false
}
