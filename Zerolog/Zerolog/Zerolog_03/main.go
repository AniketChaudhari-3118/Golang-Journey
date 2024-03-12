package main

import (
	"os"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	// The With() method which returns a zerolog.Context instance that allows you to add additional
	// properties to the logger in key-value pairs through field methods similar to those on the zerolog.Event type.

	//By using Caller() we can add file and line number to all log entries

	logger.Info().Msg("Hello! this Info with new feild timestamp")
	logger.Debug().Str("Username", "Aniket").Send()
	
}
