package middleware

import (
	"github.com/gin-gonic/gin"
)

//One way to write middleware
// func Authenticate() gin.HandlerFunc {
// 	//Write custom logicto be applied before my middlware is executed

// 	return func(c *gin.Context) {
// 		if !(c.Request.Header.Get("Token") == "auth") {
// 			c.AbortWithStatusJSON(500, gin.H{
// 				"message": "Token not present",
// 			})
// 			c.Next()
// 		}
// 	}
// }

// Other way to write middleware
func Authenticate(c *gin.Context) {
	c.Request.Header.Set("Token", "auth")
	if !(c.Request.Header.Get("Token") == "auth") {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Token not present",
		})
		c.Next()
	}
}

func AddHeader(c *gin.Context) {
	c.Writer.Header().Set("Key", "Value")
	//c.Writer.Header().Set("Aniket", "Chaudhari")
	c.Next()
}
