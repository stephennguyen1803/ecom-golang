package initialize

import (
	"ecom-project/global"
	"ecom-project/internal/middlewares"
	"ecom-project/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	// Serve static files from the "assets" directory - should be used for development only
	r.Static("/assets/images", "./assets/images")

	//call middleware
	//r.Use(middlewares.Logger())
	r.Use(middlewares.AuthenMiddleware())
	//r.Use(middlewares.AuthorizeMiddleware())
	//r.Use(middlewares.CORSMiddleware())
	//r.Use(middlewares.RateLimitMiddleware())

	manageGroupRouter := routers.RouterGroupApp.Manage
	userGroupRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1/2024")
	{
		MainGroup.GET("/chekcStatus") // tracking monitor - check health of service
	}
	{
		userGroupRouter.UserRouter.InitUserRouter(MainGroup)
		userGroupRouter.ProductRouter.InitProductRouter(MainGroup)
	}
	{
		manageGroupRouter.UserRouter.InitUserRouter(MainGroup)
		manageGroupRouter.AdminRouter.InitAdminRouter(MainGroup)
	}

	return r
}
