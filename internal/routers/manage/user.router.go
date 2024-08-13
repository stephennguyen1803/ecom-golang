package manage

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// private route - require authentication
	userRouterPrivate := router.Group("/admin/user")
	{
		//call middleware
		// 1 - rate limit
		// userRouterPrivate.Use(middlewares.RateLimitMiddleware())
		// 2 - authentication
		// userRouterPrivate.Use(middlewares.AuthenMiddleware())
		// 3 - authorization
		// userRouterPrivate.Use(middlewares.AuthorizeMiddleware())
		userRouterPrivate.POST("/active_user")
	}
}
