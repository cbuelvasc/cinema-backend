package routes

import (
	"github.com/cbuelvasc/cinema-backend/controller"
	"github.com/cbuelvasc/cinema-backend/enums"
	"github.com/labstack/echo/v4"
)

func GetCityApiRoutes(e *echo.Echo, cityController *controller.CityController) {
	v1 := e.Group(enums.BasePath)
	{
		v1.POST(enums.CreateCity, cityController.SaveCity)
		v1.GET(enums.GetCities, cityController.GetAllCities)
		v1.GET(enums.GetCityById, cityController.GetCity)
		v1.PUT(enums.UpdateCityById, cityController.UpdateCity)
		v1.DELETE(enums.DeleteCityById, cityController.DeleteCity)
	}
}
