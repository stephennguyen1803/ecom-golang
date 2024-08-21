package routers

import (
	"github.com/gin-gonic/gin"

	"ecom-project/internal/middlewares"
)

func NewServer() *gin.Engine {
	r := gin.Default()
	//call middleware
	r.Use(middlewares.AuthenMiddleware())

	//tuong dung ma ko dung =)) this is non-dependency injection
	//User handle depenedency injection
	// user := um.NewUser()
	// userRepo := ur.NewUserRepo(user)
	// userSer := us.NewUserService(userRepo)

	// v1 := r.Group("/v1/2024")
	// {
	// 	v1.GET("/ping", uc.PingController)
	// 	v1.GET("/user/:name", uc.NewUserController(userSer).GetUserInfo)
	// 	v1.GET("/userById/:id", uc.NewUserController(userSer).GetUserById)
	// 	v1.PUT("/user", uc.NewUserController(userSer).GetUserInfo)
	// 	v1.PATCH("/user", uc.NewUserController(userSer).GetUserInfo)
	// }

	return r
}
