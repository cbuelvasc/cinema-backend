package controller

import (
	"github.com/cbuelvasc/cinema-backend/model"
	"github.com/cbuelvasc/cinema-backend/repository"
	"github.com/cbuelvasc/cinema-backend/util"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type CountryControllerInterface interface {
	GetAllCountries(c echo.Context) error
	GetCountry(c echo.Context) error
	SaveCountry(c echo.Context) error
	DeleteCountry(c echo.Context) error
}

type CountryController struct {
	countryRepository repository.CountryRepository
}

func NewCountryController(countryRepository repository.CountryRepository) *CountryController {
	return &CountryController{
		countryRepository: countryRepository,
	}
}

// GetAllCountries godoc
// @Summary Get all countries
// @Description Get all country items
// @Tags countries
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(xml, json)
// @Param page query int false "page" minimum(1)
// @Param limit query int false "size" minimum(1)
// @Param countryId query string true "countryId"
// @Success 200 {array} model.Country
// @Failure 500 {object} handler.APIError
// @Router /countries [get]
// @Security ApiKeyAuth
func (countryController *CountryController) GetAllCountries(c echo.Context) error {
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64)

	pagedCountry, err := countryController.countryRepository.GetAllCountries(c.Request().Context(), page, limit)
	if err != nil {
		return err
	}
	return util.Negotiate(c, http.StatusOK, pagedCountry)
}

// GetCountry godoc
// @Summary Get a country
// @Description Get a country item
// @Tags countries
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "Country ID"
// @Success 200 {object} model.Country
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /countries/{id} [get]
// @Security ApiKeyAuth
func (countryController *CountryController) GetCountry(c echo.Context) error {
	id := c.Param("id")
	if id == "me" {
		id = util.GetUserIdFromToken(c)
	}

	country, err := countryController.countryRepository.GetCountryById(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return util.Negotiate(c, http.StatusOK, country)
}

// SaveCountry godoc
// @Summary Create a country
// @Description Create a new country item
// @Tags countries
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param country body model.CountryInput true "New country"
// @Success 200 {object} model.Country
// @Failure 400 {object} handler.APIError
// @Failure 409 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /countries [post]
// @Security ApiKeyAuth
func (countryController *CountryController) SaveCountry(c echo.Context) error {
	payload := new(model.CountryInput)
	if err := util.BindAndValidate(c, payload); err != nil {
		return err
	}

	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()
	country := &model.Country{CountryInput: payload}

	createdCountry, err := countryController.countryRepository.SaveCountry(c.Request().Context(), country)
	if err != nil {
		return err
	}

	return util.Negotiate(c, http.StatusCreated, createdCountry)
}

// UpdateCountry godoc
// @Summary Update a user
// @Description Update a user item
// @Tags users
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "Country ID"
// @Param user body model.CountryInput true "Country Info"
// @Success 200 {object} model.Country
// @Failure 400 {object} handler.APIError
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /users/{id} [put]
// @Security ApiKeyAuth
func (countryController *CountryController) UpdateCountry(c echo.Context) error {
	id := c.Param("id")

	payload := new(model.CountryInput)

	payload.UpdatedAt = time.Now()
	if err := util.BindAndValidate(c, payload); err != nil {
		return err
	}

	user, err := countryController.countryRepository.UpdateCountry(c.Request().Context(), id, &model.Country{CountryInput: payload})
	if err != nil {
		return err
	}
	return util.Negotiate(c, http.StatusOK, user)
}

// DeleteCountry godoc
// @Summary Delete a country
// @Description Delete a new country item
// @Tags countries
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "Country ID"
// @Success 204 {object} model.Country
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /countries/{id} [delete]
// @Security ApiKeyAuth
func (countryController *CountryController) DeleteCountry(c echo.Context) error {
	id := c.Param("id")

	e := countryController.countryRepository.DeleteCountry(c.Request().Context(), id)
	if e != nil {
		return e
	}
	return c.NoContent(http.StatusNoContent)
}
