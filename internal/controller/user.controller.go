package controller

import (
	"ecom-project/internal/service"
	"ecom-project/pkg/request"
	"ecom-project/pkg/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

// Define it for easy to test
type UserServiceInterface interface {
	GetUserSerivce() string
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) Register(c *gin.Context) {
	// email := c.PostForm("email")
	// purpose := c.PostForm("purpose")

	var userRequestBody request.UserRequestBody
	if err := c.BindJSON(&userRequestBody); err != nil {
		response.ErrorResponse(c, response.ErrorUserBadRequest)
		return
	}
	email := userRequestBody.Email
	purpose := userRequestBody.Purpose

	result := uc.userService.Register(email, purpose)
	if result != response.ErrorCodeSuccess {
		response.ErrorResponse(c, result)
		return
	}
	response.SuccessResponse(c, fmt.Sprintf("Register success with Email %s", email))
}

// controller -> service -> repo -> model -> dbs
// func (uc *UserController) GetUserInfo(c *gin.Context) {
// 	response.SuccessResponse(c, uc.userService.GetUserSerivce())
// }

// func (uc *UserController) GetUserById(c *gin.Context) {
// 	response.ErrorResponse(c, response.ErrorCodeParamInvalid)
// }
