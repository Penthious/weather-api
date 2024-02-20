package weather_api_handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/penthious/weather-api/external/weather_api"
	"net/http"
	"strconv"
)

type Handlers struct{}

func (h Handlers) GetWeather(ctx echo.Context) error {
	latParam := ctx.QueryParam("lat")
	longParam := ctx.QueryParam("long")

	lat, err := strconv.ParseFloat(latParam, 64)
	if err != nil {
		ctx.Error(fmt.Errorf("lat must be number"))
		return ctx.String(http.StatusBadRequest, "")
	}

	long, err := strconv.ParseFloat(longParam, 64)
	if err != nil {
		return fmt.Errorf("long must be number")
	}

	weather, err := weather_api.GetWeather(lat, long)

	if err != nil {
		return fmt.Errorf("getWeatherApi: %w", err)
	}

	var perceivedTemp string

	if weather.Main.Temp > 80 {
		perceivedTemp = "hot"
	} else if weather.Main.Temp < 50 {
		perceivedTemp = "cold"
	} else {
		perceivedTemp = "moderate"
	}

	type response struct {
		PerceivedTemp string
		Temp          float64
		Weather       string
	}

	res := response{
		PerceivedTemp: perceivedTemp,
		Temp:          weather.Main.Temp,
		Weather:       weather.Weather[0].Description,
	}

	return ctx.JSON(200, res)
}
