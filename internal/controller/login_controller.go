package controller

import (
	"ecom-project/internal/service"
	"ecom-project/pkg/response"

	"github.com/gin-gonic/gin"
)

// management controller user login ??? chua hieu ro lam
var Login = new(cUserLogin)

type cUserLogin struct {
}

func (cUser *cUserLogin) Login(ctx *gin.Context) {
	// Implement login logic here
	err := service.UserLogin().Login(ctx)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorInvalidOTP)
	}

	response.SuccessResponse(ctx, response.ErrorCodeSuccess)
}
