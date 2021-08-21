package main

import (
	"fmt"
	"log"

	"github.com/cbuelvasc/cinema-backend/config"
	"github.com/cbuelvasc/cinema-backend/controller"
	"github.com/cbuelvasc/cinema-backend/handler"
	"github.com/cbuelvasc/cinema-backend/repository"
	"github.com/cbuelvasc/cinema-backend/routes"
	"github.com/cbuelvasc/cinema-backend/security"
	"github.com/cbuelvasc/cinema-backend/util"
	"github.com/labstack/echo/v4"
)

var userController *controller.UserController
var tweetController *controller.TweetController
var movieController *controller.MovieController
var countryController *controller.CountryController
var stateController *controller.StateController
var cityController *controller.CityController

// @title Cinema REST API
// @description Provides access to the core features of Cinema REST API
// @version 1.0
// @termsOfService http://swagger.io/terms/
// license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/cinema/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	e := echo.New()

	e.HTTPErrorHandler = handler.ErrorHandler
	e.Validator = util.NewValidationUtil()
	config.CORSConfig(e)
	security.WebSecurityConfig(e)

	security.WebSecurityConfig(e)

	routes.GetUserApiRoutes(e, userController)
	routes.GetTweetApiRoutes(e, tweetController)
	routes.GetMovieApiRoutes(e, movieController)
	routes.GetCountryApiRoutes(e, countryController)
	routes.GetStateApiRoutes(e, stateController)
	routes.GetCityApiRoutes(e, cityController)
	routes.GetSwaggerRoutes(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.ServerPort)))
}

func init() {
	mongoConnection, errorMongoConn := config.MongoConnection()

	if errorMongoConn != nil {
		log.Println("Error when connect mongo : ", errorMongoConn.Error())
	}
	userRepository := repository.NewUserRepository(mongoConnection)
	authValidator := security.NewAuthValidator(userRepository)
	userController = controller.NewUserController(userRepository, authValidator)

	tweetRepository := repository.NewTeewtRepository(mongoConnection)
	tweetController = controller.NewTweetController(tweetRepository, userRepository)

	movieRepository := repository.NewMovieRepository(mongoConnection)
	movieController = controller.NewMovieController(movieRepository, userRepository)

	countryRepository := repository.NewCountryRepository(mongoConnection)
	countryController = controller.NewCountryController(countryRepository)

	stateRepository := repository.NewStateRepository(mongoConnection)
	stateController = controller.NewStateController(stateRepository, countryRepository)

	cityRepository := repository.NewCityRepository(mongoConnection)
	cityController = controller.NewCityController(cityRepository, stateRepository)
}
