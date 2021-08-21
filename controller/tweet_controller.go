package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/cbuelvasc/cinema-backend/exception"
	"github.com/cbuelvasc/cinema-backend/model"
	"github.com/cbuelvasc/cinema-backend/repository"
	"github.com/cbuelvasc/cinema-backend/util"
	"github.com/labstack/echo/v4"
)

type TweetControllerInterface interface {
	GetAllTweet(c echo.Context) error
	GetTweet(c echo.Context) error
	SaveTweet(c echo.Context) error
	DeleteTweet(c echo.Context) error
}

type TweetController struct {
	tweetRepository repository.TweetRepository
	userRepository  repository.UserRepository
}

func NewTweetController(tweetRepository repository.TweetRepository, userRepository repository.UserRepository) *TweetController {
	return &TweetController{
		tweetRepository: tweetRepository,
		userRepository:  userRepository,
	}
}

// GetAllTweet godoc
// @Summary Get all tweets
// @Description Get all tweet items
// @Tags tweets
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(xml, json)
// @Param page query int false "page" minimum(1)
// @Param limit query int false "size" minimum(1)
// @Param userId query string true "userId"
// @Success 200 {array} model.Tweet
// @Failure 500 {object} handler.APIError
// @Router /tweets [get]
// @Security ApiKeyAuth
func (tweetController *TweetController) GetAllTweet(c echo.Context) error {
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64)
	userId := c.QueryParam("userId")
	if len(userId) < 1 {
		return exception.ParameterException("userId")
	}

	_, err := tweetController.userRepository.GetUser(c.Request().Context(), userId)
	if err != nil {
		return err
	}

	pagedUser, err := tweetController.tweetRepository.GetAllTweets(c.Request().Context(), page, limit, userId)
	if err != nil {
		return err
	}
	return util.Negotiate(c, http.StatusOK, pagedUser)
}

// GetTweet godoc
// @Summary Get a tweet
// @Description Get a tweet item
// @Tags tweets
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "Tweet ID"
// @Success 200 {object} model.Tweet
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /tweets/{id} [get]
// @Security ApiKeyAuth
func (tweetController *TweetController) GetTweet(c echo.Context) error {
	id := c.Param("id")
	if id == "me" {
		id = util.GetUserIdFromToken(c)
	}

	tweet, err := tweetController.tweetRepository.GetTweet(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return util.Negotiate(c, http.StatusOK, tweet)
}

// SaveTweet godoc
// @Summary Create a tweet
// @Description Create a new tweet item
// @Tags tweets
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param tweet body model.TweetInput true "New tweet"
// @Success 200 {object} model.Tweet
// @Failure 400 {object} handler.APIError
// @Failure 409 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /tweets [post]
// @Security ApiKeyAuth
func (tweetController *TweetController) SaveTweet(c echo.Context) error {
	payload := new(model.TweetInput)
	if err := util.BindAndValidate(c, payload); err != nil {
		return err
	}

	_, err := tweetController.userRepository.GetUser(c.Request().Context(), payload.UserId)
	if err != nil {
		return err
	}

	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()
	tweet := &model.Tweet{TweetInput: payload}

	createdTweet, err := tweetController.tweetRepository.SaveTweet(c.Request().Context(), tweet)
	if err != nil {
		return err
	}

	return util.Negotiate(c, http.StatusCreated, createdTweet)
}

// DeleteTweet godoc
// @Summary Delete a tweet
// @Description Delete a new tweet item
// @Tags tweets
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "Tweet ID"
// @Success 204 {object} model.Tweet
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /tweets/{id} [delete]
// @Security ApiKeyAuth
func (tweetController *TweetController) DeleteTweet(c echo.Context) error {
	id := c.Param("id")
	userId := c.Param("userId")

	fmt.Println(id)
	fmt.Println(userId)

	_, errUser := tweetController.userRepository.GetUser(c.Request().Context(), userId)
	if errUser != nil {
		return errUser
	}

	e := tweetController.tweetRepository.DeleteTweet(c.Request().Context(), id, userId)
	if e != nil {
		return e
	}
	return c.NoContent(http.StatusNoContent)
}
