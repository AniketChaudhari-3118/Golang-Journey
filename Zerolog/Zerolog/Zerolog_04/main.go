package main

import (
	"github.com/rs/zerolog"
	"os"
)

var logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

func main() {

	// newlogger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// newlogger.UpdateContext(func(c zerolog.Context) zerolog.Context {
	// 	return c.Str("Name", "John")
	// })
	// //This UpdateContext() method updates the Logger's internal context in place (without creating a copy)
	// newlogger.Info().Msg("Info Message")

	mainLogger := logger.With().Str("Service", "Main").Logger()
	mainLogger.Info().Msg("Main logger message")

	auth()
	admin()
}

func auth() {
	authlogger := logger.With().Str("Service", "auth").Logger()
	authlogger.Info().Msg("auth logger message")
}

func admin() {
	adminlogger := logger.With().Str("Service", "admin").Logger()
	adminlogger.Info().Msg("admin logger message")
}

// We can define this as, logger is a parent logger and mainLogger, adminlogger and authlogger are child logger of logger
