package exception

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ResourceNotFoundException(resourceName string, fieldName string, fieldValue string) error {
	msg := fmt.Sprintf("%s not found with %s: %s", resourceName, fieldName, fieldValue)
	return echo.NewHTTPError(http.StatusNotFound, msg)
}

func TweetNotFoundException(resourceName string, fieldNameOne string, fieldValueOne string, fieldNameTwo string, fieldValueTwo string) error {
	msg := fmt.Sprintf("%s not found with %s: %s and %s: %s", resourceName, fieldNameOne, fieldValueOne, fieldNameTwo, fieldValueTwo)
	return echo.NewHTTPError(http.StatusBadRequest, msg)
}

func BadRequestException(msg string) error {
	return echo.NewHTTPError(http.StatusBadRequest, msg)
}

func NotFoundRequestException(msg string) error {
	return echo.NewHTTPError(http.StatusNotFound, msg)
}

func ParameterException(parameterName string) error {
	msg := fmt.Sprintf("Parameter named %s is required", parameterName)
	return echo.NewHTTPError(http.StatusBadRequest, msg)
}

func ConflictException(resourceName string, fieldName string, fieldValue string) error {
	msg := fmt.Sprintf("%s with %s: %s already exists", resourceName, fieldName, fieldValue)
	return echo.NewHTTPError(http.StatusConflict, msg)
}

func UnauthorizedException() error {
	return echo.ErrUnauthorized
}
