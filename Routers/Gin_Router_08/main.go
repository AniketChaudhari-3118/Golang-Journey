// What is logrus
// Installing and using logrus
// LogLevels in Logrus
// Log messages to multiple options. //We can print messages to the console and log file as well
// Format messages in logrus
// Logging in JSON Format
// LogwithField and LogwithFields in logrus

package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetReportCaller(true) // To give and store each and every information like function name, file name,line number the caller etc

	//Everything will be printed in text format
	// logrus.SetFormatter(&logrus.TextFormatter{
	// 	DisableTimestamp: false,
	// 	FullTimestamp:    true,
	// })

	//Everything will be printed in Json format
	logrus.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: true,
		PrettyPrint:      true,
	})

	f, _ := os.Create("logrus.log")
	multi := io.MultiWriter(f, os.Stdout)
	logrus.SetOutput(multi)

	logrus.Traceln("Trace")
	// logrus.SetLevel(logrus.DebugLevel)
	logrus.Debugln("Debug")
	logrus.Infoln("info")
	logrus.Warnln("Warn")
	logrus.Errorln("Error")
	// logrus.Panicln("Panic") // Whenever we call panic the application will exit with the status code of 2 giving an error
	// logrus.Fatalln("Fatal") // Whenever we call panic the application will exit with the status code of 1 giving an error

	router := gin.Default()
	logrus.Println("Hi i am logrus")
	router.GET("/getData", getData)
	router.Run()
}

func getData(c *gin.Context) {
	//It prints all the information like what it is doing, filename, which func is called, its log level and msg
	logrus.WithField("Info", "CreateFile").Info("Starting File creation Info")
	logrus.WithField("Debug", "CreateFile").Debug("Starting File creation Info")
	f, err := os.Create("logrus.log")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Method": "CreateFile",
			"error":  true,
		}).Error(err.Error())
	}
	logrus.WithField("Info", "CreateFile").Debug("End File Creation")

	multi := io.MultiWriter(f, os.Stdout)
	logrus.SetOutput(multi)

	c.JSON(200, gin.H{
		"data": "This is the getData Handler",
	})
}

//Different log levels are provided by the log levels
/*
	Trace
	Debug
	Info
	Warn
	Error
	Fatal
	Panic

	//If i say Trace then below it, it will print all info from Debug-->Panic
	//If i say Warn then below it, it will print all info from Warn-->Panic
*/
