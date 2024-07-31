package controller

import "github.com/gin-gonic/gin"

// using for health check server service
func PingController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
