package controller

import (
	"ecom-project/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserControler struct {
	UserService UserServiceInterface
}

// Define it for easy to test
type UserServiceInterface interface {
	GetUserSerivce() string
}

func NewUserController(UserService UserServiceInterface) *UserControler {
	return &UserControler{UserService: UserService}
}

// controller -> service -> repo -> model -> dbs
func (uc *UserControler) GetUserInfo(c *gin.Context) {
	response.SuccessResponse(c, uc.UserService.GetUserSerivce())
}

func (uc *UserControler) GetUserById(c *gin.Context) {
	response.ErrorResponse(c, response.ErrorCodeParamInvalid)
}
