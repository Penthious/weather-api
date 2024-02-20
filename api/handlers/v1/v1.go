package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/penthious/weather-api/api/handlers/v1/weather_api_handler"
	"github.com/rs/zerolog"
)

func Routes(e *echo.Echo, log *zerolog.Logger) {
	weatherHandlers := weather_api_handler.Handlers{
		Log: log,
	}
	v1 := e.Group("v1")
	v1.GET("/weather", weatherHandlers.GetWeather)
}
