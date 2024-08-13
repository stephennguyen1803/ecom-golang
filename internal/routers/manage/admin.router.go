package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (ar *AdminRouter) InitAdminRouter(router *gin.RouterGroup) {
	// public router
	adminRouterPublic := router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}
	// private router
	adminRouterPrivate := router.Group("/admin")
	{
		//call middleware
		// 1 - rate limit
		// adminRouterPrivate.Use(middlewares.RateLimitMiddleware())
		// 2 - authentication
		// adminRouterPrivate.Use(middlewares.AuthenMiddleware())
		// 3 - authorization
		// adminRouterPrivate.Use(middlewares.AuthorizeMiddleware())
		adminRouterPrivate.POST("/active_user")
	}
}
