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

type StateControllerInterface interface {
	GetAllStates(c echo.Context) error
	GetState(c echo.Context) error
	SaveState(c echo.Context) error
	DeleteState(c echo.Context) error
}

type StateController struct {
	stateRepository repository.StateRepository
}

func NewStateController(stateRepository repository.StateRepository) *StateController {
	return &StateController{
		stateRepository: stateRepository,
	}
}

// GetAllStates godoc
// @Summary Get all countries
// @Description Get all states items
// @Tags countries
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(xml, json)
// @Param page query int false "page" minimum(1)
// @Param limit query int false "size" minimum(1)
// @Param statesId query string true "statesId"
// @Success 200 {array} model.State
// @Failure 500 {object} handler.APIError
// @Router /countries [get]
// @Security ApiKeyAuth
func (stateController *StateController) GetAllStates(c echo.Context) error {
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64)

	pagedState, err := stateController.stateRepository.GetAllStates(c.Request().Context(), page, limit)
	if err != nil {
		return err
	}
	return util.Negotiate(c, http.StatusOK, pagedState)
}

// GetState godoc
// @Summary Get a states
// @Description Get a states item
// @Tags countries
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "State ID"
// @Success 200 {object} model.State
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /countries/{id} [get]
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
// @Tags countries
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param states body model.StateInput true "New states"
// @Success 200 {object} model.State
// @Failure 400 {object} handler.APIError
// @Failure 409 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /countries [post]
// @Security ApiKeyAuth
func (stateController *StateController) SaveState(c echo.Context) error {
	payload := new(model.StateInput)
	if err := util.BindAndValidate(c, payload); err != nil {
		return err
	}

	payload.CreatedAt = time.Now()
	states := &model.State{StateInput: payload}

	createdState, err := stateController.stateRepository.SaveState(c.Request().Context(), states)
	if err != nil {
		return err
	}

	return util.Negotiate(c, http.StatusCreated, createdState)
}

// UpdateState godoc
// @Summary Update a user
// @Description Update a user item
// @Tags users
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "State ID"
// @Param user body model.StateInput true "State Info"
// @Success 200 {object} model.State
// @Failure 400 {object} handler.APIError
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /users/{id} [put]
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
// @Tags countries
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "State ID"
// @Success 204 {object} model.State
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /countries/{id} [delete]
// @Security ApiKeyAuth
func (stateController *StateController) DeleteState(c echo.Context) error {
	id := c.Param("id")

	e := stateController.stateRepository.DeleteState(c.Request().Context(), id)
	if e != nil {
		return e
	}
	return c.NoContent(http.StatusNoContent)
}
