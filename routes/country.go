package routes

import (
	"github.com/cbuelvasc/cinema-backend/controller"
	"github.com/cbuelvasc/cinema-backend/enums"
	"github.com/labstack/echo/v4"
)

func GetCountryApiRoutes(e *echo.Echo, movieController *controller.CountryController) {
	v1 := e.Group(enums.BasePath)
	{
		v1.POST(enums.CreateCountry, movieController.SaveCountry)
		v1.GET(enums.GetCountries, movieController.GetAllCountries)
		v1.GET(enums.GetCountryById, movieController.GetCountry)
		v1.PUT(enums.UpdateCountryById, movieController.UpdateCountry)
		v1.DELETE(enums.DeleteCountryById, movieController.DeleteCountry)
	}
}
