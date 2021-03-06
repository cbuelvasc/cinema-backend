package controller

import (
	"net/http"
	"strconv"

	"github.com/cbuelvasc/cinema-backend/exception"
	"github.com/cbuelvasc/cinema-backend/model"
	"github.com/cbuelvasc/cinema-backend/repository"
	"github.com/cbuelvasc/cinema-backend/security"
	"github.com/cbuelvasc/cinema-backend/util"
	"github.com/labstack/echo/v4"
)

type UserControllerInterface interface {
	AuthenticateUser(c echo.Context) error
	SaveUser(c echo.Context) error
	GetAllUser(c echo.Context) error
	GetUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

type UserController struct {
	userRepository repository.UserRepository
	authValidator  *security.AuthValidator
}

func NewUserController(userRepository repository.UserRepository, authValidator *security.AuthValidator) *UserController {
	return &UserController{userRepository: userRepository, authValidator: authValidator}
}

// AuthenticateUser godoc
// @Summary Authenticate User
// @Description Authenticate a user item
// @Tags users
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(xml, json)
// @Param user body model.SignInInput true "SignIn"
// @Success 200 {array} model.User
// @Failure 400 {object} handler.APIError
// @Failure 401 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /signin [post]
func (userController *UserController) AuthenticateUser(c echo.Context) error {
	payload := new(model.SignInInput)
	if err := util.BindAndValidate(c, payload); err != nil {
		return err
	}

	user, valid := userController.authValidator.ValidateCredentials(c.Request().Context(), payload.Email, payload.Password)
	if !valid {
		return exception.UnauthorizedException()
	}

	jwt, err := util.GenerateJwtToken(user)
	if err != nil {
		return err
	}

	return util.Negotiate(c, http.StatusOK, model.Token{Token: jwt})
}

// SaveUser godoc
// @Summary Create a user
// @Description Create a new user item
// @Tags users
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param user body model.UserInput true "New User"
// @Success 200 {object} model.User
// @Failure 400 {object} handler.APIError
// @Failure 409 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /signup [post]
func (userController *UserController) SaveUser(c echo.Context) error {
	payload := new(model.UserInput)
	if err := util.BindAndValidate(c, payload); err != nil {
		return err
	}

	_, err := userController.userRepository.FindByEmail(c.Request().Context(), payload.Email)
	if err == nil {
		return exception.ConflictException("User", "email", payload.Email)
	}

	user := &model.User{UserInput: payload}

	//encrypt password
	err = beforeSave(user)
	if err != nil {
		return err
	}

	createdUser, err := userController.userRepository.SaveUser(c.Request().Context(), user)
	if err != nil {
		return err
	}

	return util.Negotiate(c, http.StatusCreated, createdUser)
}

// GetAllUser godoc
// @Summary Get all users
// @Description Get all user items
// @Tags users
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(xml, json)
// @Param page query int false "page" minimum(1)
// @Param limit query int false "size" minimum(1)
// @Success 200 {array} model.User
// @Failure 500 {object} handler.APIError
// @Router /users [get]
// @Security ApiKeyAuth
func (userController *UserController) GetAllUser(c echo.Context) error {
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64)

	pagedUser, _ := userController.userRepository.GetAllUser(c.Request().Context(), page, limit)
	return util.Negotiate(c, http.StatusOK, pagedUser)
}

// GetUser godoc
// @Summary Get a user
// @Description Get a user item
// @Tags users
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "User ID"
// @Success 200 {object} model.User
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /users/{id} [get]
// @Security ApiKeyAuth
func (userController *UserController) GetUser(c echo.Context) error {
	id := c.Param("id")
	if id == "me" {
		id = util.GetUserIdFromToken(c)
	}

	user, err := userController.userRepository.GetUser(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return util.Negotiate(c, http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user item
// @Tags users
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "User ID"
// @Param user body model.UserInput true "User Info"
// @Success 200 {object} model.User
// @Failure 400 {object} handler.APIError
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /users/{id} [put]
// @Security ApiKeyAuth
func (userController *UserController) UpdateUser(c echo.Context) error {
	id := c.Param("id")

	payload := new(model.UserInput)

	if err := util.BindAndValidate(c, payload); err != nil {
		return err
	}

	user, err := userController.userRepository.UpdateUser(c.Request().Context(), id, &model.User{UserInput: payload})
	if err != nil {
		return err
	}
	return util.Negotiate(c, http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a new user item
// @Tags users
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "User ID"
// @Success 204 {object} model.User
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /users/{id} [delete]
// @Security ApiKeyAuth
func (userController *UserController) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	err := userController.userRepository.DeleteUser(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

func beforeSave(user *model.User) (err error) {
	hashedPassword, err := util.EncryptPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}
