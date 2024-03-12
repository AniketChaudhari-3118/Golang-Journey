package main

import (
	"errors"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()
	logger.Error().Err(errors.New("file open Failed")).Msg("Something Happened!")
	// OR
	//logger.Fatal().Err(err).Msg("something catastrophic happened!")

	//If you want to log a FATAL or PANIC level message without calling os.Exit(1) and panic() respectively,
	// you must use the WithLevel() method
	err := errors.New("failed to connect to database")
	logger.WithLevel(zerolog.FatalLevel).Err(err).Msg("Something catastrophic happened!")

	//While the above output gives you details about the error that occurred,
	//it does not show the path of code execution that led to the error which can be crucial for debugging the issue.
	// You can fix this by including a stack trace in your error log through the Stack() method on an Event,
	// but before it can have an effect, you must assign zerolog.
	//ErrorStackMarshaler to a function that can extract stack traces from an error.
	//You can combine pkg/errors with the zerolog/pkgerrors helper to add a stack trace to an error log

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	newlogger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	newlogger.Error().
		Stack().
		Err(errors.New("file open failed")).
		Msg("Something Happened!")
}
