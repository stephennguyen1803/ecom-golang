package user

import (
	"ecom-project/internal/controller"
	"ecom-project/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	//public route - not require authentication

	// init controller
	// Using Wire Go to inject dependency
	// Dependency Injection
	userController, _ := wire.InitUserRouterHandler()

	userRouterPublic := router.Group("/user")
	{
		//call middleware
		userRouterPublic.POST("/register", userController.Register) // register -> YES -> send OTP -> verify OTP -> create account ||
		userRouterPublic.POST("/otp")
		userRouterPublic.POST("/login", controller.Login.Login)
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
