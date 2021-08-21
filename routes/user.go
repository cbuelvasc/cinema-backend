package routes

import (
	"github.com/cbuelvasc/cinema-backend/controller"
	"github.com/cbuelvasc/cinema-backend/enums"
	"github.com/labstack/echo/v4"
)

func GetUserApiRoutes(e *echo.Echo, userController *controller.UserController) {
	v1 := e.Group(enums.BasePath)
	{
		v1.POST(enums.SignIn, userController.AuthenticateUser)
		v1.POST(enums.SignUp, userController.SaveUser)

		v1.GET(enums.GetUsers, userController.GetAllUser)
		v1.GET(enums.GetUserById, userController.GetUser)
		v1.PUT(enums.UpdateUserById, userController.UpdateUser)
		v1.DELETE(enums.DeleteUserById, userController.DeleteUser)

	}
}
