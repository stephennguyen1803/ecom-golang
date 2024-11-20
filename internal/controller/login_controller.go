package controller

import (
	"ecom-project/global"
	"ecom-project/internal/model"
	"ecom-project/internal/service"
	"ecom-project/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

func (cUser *cUserLogin) VerifyOTP(ctx *gin.Context) {
	var params model.VerifyOTPInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeParamInvalid)
		return
	}

	result, err := service.UserLogin().VerifyOTP(ctx, &params)
	if err != nil {
		global.Logger.Error("Error Verifying OTP: ", zap.Error(err), zapcore.Field{Key: "params", Type: zapcore.ObjectMarshalerType})
		response.ErrorResponse(ctx, response.ErrorInvalidOTP)
		return
	}

	response.SuccessResponse(ctx, result)
}

func (cUser *cUserLogin) Register(ctx *gin.Context) {
	var params model.RegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeParamInvalid)
		return
	}

	codeStatus, err := service.UserLogin().Register(ctx, &params)
	if err != nil {
		global.Logger.Error("Error Registering OTP: ", zap.Error(err), zapcore.Field{Key: "params", Type: zapcore.ObjectMarshalerType})
		response.ErrorResponse(ctx, codeStatus)
		return
	}

	response.SuccessResponse(ctx, codeStatus)
}
