package main

import (
	"Gin_Router_09/config"
	"Gin_Router_09/database"
	"Gin_Router_09/logger"
	"Gin_Router_09/router"
)

func main() {
	config.Init()
	config.Appconfig = config.GetConfig()
	logger.Init()
	logger.InfoLn("Logger Initialized successfully")

	database.Init()
	if config.Appconfig.GetBool("senddata") {
		// logic to send data to database
		logger.InfoLn("Data Sent Successfully")
	}

	logger.InfoLn("Started Router Initializing")
	router.Init()
	logger.InfoLn("Router Initializing Successfully")
}
