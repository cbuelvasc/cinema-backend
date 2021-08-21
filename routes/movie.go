package routes

import (
	"github.com/cbuelvasc/cinema-backend/controller"
	"github.com/cbuelvasc/cinema-backend/enums"
	"github.com/labstack/echo/v4"
)

func GetMovieApiRoutes(e *echo.Echo, movieController *controller.MovieController) {
	v1 := e.Group(enums.BasePath)
	{
		v1.POST(enums.CreateMovie, movieController.SaveMovie)
		v1.GET(enums.GetMovies, movieController.GetAllMovie)
		v1.GET(enums.GetMovieById, movieController.GetMovie)
		v1.PUT(enums.UpdateMovieById, movieController.UpdateMovie)
		v1.DELETE(enums.DeleteMovieById, movieController.DeleteMovie)

	}
}
