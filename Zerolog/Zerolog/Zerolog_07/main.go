package main

import (
	"os"

	"github.com/rs/zerolog"
)

func main() {
	file, err := os.OpenFile("myapp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	logger := zerolog.New(file).With().Timestamp().Caller().Logger()
	logger.Info().Msg("Info Message")
}



//Logging to a File