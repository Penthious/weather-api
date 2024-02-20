package logger

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

type Config struct {
	App   string
	Now   string
	Debug bool
}

func New(c Config) *zerolog.Logger {
	cw := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339Nano}
	zlog := zerolog.New(cw).
		With().
		Dict("ctx",
			zerolog.Dict().
				Str("APP", c.App).
				Str("TIME", c.Now),
		).
		Stack().
		Timestamp().
		Logger()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if c.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	return &zlog
}
