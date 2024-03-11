//Gin Router

package main

import (
	"github.com/gin-gonic/gin"
	// "net/http"
)

func main() {
	router := gin.Default()
	//We create a new Gin router using gin.Default(), which comes with some default middleware (like logging and recovery).

	//Logging Middleware: This middleware logs every incoming HTTP request. It logs information such as the HTTP method, request path, response status code,
	//and the time taken to process the request. Logging is essential for debugging, monitoring, and analyzing the behavior of
	// your web server.

	//Recovery Middleware: This middleware recovers from any panics that occur during request handling. If
	// a handler function panics (e.g., due to a runtime error), the recovery middleware catches the panic, logs the error,
	// and sends a 500 Internal Server Error response to the client. This prevents the entire server from crashing
	//due to a single request causing a panic.
	
	//we can use many functions like router.GET || router.POST || router.Delete || router.PUT etc.
	router.GET("/getData", func(c *gin.Context) { //if the route is comming as get request and the route is /getData then call the func.
		c.JSON(200, gin.H{ //func(c *gin.Context)  is an Handler (The gin.Context contains all the information of response and the request(each and every bit of information) which is stored in a variable called 'c')
			"data": "Hi i am Gin framework",
			//c.JSOn is context.JSON which is object of the Response Writer
			//200 is a response code, gin.H is a kind of interface
			//here we are saying that return this output in the form of JSON with an Response code of 200
		})
	})

	router.Run() // By default it uses 8080, if we want to use other port than we can specify using router.Run(":5000")
}
