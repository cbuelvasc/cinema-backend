package controller

import (
	"github.com/cbuelvasc/cinema-backend/model"
	"github.com/cbuelvasc/cinema-backend/repository"
	"github.com/cbuelvasc/cinema-backend/util"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type MovieControllerInterface interface {
	GetAllMovie(c echo.Context) error
	GetMovie(c echo.Context) error
	SaveMovie(c echo.Context) error
	DeleteMovie(c echo.Context) error
}

type MovieController struct {
	movieRepository repository.MovieRepository
	userRepository  repository.UserRepository
}

func NewMovieController(movieRepository repository.MovieRepository, userRepository repository.UserRepository) *MovieController {
	return &MovieController{
		movieRepository: movieRepository,
		userRepository:  userRepository,
	}
}

// GetAllMovie godoc
// @Summary Get all movies
// @Description Get all movie items
// @Tags movies
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(xml, json)
// @Param page query int false "page" minimum(1)
// @Param limit query int false "size" minimum(1)
// @Param movieId query string true "movieId"
// @Success 200 {array} model.Movie
// @Failure 500 {object} handler.APIError
// @Router /movies [get]
// @Security ApiKeyAuth
func (movieController *MovieController) GetAllMovie(c echo.Context) error {
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64)

	pagedMovie, err := movieController.movieRepository.GetAllMovies(c.Request().Context(), page, limit)
	if err != nil {
		return err
	}
	return util.Negotiate(c, http.StatusOK, pagedMovie)
}

// GetMovie godoc
// @Summary Get a movie
// @Description Get a movie item
// @Tags movies
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "Movie ID"
// @Success 200 {object} model.Movie
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /movies/{id} [get]
// @Security ApiKeyAuth
func (movieController *MovieController) GetMovie(c echo.Context) error {
	id := c.Param("id")
	if id == "me" {
		id = util.GetUserIdFromToken(c)
	}

	movie, err := movieController.movieRepository.GetMovie(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return util.Negotiate(c, http.StatusOK, movie)
}

// SaveMovie godoc
// @Summary Create a movie
// @Description Create a new movie item
// @Tags movies
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param movie body model.MovieInput true "New movie"
// @Success 200 {object} model.Movie
// @Failure 400 {object} handler.APIError
// @Failure 409 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /movies [post]
// @Security ApiKeyAuth
func (movieController *MovieController) SaveMovie(c echo.Context) error {
	payload := new(model.MovieInput)
	if err := util.BindAndValidate(c, payload); err != nil {
		return err
	}

	movie := &model.Movie{MovieInput: payload}

	createdMovie, err := movieController.movieRepository.SaveMovie(c.Request().Context(), movie)
	if err != nil {
		return err
	}

	return util.Negotiate(c, http.StatusCreated, createdMovie)
}

// UpdateMovie godoc
// @Summary Update a user
// @Description Update a user item
// @Tags users
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "Movie ID"
// @Param user body model.MovieInput true "Movie Info"
// @Success 200 {object} model.Movie
// @Failure 400 {object} handler.APIError
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /users/{id} [put]
// @Security ApiKeyAuth
func (movieController *MovieController) UpdateMovie(c echo.Context) error {
	id := c.Param("id")

	payload := new(model.MovieInput)

	if err := util.BindAndValidate(c, payload); err != nil {
		return err
	}

	user, err := movieController.movieRepository.UpdateMovie(c.Request().Context(), id, &model.Movie{MovieInput: payload})
	if err != nil {
		return err
	}
	return util.Negotiate(c, http.StatusOK, user)
}

// DeleteMovie godoc
// @Summary Delete a movie
// @Description Delete a new movie item
// @Tags movies
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "Movie ID"
// @Success 204 {object} model.Movie
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /movies/{id} [delete]
// @Security ApiKeyAuth
func (movieController *MovieController) DeleteMovie(c echo.Context) error {
	id := c.Param("id")
	movieId := c.Param("movieId")

	e := movieController.movieRepository.DeleteMovie(c.Request().Context(), id, movieId)
	if e != nil {
		return e
	}
	return c.NoContent(http.StatusNoContent)
}
