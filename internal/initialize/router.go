package initialize

import (
	"ecom-project/internal/controller"
	"ecom-project/internal/middlewares"
	"ecom-project/internal/model"
	"ecom-project/internal/repo"
	"ecom-project/internal/service"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// Router initialization
	r := gin.Default()
	//call middleware
	r.Use(middlewares.AuthenMiddleware())

	//User handle depenedency injection
	user := model.NewUser()
	userRepo := repo.NewUserRepo(user)
	userSer := service.NewUserService(userRepo)

	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", controller.PingController)
		v1.GET("/user/:name", controller.NewUserController(userSer).GetUserInfo)
		v1.GET("/userById/:id", controller.NewUserController(userSer).GetUserById)
		v1.PUT("/user", controller.NewUserController(userSer).GetUserInfo)
		v1.PATCH("/user", controller.NewUserController(userSer).GetUserInfo)
	}

	return r
}
