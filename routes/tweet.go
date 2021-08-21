package routes

import (
	"github.com/cbuelvasc/cinema-backend/controller"
	"github.com/cbuelvasc/cinema-backend/enums"
	"github.com/labstack/echo/v4"
)

func GetTweetApiRoutes(e *echo.Echo, tweetController *controller.TweetController) {
	v1 := e.Group(enums.BasePath)
	{
		v1.GET(enums.GetTweets, tweetController.GetAllTweet)
		v1.GET(enums.GetTweetById, tweetController.GetTweet)
		v1.POST(enums.CreateTweets, tweetController.SaveTweet)
		//v1.PUT(enums.UpdateTweetById, tweetController.UpdateTweet)
		v1.DELETE(enums.DeleteTweetById, tweetController.DeleteTweet)
	}
}
