//What is a middleware
//How to use middleware in go
//Apply Middleware to routes, routes group and Whole application at once.

package main

import (
	"Gin_Router_06/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	//router.Use(middleware.Authenticate) //Apply to all routes

	//Apply middleware for group of routes
	admin := router.Group("/admin", middleware.Authenticate, middleware.AddHeader)
	{
		admin.GET("/getData1", getData1)
		admin.GET("/getData2", getData2)
	}

	//Apply middleware for one route
	router.GET("/getData", middleware.Authenticate, middleware.AddHeader, getData)

	router.Run()
}

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi i am getData Method",
	})
}

func getData1(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi i am getData1 Method",
	})
}

func getData2(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi i am getData2 Method",
	})
}
