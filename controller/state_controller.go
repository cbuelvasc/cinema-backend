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

type StateControllerInterface interface {
	GetAllStates(c echo.Context) error
	GetState(c echo.Context) error
	SaveState(c echo.Context) error
	DeleteState(c echo.Context) error
}

type StateController struct {
	stateRepository   repository.StateRepository
	countryRepository repository.CountryRepository
}

func NewStateController(stateRepository repository.StateRepository, countryRepository repository.CountryRepository) *StateController {
	return &StateController{
		stateRepository:   stateRepository,
		countryRepository: countryRepository,
	}
}

// GetAllStates godoc
// @Summary Get all states
// @Description Get all states items
// @Tags states
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(xml, json)
// @Param page query int false "page" minimum(1)
// @Param limit query int false "size" minimum(1)
// @Param countryId query string true "countryId"
// @Success 200 {array} model.State
// @Failure 500 {object} handler.APIError
// @Router /states [get]
// @Security ApiKeyAuth
func (stateController *StateController) GetAllStates(c echo.Context) error {
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64)
	countryId := c.QueryParam("countryId")
	if len(countryId) > 0 {
		_, err := stateController.countryRepository.GetCountryById(c.Request().Context(), countryId)
		if err != nil {
			return err
		}
	}

	pagedState, err := stateController.stateRepository.GetAllStates(c.Request().Context(), page, limit, countryId)
	if err != nil {
		return err
	}
	return util.Negotiate(c, http.StatusOK, pagedState)
}

// GetState godoc
// @Summary Get a states
// @Description Get a states item
// @Tags states
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "State ID"
// @Success 200 {object} model.State
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /states/{id} [get]
// @Security ApiKeyAuth
func (stateController *StateController) GetState(c echo.Context) error {
	id := c.Param("id")
	if id == "me" {
		id = util.GetUserIdFromToken(c)
	}

	states, err := stateController.stateRepository.GetStateById(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return util.Negotiate(c, http.StatusOK, states)
}

// SaveState godoc
// @Summary Create a states
// @Description Create a new states item
// @Tags states
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param states body model.StateInput true "New states"
// @Success 200 {object} model.State
// @Failure 400 {object} handler.APIError
// @Failure 409 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /states [post]
// @Security ApiKeyAuth
func (stateController *StateController) SaveState(c echo.Context) error {
	payload := new(model.StateInput)
	if err := util.BindAndValidate(c, payload); err != nil {
		return err
	}

	_, err := stateController.countryRepository.GetCountryById(c.Request().Context(), payload.CountryId)
	if err != nil {
		return err
	}

	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()
	states := &model.State{StateInput: payload}

	createdState, err := stateController.stateRepository.SaveState(c.Request().Context(), states)
	if err != nil {
		return err
	}

	return util.Negotiate(c, http.StatusCreated, createdState)
}

// UpdateState godoc
// @Summary Update a state
// @Description Update a state item
// @Tags states
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "State ID"
// @Param state body model.StateInput true "State Info"
// @Success 200 {object} model.State
// @Failure 400 {object} handler.APIError
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /states/{id} [put]
// @Security ApiKeyAuth
func (stateController *StateController) UpdateState(c echo.Context) error {
	id := c.Param("id")

	payload := new(model.StateInput)

	if err := util.BindAndValidate(c, payload); err != nil {
		return err
	}

	user, err := stateController.stateRepository.UpdateState(c.Request().Context(), id, &model.State{StateInput: payload})
	if err != nil {
		return err
	}
	return util.Negotiate(c, http.StatusOK, user)
}

// DeleteState godoc
// @Summary Delete a states
// @Description Delete a new states item
// @Tags states
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "State ID"
// @Success 204 {object} model.State
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /states/{id} [delete]
// @Security ApiKeyAuth
func (stateController *StateController) DeleteState(c echo.Context) error {
	id := c.Param("id")

	e := stateController.stateRepository.DeleteState(c.Request().Context(), id)
	if e != nil {
		return e
	}
	return c.NoContent(http.StatusNoContent)
}
