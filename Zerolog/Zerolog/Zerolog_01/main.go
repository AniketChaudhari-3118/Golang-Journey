package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Hello from zerolog global logger")
	// The above line prints the JSON formated log enrty to the console

	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	log.Error().Msg("Error Message")

	logger := zerolog.New(os.Stdout)
	logger.Info().Msg("logger is a new logger instance")
	//logger.With().Str("key", "value")   //log a message with additional fields using the `With` method, which allows you to attach structured data to the log entry.
}
