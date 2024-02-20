package weather_api_handler

import (
	"github.com/labstack/echo/v4"
	"github.com/penthious/weather-api/external/weather_api"
	"github.com/rs/zerolog"
	"net/http"
	"strconv"
)

type Handlers struct {
	Log *zerolog.Logger
}

func (h Handlers) GetWeather(ctx echo.Context) error {
	latParam := ctx.QueryParam("lat")
	longParam := ctx.QueryParam("long")

	lat, err := strconv.ParseFloat(latParam, 64)
	if err != nil {
		h.Log.Error().Err(err).Msg("lat query parse")
		return ctx.String(http.StatusBadRequest, "lat query must be number")
	}

	long, err := strconv.ParseFloat(longParam, 64)
	if err != nil {
		h.Log.Error().Err(err).Msg("long query parse")
		return ctx.String(http.StatusBadRequest, "long query must be number")
	}

	weather, err := weather_api.GetWeather(lat, long)
	if err != nil {
		h.Log.Error().Err(err).Msg("get weather api failed")
		return ctx.String(http.StatusInternalServerError, "Something broke")
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
