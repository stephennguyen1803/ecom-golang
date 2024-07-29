package routers

import (
	"github.com/gin-gonic/gin"

	uc "ecom-project/internal/controller"
	um "ecom-project/internal/model"
	ur "ecom-project/internal/repo"
	us "ecom-project/internal/service"
)

func NewServer() *gin.Engine {
	r := gin.Default()

	//User handle depenedency injection
	user := um.NewUser()
	userRepo := ur.NewUserRepo(user)
	userSer := us.NewUserService(userRepo)

	v1 := r.Group("/v1/2024")
	{
		v1.GET("/user/:name", uc.NewUserController(userSer).GetUserInfo)
		v1.GET("/userById/:id", uc.NewUserController(userSer).GetUserById)
		v1.PUT("/user", uc.NewUserController(userSer).GetUserInfo)
		v1.PATCH("/user", uc.NewUserController(userSer).GetUserInfo)
	}

	return r
}
