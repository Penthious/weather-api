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

	v1.Routes(e, cfg.Log)

	return e
}
