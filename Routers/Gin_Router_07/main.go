// Logging in GIN
// How default logging works in GIN
// Define format for the log of routes in GIN
// Define format of logs with GIN
// Write logs to file in Gin
// Controlling log output colouring in console with GIN
// Logging in JSON Format in GIN. (Real Work Situation)

package main

import (
	"Gin_Router_07/logger"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
)

func main() {
	router := gin.Default()

	//Define format for the log of routes in GIN
	//Before Defining the format
	// GET    /getData                  --> main.getData (3 handlers)
	gin.DebugPrintRouteFunc = func(httpMethod string, absolutePath string, handlerName string, nuHandlers int) {
		log.Printf("Endpoint formatted information is %v %v %v %v\n", absolutePath, httpMethod, handlerName, nuHandlers)
	}
	//After Defining the Format
	//Endpoint formatted information is /getData GET main.getData 3

	// Controlling log output colouring in console with GIN
	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()

	// Write logs to file in Gin
	f, _ := os.Create("ginLogging.log")
	//gin.DefaultWriter = io.MultiWriter(f) // If i want to write the data of logging to the file
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// Define format of logs with GIN
	router.Use(gin.LoggerWithFormatter(logger.FormatLogsJson))

	router.GET("/getData", getData)
	router.Run()
}

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hii this is getData Route",
	})
}
