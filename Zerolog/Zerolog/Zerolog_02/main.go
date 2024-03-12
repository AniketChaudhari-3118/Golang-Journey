package main

import (
	"os"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout)

	logger.Info().
		Str("Name", "Aniket").
		Int("Age", 22).
		Bool("registered", true).
		Msg("New Signup!")
	// Send()
	// If not using Msg() then use Send(), if we dont use any method than it will give us blank screen
}
