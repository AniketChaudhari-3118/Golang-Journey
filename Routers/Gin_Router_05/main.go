package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//Using basic Auth functionality with GIN
	auth := gin.BasicAuth(gin.Accounts{
		"user":   "pass",
		"user2":  "pass2",
		"Aniket": "Aniket@123",
	})

	//Route Grouping in GIN
	router.GET("/getUrlData/:name/:age", getUrlData)
	admin := router.Group("/admin", auth)
	{
		admin.GET("/getData", getData)
	}
	client := router.Group("/client", auth)
	{
		client.GET("/getQueryString", getQueryString)
	}

	//Custom HTTP configuration with GIN
	server := &http.Server{
		Addr:         ":9091",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
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

/*
Custom HTTP configuration with GIN
Route Grouping in GIN
Using basic Auth functionality with GIN
*/
