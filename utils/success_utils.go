package utils

import "github.com/gin-gonic/gin"

func ProcessSuccess(c *gin.Context, message string, status int, data any) {
	if data != nil {
		c.JSON(status, gin.H{
			"message" : message,
			"data" : data,
		})
	} else{
		c.JSON(status, gin.H{
			"message" : message,
		})
	}
}