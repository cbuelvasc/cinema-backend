package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/cbuelvasc/cinema-backend/model"
	"github.com/cbuelvasc/cinema-backend/repository"
	"github.com/cbuelvasc/cinema-backend/util"
	"github.com/labstack/echo/v4"
)

type CityControllerInterface interface {
	GetAllCities(c echo.Context) error
	GetCity(c echo.Context) error
	SaveCity(c echo.Context) error
	DeleteCity(c echo.Context) error
}

type CityController struct {
	cityRepository  repository.CityRepository
	stateRepository repository.StateRepository
}

func NewCityController(cityRepository repository.CityRepository, stateRepository repository.StateRepository) *CityController {
	return &CityController{
		cityRepository:  cityRepository,
		stateRepository: stateRepository,
	}
}

// GetAllCities godoc
// @Summary Get all cities
// @Description Get all cities items
// @Tags cities
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(xml, json)
// @Param page query int false "page" minimum(1)
// @Param limit query int false "size" minimum(1)
// @Param stateId query string true "stateId"
// @Success 200 {array} model.City
// @Failure 500 {object} handler.APIError
// @Router /cities [get]
// @Security ApiKeyAuth
func (cityController *CityController) GetAllCities(c echo.Context) error {
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64)
	stateId := c.QueryParam("stateId")
	if len(stateId) > 0 {
		_, err := cityController.stateRepository.GetStateById(c.Request().Context(), stateId)
		if err != nil {
			return err
		}
	}

	pagedCity, err := cityController.cityRepository.GetAllCities(c.Request().Context(), page, limit, stateId)
	if err != nil {
		return err
	}
	return util.Negotiate(c, http.StatusOK, pagedCity)
}

// GetCity godoc
// @Summary Get a City
// @Description Get a city item
// @Tags cities
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "City ID"
// @Success 200 {object} model.City
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /cities/{id} [get]
// @Security ApiKeyAuth
func (cityController *CityController) GetCity(c echo.Context) error {
	id := c.Param("id")
	if id == "me" {
		id = util.GetUserIdFromToken(c)
	}

	city, err := cityController.cityRepository.GetCityById(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return util.Negotiate(c, http.StatusOK, city)
}

// SaveCity godoc
// @Summary Create a City
// @Description Create a new city item
// @Tags cities
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param City body model.CityInput true "New City"
// @Success 200 {object} model.City
// @Failure 400 {object} handler.APIError
// @Failure 409 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /cities [post]
// @Security ApiKeyAuth
func (cityController *CityController) SaveCity(c echo.Context) error {
	payload := new(model.CityInput)
	if err := util.BindAndValidate(c, payload); err != nil {
		return err
	}

	_, err := cityController.stateRepository.GetStateById(c.Request().Context(), payload.StateId)
	if err != nil {
		return err
	}

	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()
	city := &model.City{CityInput: payload}

	createdCity, err := cityController.cityRepository.SaveCity(c.Request().Context(), city)
	if err != nil {
		return err
	}

	return util.Negotiate(c, http.StatusCreated, createdCity)
}

// UpdateCity godoc
// @Summary Update a city
// @Description Update a city item
// @Tags cities
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "City ID"
// @Param city body model.CityInput true "City Info"
// @Success 200 {object} model.City
// @Failure 400 {object} handler.APIError
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /cities/{id} [put]
// @Security ApiKeyAuth
func (cityController *CityController) UpdateCity(c echo.Context) error {
	id := c.Param("id")

	payload := new(model.CityInput)

	if err := util.BindAndValidate(c, payload); err != nil {
		return err
	}

	payload.UpdatedAt = time.Now()
	user, err := cityController.cityRepository.UpdateCity(c.Request().Context(), id, &model.City{CityInput: payload})
	if err != nil {
		return err
	}
	return util.Negotiate(c, http.StatusOK, user)
}

// DeleteCity godoc
// @Summary Delete a city
// @Description Delete a city item
// @Tags cities
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "City ID"
// @Success 204 {object} model.City
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /cities/{id} [delete]
// @Security ApiKeyAuth
func (cityController *CityController) DeleteCity(c echo.Context) error {
	id := c.Param("id")
	stateId := c.Param("stateId")

	_, err := cityController.stateRepository.GetStateById(c.Request().Context(), stateId)
	if err != nil {
		return err
	}

	e := cityController.cityRepository.DeleteCity(c.Request().Context(), id, stateId)
	if e != nil {
		return e
	}
	return c.NoContent(http.StatusNoContent)
}
