package routes

import (
	"github.com/cbuelvasc/cinema-backend/controller"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func GetSwaggerRoutes(e *echo.Echo) {
	e.GET("/api", controller.RedirectIndexPage)
	e.GET("/api/*", echoSwagger.WrapHandler)
}
