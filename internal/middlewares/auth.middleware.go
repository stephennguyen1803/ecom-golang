package middlewares

import (
	"ecom-project/pkg/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before Run !!!")
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrorTokenInvalid)
			c.Abort()
			return
		}
		c.Next()
		fmt.Println("After Run !!!")
	}
}
