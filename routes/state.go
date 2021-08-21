package routes

import (
	"github.com/cbuelvasc/cinema-backend/controller"
	"github.com/cbuelvasc/cinema-backend/enums"
	"github.com/labstack/echo/v4"
)

func GetStateApiRoutes(e *echo.Echo, stateController *controller.StateController) {
	v1 := e.Group(enums.BasePath)
	{
		v1.POST(enums.CreateState, stateController.SaveState)
		v1.GET(enums.GetStates, stateController.GetAllStates)
		v1.GET(enums.GetStateById, stateController.GetState)
		v1.PUT(enums.UpdateStateById, stateController.UpdateState)
		v1.DELETE(enums.DeleteStateById, stateController.DeleteState)
	}
}
