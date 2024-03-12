package main

import (
	"os"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
)

func main() {
	buildInfo, _ := debug.ReadBuildInfo()

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Int("pid", os.Getpid()).
		Str("go_version", buildInfo.GoVersion).
		Logger()

	logger.Trace().Msg("Trace message")
	logger.Debug().Msg("Debug message")
	logger.Info().Msg("Info message")
	logger.Warn().Msg("Warn message")
	logger.Error().Msg("Error message")
	logger.WithLevel(zerolog.FatalLevel).Msg("Fatal message")
	logger.WithLevel(zerolog.PanicLevel).Msg("Panic message")
}
