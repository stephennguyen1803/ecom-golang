package user

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	//public route - not require authentication
	userRouterPublic := router.Group("/user")
	{
		//call middleware
		userRouterPublic.GET("/register") // register -> YES -> send OTP -> verify OTP -> create account ||
		userRouterPublic.POST("/otp")
	}
	//private route - require authentication
	userRouterPrivate := router.Group("/user")
	{
		//call middleware
		// 1 - rate limit
		// userRouterPrivate.Use(middlewares.RateLimitMiddleware())
		// 2 - authentication
		// userRouterPrivate.Use(middlewares.AuthenMiddleware())
		// 3 - authorization
		// userRouterPrivate.Use(middlewares.AuthorizeMiddleware())
		userRouterPrivate.GET("/profile")
	}
}
