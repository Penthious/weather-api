package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/penthious/weather-api/api/handlers/v1/weather_api_handler"
)

func Routes(e *echo.Echo) {
	weatherHandlers := weather_api_handler.Handlers{}
	v1 := e.Group("v1")
	v1.GET("/weather", weatherHandlers.GetWeather)
}
