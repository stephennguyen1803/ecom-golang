package controller

import (
	"ecom-project/internal/service"
	"ecom-project/internal/vo"
	"ecom-project/pkg/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) Register(c *gin.Context) {
	var params vo.UserRegistratorRequest

	// Bind the incoming JSON to the combined struct
	err := c.BindJSON(&params)
	if err != nil {
		response.ErrorResponse(c, response.ErrorUserBadRequest)
		return
	}

	// Process Email registration if email is provided
	if params.Email != "" {
		result := uc.userService.Register(params.Email, params.Purpose)
		if result != response.ErrorCodeSuccess {
			response.ErrorResponse(c, result)
			return
		}
		response.SuccessResponse(c, fmt.Sprintf("Register success with Email %s", params.Email))
		return
	}

	// Process Phone registration if phone is provided
	if params.Phone != "" {
		result := uc.userService.Register(params.Phone, params.Purpose)
		if result != response.ErrorCodeSuccess {
			response.ErrorResponse(c, result)
			return
		}
		response.SuccessResponse(c, fmt.Sprintf("Register success with Phone %s", params.Phone))
		return
	}

	// If neither email nor phone is provided, return an error
	response.ErrorResponse(c, response.ErrorUserBadRequest)
}

// controller -> service -> repo -> model -> dbs
// func (uc *UserController) GetUserInfo(c *gin.Context) {
// 	response.SuccessResponse(c, uc.userService.GetUserSerivce())
// }

// func (uc *UserController) GetUserById(c *gin.Context) {
// 	response.ErrorResponse(c, response.ErrorCodeParamInvalid)
// }
