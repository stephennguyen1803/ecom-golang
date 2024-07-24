package routers

import (
	"github.com/gin-gonic/gin"

	uc "ecom-project/internal/controller"
	um "ecom-project/internal/model"
	ur "ecom-project/internal/repo"
)

func NewServer() *gin.Engine {
	r := gin.Default()

	//User handle depenedency injection
	user := um.NewUser()
	userRepo := ur.NewUserRepo(user)

	v1 := r.Group("/v1/2024")
	{
		v1.GET("/user/:name", uc.NewUserController(userRepo).GetUserName)
		v1.PUT("/user", uc.NewUserController(userRepo).GetUserName)
		v1.PATCH("/user", uc.NewUserController(userRepo).GetUserName)
	}

	return r
}
