// Gin Router

package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router := gin.New()
	// Same as gin.Default() but as gin.Default() show all the information on console screen gin.New() will not show
	//So if we use it and still we want to show all the info on console we can use...
	// router.Use(gin.Logger())

	router.GET("/getData", getData)
	router.POST("/getDataPost", getDataPost)
	router.GET("/getQueryString", getQueryString)
	router.GET("/getUrlData/:name/:age", getUrlData)
	router.Run()
}

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi i am a Get Method from Gin Framework",
	})
}

// if this is the url request http://localhost:8080/getQueryString?name=Mark&age=30
func getQueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"data": "Hi i am in getQueryString method",
		"name": name,
		"age":  age,
	})
}

// if this is the url request http://localhost:8080/getUrlData/name/Mark/age=30
func getUrlData(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"data": "Hi i am in getQueryString method",
		"name": name,
		"age":  age,
	})
}

func getDataPost(c *gin.Context) {
	body := c.Request.Body
	value, _ := ioutil.ReadAll(body)
	// Here we can store the data sent by the user in the body and then store it in variable called value
	// and then convert it into string and again store it in response writer in the format of Json in the
	// variable called bodyData
	c.JSON(200, gin.H{
		"data":     "Hi i am a Post Method from Gin Framework",
		"bodyData": string(value),
	})
}

/*
	what is Gin
	How to use/install Gin
	Gin basics
	Routing with Gin. i.e GET and POST
    Handle Query string with GIN
	Handle Url parameters with GIN
*/
