package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	v1 "github.com/penthious/weather-api/api/handlers/v1"
	"github.com/rs/zerolog"
	"net/http"
	"os"
)

type APIMuxConfig struct {
	Log         *zerolog.Logger
	ServiceName string
	Shutdown    chan os.Signal
}

func APIMux(cfg APIMuxConfig) http.Handler {
	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			cfg.Log.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))

	eh := errorHandler{Log: cfg.Log}

	e.HTTPErrorHandler = eh.handle

	v1.Routes(e)

	return e
}

type errorHandler struct {
	Log *zerolog.Logger
}

func (eh errorHandler) handle(err error, c echo.Context) {
	c.Logger().Error(err)
}
