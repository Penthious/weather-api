package api

import (
	"context"
	"fmt"
	"github.com/penthious/weather-api/api/handlers"
	"github.com/penthious/weather-api/foundation/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
)

// Start builds out base requirements for the server
func Start() {
	appName := "weather-app"

	log := logger.New(logger.Config{
		App:   appName,
		Now:   time.Now().String(),
		Debug: true,
	})

	if err := server(log, appName); err != nil {
		log.Error().Err(err).Msg("startup")
	}
}

func server(log *zerolog.Logger, appName string) error {
	log.Info().Str("status", "").Msg("startup")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	apiMux := handlers.APIMux(handlers.APIMuxConfig{
		Log:         log,
		ServiceName: appName,
		Shutdown:    shutdown,
	})
	api := http.Server{
		Addr:    ":8080",
		Handler: apiMux,
	}

	serverErrors := make(chan error, 1)
	go func() {
		log.Info().Str("host", api.Addr).Msg("Application started")
		serverErrors <- api.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server err: %w", err)

	case sig := <-shutdown:
		log.Info().Any("signal", sig).Msg("shutdown started")
		defer log.Info().Any("signal", sig).Msg("shutdown finished")

		//Give outstanding requests a deadline for completion
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			_ = api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}
	return nil
}
