package controller

import (
	"ecom-project/global"
	"ecom-project/internal/model"
	"ecom-project/internal/service"
	"ecom-project/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

// VerifyOTP godoc
// @Summary      Verify Account Using OTP
// @Description  Verfiy OTP login by user
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body   model.VerifyOTPInput  true  "payload"
// @Success      200  {object}  response.Response
// @Failure      200  {object}  response.Response
// @Failure      500
// @Router       /user/verify_account [post]
func (cUser *cUserLogin) VerifyOTP(ctx *gin.Context) {
	var params model.VerifyOTPInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeParamInvalid)
		return
	}

	result, err := service.UserLogin().VerifyOTP(ctx, &params)
	if err != nil {
		global.Logger.Error("Error Verifying OTP: ", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrorInvalidOTP)
		return
	}

	response.SuccessResponse(ctx, result)
}

// Register godoc
// @Summary      Register User
// @Description  Register User Using Verify Key
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body   model.RegisterInput  true  "payload"
// @Success      200  {object}  response.Response
// @Failure      200  {object}  response.Response
// @Failure      500
// @Router       /user/register [post]
func (cUser *cUserLogin) Register(ctx *gin.Context) {
	var params model.RegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeParamInvalid)
		return
	}

	codeStatus, err := service.UserLogin().Register(ctx, &params)
	if err != nil {
		global.Logger.Error("Error Registering OTP: ", zap.Error(err))
		response.ErrorResponse(ctx, codeStatus)
		return
	}

	response.SuccessResponse(ctx, codeStatus)
}

// Register godoc
// @Summary      Update User Password
// @Description  Update User Password Using  User Token
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body   model.UpdateUserPasswordInput  true  "payload"
// @Success      200  {object}  response.Response
// @Failure      200  {object}  response.Response
// @Failure      500
// @Router       /user/update_password_register [post]
func (cUser *cUserLogin) UpdatePasswordRegister(ctx *gin.Context) {
	// Implement UpdatePasswordRegister logic here
	var params model.UpdateUserPasswordInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeParamInvalid)
		return
	}

	result, err := service.UserLogin().UpdatePasswordRegister(ctx, params.UserToken, params.UserPassword)
	if err != nil {
		global.Logger.Error("Error Update User Password: ", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrorInvalidOTP)
		return
	}

	response.SuccessResponse(ctx, result)
}
