package middlewares

import "github.com/gin-gonic/gin"

func ErrorHandler() gin.HandlerFunc {

	//todo implement error handler
	return func(c *gin.Context) {
		c.Next()
	}
}
