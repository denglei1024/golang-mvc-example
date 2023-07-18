package utils

import (
	"github.com/gin-gonic/gin"
)

// handle success response
func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

// handle fail response
func Fail(c *gin.Context, code int, message string) {
	c.JSON(200, gin.H{
		"code":    code,
		"message": message,
		"data":    nil,
	})
}

// handle error response
func Error(c *gin.Context, err error) {
	c.JSON(200, gin.H{
		"code":    500,
		"message": "internal error",
		"data":    err.Error(),
	})
}
